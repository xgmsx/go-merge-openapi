package openapi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func WriteSpec(node *yaml.Node, path string) error {
	var data []byte
	var err error

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		var spec interface{}
		if err := node.Decode(&spec); err != nil {
			return fmt.Errorf("failed to decode YAML node: %w", err)
		}
		data, err = json.MarshalIndent(spec, "", "  ")
	case ".yaml", ".yml":
		data, err = yaml.Marshal(node)
	default:
		data, err = yaml.Marshal(node)
	}

	if err != nil {
		return fmt.Errorf("failed to marshal spec: %w", err)
	}

	return os.WriteFile(filepath.Clean(path), data, 0o600)
}
