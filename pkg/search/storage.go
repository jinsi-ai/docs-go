package search

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Storage 存储实现
type Storage struct {
	db *sql.DB
}

// NewStorage 创建新的存储实例
func NewStorage(dataDir string) (*Storage, error) {
	sqlitePath := filepath.Join(dataDir, "search.db")
	
	// 确保数据目录存在
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("创建数据目录失败: %v", err)
	}

	db, err := sql.Open("sqlite3", sqlitePath)
	if err != nil {
		return nil, fmt.Errorf("打开SQLite数据库失败: %v", err)
	}

	s := &Storage{db: db}
	
	// 初始化数据库表
	if err := s.initTables(); err != nil {
		db.Close()
		return nil, fmt.Errorf("初始化数据库表失败: %v", err)
	}

	return s, nil
}

// initTables 初始化数据库表
func (s *Storage) initTables() error {
	// 创建文档表
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS documents (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			path TEXT NOT NULL,
			updated_at DATETIME NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("创建文档表失败: %v", err)
	}

	// 创建倒排索引表
	_, err = s.db.Exec(`
		CREATE TABLE IF NOT EXISTS inverted_index (
			term TEXT NOT NULL,
			document_id TEXT NOT NULL,
			positions TEXT NOT NULL, -- 存储位置数组的JSON字符串
			frequency INTEGER NOT NULL,
			PRIMARY KEY (term, document_id)
		)
	`)
	if err != nil {
		return fmt.Errorf("创建倒排索引表失败: %v", err)
	}

	// 尝试创建全文搜索索引（如果FTS5可用）
	_, err = s.db.Exec(`
		CREATE VIRTUAL TABLE IF NOT EXISTS documents_fts USING fts5(
			id UNINDEXED,
			title,
			content,
			path UNINDEXED,
			content=documents,
			content_rowid=rowid
		)
	`)
	if err != nil {
		// 如果FTS5不可用，记录警告但继续初始化
		log.Printf("警告: FTS5全文搜索不可用，将使用倒排索引搜索: %v", err)
	} else {
		// 创建触发器以保持FTS索引与主表同步
		_, err = s.db.Exec(`
			CREATE TRIGGER IF NOT EXISTS documents_ai AFTER INSERT ON documents BEGIN
				INSERT INTO documents_fts(rowid, title, content) VALUES (new.rowid, new.title, new.content);
			END;
		`)
		if err != nil {
			log.Printf("警告: 创建插入触发器失败: %v", err)
		}

		_, err = s.db.Exec(`
			CREATE TRIGGER IF NOT EXISTS documents_ad AFTER DELETE ON documents BEGIN
				INSERT INTO documents_fts(documents_fts, rowid, title, content) VALUES('delete', old.rowid, old.title, old.content);
			END;
		`)
		if err != nil {
			log.Printf("警告: 创建删除触发器失败: %v", err)
		}

		_, err = s.db.Exec(`
			CREATE TRIGGER IF NOT EXISTS documents_au AFTER UPDATE ON documents BEGIN
				INSERT INTO documents_fts(documents_fts, rowid, title, content) VALUES('delete', old.rowid, old.title, old.content);
				INSERT INTO documents_fts(rowid, title, content) VALUES (new.rowid, new.title, new.content);
			END;
		`)
		if err != nil {
			log.Printf("警告: 创建更新触发器失败: %v", err)
		}
	}

	return nil
}

// SaveDocument 保存文档到数据库
func (s *Storage) SaveDocument(doc *DocumentIndex) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 插入或更新文档
	_, err = tx.Exec(`
		INSERT OR REPLACE INTO documents (id, title, content, path, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, doc.ID, doc.Title, doc.Content, doc.Path, doc.UpdatedAt)
	if err != nil {
		return fmt.Errorf("保存文档失败: %v", err)
	}

	return tx.Commit()
}

// GetDocument 根据ID获取文档
func (s *Storage) GetDocument(id string) (*DocumentIndex, error) {
	var doc DocumentIndex
	err := s.db.QueryRow(`
		SELECT id, title, content, path, updated_at
		FROM documents WHERE id = ?
	`, id).Scan(&doc.ID, &doc.Title, &doc.Content, &doc.Path, &doc.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

// DeleteDocument 删除文档
func (s *Storage) DeleteDocument(id string) error {
	_, err := s.db.Exec("DELETE FROM documents WHERE id = ?", id)
	return err
}

// GetAllDocuments 获取所有文档
func (s *Storage) GetAllDocuments() ([]*DocumentIndex, error) {
	rows, err := s.db.Query(`
		SELECT id, title, content, path, updated_at
		FROM documents ORDER BY updated_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docs []*DocumentIndex
	for rows.Next() {
		var doc DocumentIndex
		if err := rows.Scan(&doc.ID, &doc.Title, &doc.Content, &doc.Path, &doc.UpdatedAt); err != nil {
			return nil, err
		}
		docs = append(docs, &doc)
	}

	return docs, nil
}

// SaveInvertedIndex 保存倒排索引
func (s *Storage) SaveInvertedIndex(term string, documentID string, positions []int) error {
	positionsJSON := fmt.Sprintf("[%s]", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(positions)), ","), "[]"))
	
	_, err := s.db.Exec(`
		INSERT OR REPLACE INTO inverted_index (term, document_id, positions, frequency)
		VALUES (?, ?, ?, ?)
	`, term, documentID, positionsJSON, len(positions))
	
	return err
}

// GetDocumentsByTerm 根据关键词获取文档列表
func (s *Storage) GetDocumentsByTerm(term string) ([]string, error) {
	rows, err := s.db.Query(`
		SELECT document_id FROM inverted_index 
		WHERE term = ? ORDER BY frequency DESC
	`, term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docIDs []string
	for rows.Next() {
		var docID string
		if err := rows.Scan(&docID); err != nil {
			return nil, err
		}
		docIDs = append(docIDs, docID)
	}

	return docIDs, nil
}

// ClearAll 清除所有数据
func (s *Storage) ClearAll() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 清空基础表
	tables := []string{"documents", "inverted_index"}
	for _, table := range tables {
		_, err := tx.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			return fmt.Errorf("清空表 %s 失败: %v", table, err)
		}
	}

	// 检查FTS表是否存在，如果存在则清空
	var ftsTableExists int
	err = tx.QueryRow(`
		SELECT COUNT(*) FROM sqlite_master 
		WHERE type='table' AND name='documents_fts'
	`).Scan(&ftsTableExists)
	
	if err == nil && ftsTableExists > 0 {
		_, err = tx.Exec("DELETE FROM documents_fts")
		if err != nil {
			return fmt.Errorf("清空表 documents_fts 失败: %v", err)
		}
	}

	return tx.Commit()
}

// GetStats 获取统计信息
func (s *Storage) GetStats() (*IndexStats, error) {
	var totalDocs int
	err := s.db.QueryRow("SELECT COUNT(*) FROM documents").Scan(&totalDocs)
	if err != nil {
		return nil, err
	}

	var lastUpdated time.Time
	err = s.db.QueryRow("SELECT MAX(updated_at) FROM documents").Scan(&lastUpdated)
	if err != nil {
		// 如果没有文档，返回默认时间
		lastUpdated = time.Time{}
	}

	// 获取数据库文件大小（由于SQLite连接不直接提供文件路径，暂时返回0）
	var fileSize int64 = 0

	return &IndexStats{
		TotalDocuments: totalDocs,
		LastUpdated:    lastUpdated,
		IndexSize:      fileSize,
	}, nil
}

// Close 关闭存储
func (s *Storage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// FullTextSearch 全文搜索（使用SQLite FTS5或回退到倒排索引）
func (s *Storage) FullTextSearch(query string, limit int, offset int) ([]*SearchResult, error) {
	// 检查FTS5表是否存在
	var tableExists int
	err := s.db.QueryRow(`
		SELECT COUNT(*) FROM sqlite_master 
		WHERE type='table' AND name='documents_fts'
	`).Scan(&tableExists)
	
	if err != nil || tableExists == 0 {
		// FTS5不可用，返回空结果，让搜索器使用倒排索引
		return []*SearchResult{}, nil
	}

	// 构建FTS查询
	ftsQuery := fmt.Sprintf("title:%s OR content:%s", query, query)
	
	rows, err := s.db.Query(`
		SELECT d.id, d.title, d.content, d.path, d.updated_at,
		       snippet(documents_fts, 2, '<mark>', '</mark>', '...', 64) as excerpt,
		       bm25(documents_fts) as score
		FROM documents_fts fts
		JOIN documents d ON d.id = fts.id
		WHERE documents_fts MATCH ?
		ORDER BY score
		LIMIT ? OFFSET ?
	`, ftsQuery, limit, offset)
	
	if err != nil {
		// 如果FTS查询失败，返回空结果，让搜索器使用倒排索引
		log.Printf("FTS5查询失败，回退到倒排索引: %v", err)
		return []*SearchResult{}, nil
	}
	defer rows.Close()

	var results []*SearchResult
	for rows.Next() {
		var doc DocumentIndex
		var excerpt string
		var score float64
		
		if err := rows.Scan(&doc.ID, &doc.Title, &doc.Content, &doc.Path, &doc.UpdatedAt, &excerpt, &score); err != nil {
			return nil, err
		}

		results = append(results, &SearchResult{
			Document: &doc,
			Excerpt:  excerpt,
			Score:    score,
		})
	}

	return results, nil
}

// GetTotalSearchCount 获取搜索结果总数
func (s *Storage) GetTotalSearchCount(query string) (int, error) {
	ftsQuery := fmt.Sprintf("title:%s OR content:%s", query, query)
	
	var count int
	err := s.db.QueryRow(`
		SELECT COUNT(*)
		FROM documents_fts fts
		JOIN documents d ON d.id = fts.id
		WHERE documents_fts MATCH ?
	`, ftsQuery).Scan(&count)
	
	if err != nil {
		return 0, err
	}

	return count, nil
}