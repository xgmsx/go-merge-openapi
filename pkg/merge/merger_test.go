package merge_test

import (
	"os"
	"testing"

	"github.com/xgmsx/go-merge-openapi/pkg/merge"
	"gopkg.in/yaml.v3"
)

func TestMerge(t *testing.T) {
	node, err := merge.Merge("../../examples/api/petstore.openapi.yaml")
	if err != nil {
		t.Fatalf("merge failed: %v", err)
	}
	if node == nil {
		t.Fatal("want non-nil result")
	}
}

func TestWrite(t *testing.T) {
	node := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "key"},
			{Kind: yaml.ScalarNode, Value: "value"},
		},
	}

	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "YAML file",
			path: "./test_output.yaml",
			want: "key: value\n",
		},
		{
			name: "JSON file",
			path: "./test_output.json",
			want: "{\n  \"key\": \"value\"\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() { _ = os.Remove(tt.path) }()

			// ext := strings.ToLower(filepath.Ext(tt.path))
			err := merge.Write(node, tt.path)
			if err != nil {
				t.Fatalf("write failed: %v", err)
			}

			if _, err := os.Stat(tt.path); os.IsNotExist(err) {
				t.Fatal("want output file to be created")
			}

			data, err := os.ReadFile(tt.path)
			if err != nil {
				t.Fatalf("failed to read output file: %v", err)
			}

			got := string(data)
			if got != tt.want {
				t.Fatalf("unexpected file content: got %q, want %q", got, tt.want)
			}
		})
	}
}
