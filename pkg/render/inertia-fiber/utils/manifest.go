package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Manifest map[string]struct {
	Css            []string `json:"css"`
	File           string   `json:"file"`
	IsEntry        bool     `json:"isEntry"`
	Imports        []string `json:"imports"`
	DynamicImports []string `json:"dynamicImports"`
	IsDynamicEntry bool     `json:"isDynamicEntry"`
	Src            string   `json:"src"`
}

func manifest(buildDirectory []string) Manifest {
	strings := append(buildDirectory, "public", "build", "manifest.json")
	path := filepath.Join(strings...)
	log.Println("manifest path:", path)

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var manifest Manifest
	err = json.Unmarshal(data, &manifest)
	if err != nil {
		panic(err)
	}

	//log.Println(manifest)

	return manifest
}
