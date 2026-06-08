package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Aeres-u99/hermes/v2/internal"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var test string
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "%s - The Code Map you will Ever need!", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(w, "Have Fun!")
	}
	flag.StringVar(&test, "input", "code.py", "Code to Parse")
	flag.Parse()
	content, err := os.ReadFile(test)
	if err != nil {
		panic(err)
	}
	output := internal.Output{
		Version:   1,
		Generated: time.Now().UTC().Format(time.RFC3339),
		Files:     make(map[string]internal.FileInfo),
		Index:     make(map[string]internal.Location),
	}
	lang := detectLanguage(test)
	loc := len(strings.Split(string(content), "\n"))
	output.Files[test] = internal.FileInfo{
		Lang:    lang,
		LOC:     loc,
		Imports: []string{},
		Symbols: []internal.Symbol{},
	}
	data, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func detectLanguage(path string) string {
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
