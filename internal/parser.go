package internal

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	py_ts "github.com/tree-sitter/tree-sitter-python/bindings/go"
	// go_ts "github.com/tree-sitter/tree-sitter-go/bindings/go"
	// js_ts "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
	//py_ts "github.com/tree-sitter/tree-sitter-python/bindings/go"
	// rs_ts "github.com/tree-sitter/tree-sitter-rust/bindings/go"
	// ts_ts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

func ExtractImports(content []byte, lang string) []string {
	if lang != "python" {
		return []string{}
	}

	parser := sitter.NewParser()

	tsLang := sitter.NewLanguage(py_ts.Language())

	parser.SetLanguage(tsLang)

	tree := parser.Parse(content, nil)

	root := tree.RootNode()

	imports := []string{}

	for i := uint(0); i < root.ChildCount(); i++ {
		child := root.Child(i)

		if child.Kind() != "import_statement" {
			continue
		}

		for j := uint(0); j < child.ChildCount(); j++ {
			grandchild := child.Child(j)

			if grandchild.Kind() == "dotted_name" {
				imports = append(imports, grandchild.Utf8Text(content))
			}
		}
	}

	return imports
}
