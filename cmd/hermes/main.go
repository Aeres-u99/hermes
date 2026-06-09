package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Aeres-u99/hermes/v2/internal"
	"os"
	"strings"
	"time"
)

func main() {
	var inputFile string
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "%s - The Code Map you will Ever need!", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(w, "Have Fun!")
	}
	flag.StringVar(&inputFile, "input", "code.py", "Code to Parse")
	flag.Parse()
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	output := internal.Output{
		Version:   1,
		Generated: time.Now().UTC().Format(time.RFC3339),
		Files:     make(map[string]internal.FileInfo),
		Index:     make(map[string]internal.Location),
	}
	lang := internal.DetectLanguage(inputFile)
	loc := len(strings.Split(string(content), "\n"))
	output.Files[inputFile] = internal.FileInfo{
		Lang:    lang,
		LOC:     loc,
		Imports: []string{},
		Symbols: []internal.Symbol{},
	}
	imports := internal.ExtractImports(content, lang)
	tags, err := internal.GetTags(inputFile)
	if err != nil {
		panic(err)
	}
	symbols, index := internal.BuildSymbols(tags, inputFile)
	for k, v := range index {
		output.Index[k] = v
	}
	fileInfo := output.Files[inputFile]
	fileInfo.Symbols = symbols
	fileInfo.Imports = imports
	output.Files[inputFile] = fileInfo
	data, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
