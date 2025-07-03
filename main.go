package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/xgmsx/go-merge-openapi/pkg/merge"
)

func main() {
	var configPath, outputPath string
	flag.StringVar(&configPath, "config", "", "Path to OpenAPI config file")
	flag.StringVar(&outputPath, "output", "", "Optional output file path")
	flag.Parse()

	if configPath == "" {
		log.Fatal("missing -config argument")
	}

	merged, err := merge.Merge(configPath)
	if err != nil {
		log.Fatalf("merge failed: %v", err)
	}

	if outputPath == "" {
		base := strings.TrimSuffix(filepath.Base(configPath), filepath.Ext(configPath))
		outputPath = filepath.Join(filepath.Dir(configPath), base+".merged.yaml")
	}

	if err := merge.Write(merged, filepath.Clean(outputPath)); err != nil {
		log.Fatalf("write failed: %v", err)
	}

	log.Print("merged to", outputPath)
}
