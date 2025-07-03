package openapi

import (
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func ResolveReferences(node *yaml.Node, baseDir string, visited map[string]*yaml.Node) error {
	if node == nil {
		return nil
	}

	switch node.Kind {
	case yaml.MappingNode:
		// Check if this is a $ref object
		if len(node.Content) >= 2 && node.Content[0].Value == "$ref" {
			refPath := node.Content[1].Value
			resolved, err := resolveReference(refPath, baseDir, visited)
			if err != nil {
				return err
			}
			*node = *resolved
			return nil
		}

		// Traverse each key-value pair
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 < len(node.Content) {
				if err := ResolveReferences(node.Content[i+1], baseDir, visited); err != nil {
					return err
				}
			}
		}

	case yaml.SequenceNode:
		for _, item := range node.Content {
			if err := ResolveReferences(item, baseDir, visited); err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveReference(refPath, baseDir string, visited map[string]*yaml.Node) (*yaml.Node, error) {
	var filePath, fragment string

	if strings.Contains(refPath, "#") {
		parts := strings.SplitN(refPath, "#", 2)
		filePath = parts[0]
		fragment = parts[1]
	} else {
		filePath = refPath
	}

	if !filepath.IsAbs(filePath) {
		filePath = filepath.Join(baseDir, filePath)
	}

	cacheKey := filePath
	if cachedNode, ok := visited[cacheKey]; ok {
		if fragment != "" {
			return extractFragment(cachedNode, fragment)
		}
		return cachedNode, nil
	}

	refNode, err := LoadYAMLNode(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load referenced file '%s': %w", filePath, err)
	}

	visited[cacheKey] = refNode

	if err := ResolveReferences(refNode, filepath.Dir(filePath), visited); err != nil {
		return nil, err
	}

	if fragment != "" {
		return extractFragment(refNode, fragment)
	}

	return refNode, nil
}

func extractFragment(node *yaml.Node, fragment string) (*yaml.Node, error) {
	if fragment == "" || fragment == "/" {
		return node, nil
	}

	fragment = strings.TrimPrefix(fragment, "/")

	for _, part := range strings.Split(fragment, "/") {
		if part == "" {
			continue
		}

		if node.Kind != yaml.MappingNode {
			return nil, fmt.Errorf("cannot navigate fragment path '%s': not a mapping", fragment)
		}

		found := false
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 < len(node.Content) && node.Content[i].Value == part {
				node = node.Content[i+1]
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("fragment path '%s' not found", fragment)
		}
	}

	return node, nil
}
