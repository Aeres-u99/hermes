package internal

func MapKind(kind string) string {
	switch kind {
	case "class":
		return "cls"
	case "function":
		return "fn"
	case "member":
		return "method"
	case "variable":
		return "var"
	case "struct":
		return "struct"
	case "const":
		return "const"
	case "package":
		return "package"
	default:
		return kind
	}
}

func BuildSymbols(tags []CTag, file string) ([]Symbol, map[string]Location) {
	symbols := []Symbol{}
	index := make(map[string]Location)

	for _, tag := range tags {
		name := tag.Name

		if tag.Scope != "" {
			name = tag.Scope + "." + tag.Name
		}

		symbols = append(symbols, Symbol{
			Name: name,
			Type: MapKind(tag.Kind),
			Line: tag.Line,
		})

		index[name] = Location{
			File: file,
			Line: tag.Line,
		}
	}

	return symbols, index
}
