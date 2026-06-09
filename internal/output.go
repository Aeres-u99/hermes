package internal

import (
	"time"
)

func BuildOutput(result *AnalysisResult) Output {
	output := Output{
		Version:   1,
		Generated: time.Now().UTC().Format(time.RFC3339),
		Files:     make(map[string]FileInfo),
		Index:     make(map[string]Location),
	}

	for path, fileInfo := range result.Files {
		output.Files[path] = fileInfo
	}

	for k, v := range result.Index {
		output.Index[k] = v
	}

	return output
}
