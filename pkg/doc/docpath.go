package doc

import (
	"os"
	"path/filepath"
	"strings"
)

// NormalizePath 规范化文档路径
func NormalizePath(docsDir, path string) string {
	// 移除默认基础路径前缀
	path = strings.TrimPrefix(path, "/")

	// 处理空路径或目录路径
	if path == "" || strings.HasSuffix(path, "/") {
		return SelectIndexFile(docsDir, "")
	}

	// 如果已经有扩展名，直接返回
	if strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".html") {
		return path
	}

	// 检查是否为目录
	dirPath := filepath.Join(docsDir, path)
	if stat, err := os.Stat(dirPath); err == nil && stat.IsDir() {
		return SelectIndexFile(docsDir, path+"/")
	}

	// 检查文件是否存在（优先.html，其次.md）
	extensions := []string{".html", ".md"}
	for _, ext := range extensions {
		filePath := filepath.Join(docsDir, path+ext)
		if _, err := os.Stat(filePath); err == nil {
			return path + ext
		}
	}

	// 对于 index 路径，如果 .md 不存在，尝试 README.md
	if path == "index" {
		for _, readme := range []string{"README.md", "readme.md"} {
			if _, err := os.Stat(filepath.Join(docsDir, readme)); err == nil {
				return readme
			}
		}
	}

	// 默认返回.md扩展名
	return path + ".md"
}

// SelectIndexFile 选择合适的索引文件（优先.html，其次.md）
func SelectIndexFile(docsDir, basePath string) string {
	// 定义文件优先级顺序
	candidateFiles := []string{
		"index.html",
		"readme.md",
		"README.md",
		"index.md",
	}

	// 遍历候选文件，返回第一个存在的文件
	for _, filename := range candidateFiles {
		filePath := filepath.Join(docsDir, basePath, filename)
		if _, err := os.Stat(filePath); err == nil {
			return basePath + filename
		}
	}

	// 否则返回默认的index.md
	return filepath.Join(basePath, "index.md")
}

// GenBreadcrumb 生成面包屑导航信息
func GenBreadcrumb(path string) []map[string]string {
	// 移除文件扩展名
	cleanPath := strings.TrimSuffix(strings.TrimSuffix(path, ".md"), ".html")

	// 分割路径并过滤空部分
	parts := strings.Split(cleanPath, string(os.PathSeparator))
	filteredParts := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			filteredParts = append(filteredParts, part)
		}
	}

	// 构建面包屑
	breadcrumb := make([]map[string]string, 0, len(filteredParts))
	currentPath := ""

	for i, part := range filteredParts {
		// 跳过最后的"index"
		if part == "index" && i == len(filteredParts)-1 {
			continue
		}

		// 构建当前路径
		if currentPath == "" {
			currentPath = part
		} else {
			currentPath = filepath.Join(currentPath, part)
		}

		breadcrumb = append(breadcrumb, map[string]string{
			"name": part,
			"path": "/" + currentPath,
		})
	}

	return breadcrumb
}
