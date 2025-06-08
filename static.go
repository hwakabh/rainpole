package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

//go:embed public/*
var publicDir embed.FS

func FetchHtmlFileServer() http.Handler {
	// load staticfiles
	public, err := fs.Sub(publicDir, "public")
	if err != nil {
		os.Exit(1)
		fmt.Println("Failed to load static file")
		fmt.Println(err)
	}
	return http.FileServer(http.FS(public))
}
