package internal

import "path/filepath"

func DetectLanguage(path string) string {
	switch filepath.Ext(path) {
	case ".py":
		return "python"
	case ".go":
		return "go"
	case ".js":
		return "javascript"
	case ".ts":
		return "typescript"
	case ".rs":
		return "rust"
	case ".c":
		return "c"
	case ".bazel":
		return "bazel"
	case ".cpp":
		return "cpp"
	case ".java":
		return "java"
	case ".lua":
		return "lua"
	default:
		return "Unknown"
	}
}
