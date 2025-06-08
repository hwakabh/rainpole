package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
)

//go:embed .release-please-manifest.json
var version_strings []byte

type VersionSchema struct {
	Path string `json:"."`
}

type VersionReponseSchema struct {
	VersionString string `json:"version"`
}

func GetVersionFileContent() string {
	var version VersionSchema
	fmt.Println(version_strings)

	json.Unmarshal(version_strings, &version)
	return version.Path
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := VersionReponseSchema{
		VersionString: GetVersionFileContent(),
	}
	json.NewEncoder(w).Encode(body)
}
