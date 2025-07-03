package merge

import (
	"fmt"
	"path/filepath"

	"github.com/xgmsx/go-merge-openapi/internal/openapi"
	"gopkg.in/yaml.v3"
)

func Merge(inputPath string) (*yaml.Node, error) {
	baseDir := filepath.Dir(inputPath)
	node, err := openapi.LoadYAMLNode(inputPath)
	if err != nil {
		return nil, fmt.Errorf("load failed: %w", err)
	}

	visited := make(map[string]*yaml.Node)
	if err := openapi.ResolveReferences(node, baseDir, visited); err != nil {
		return nil, fmt.Errorf("resolve failed: %w", err)
	}

	return node, nil
}

func Write(node *yaml.Node, outputPath string) error {
	return openapi.WriteSpec(node, outputPath)
}
