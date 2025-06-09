package main

import (
	"embed"
	"fmt"
	"io"
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
		fmt.Println("Failed to load static file")
		fmt.Println(err)
		os.Exit(1)
	}
	return http.FileServer(http.FS(public))
}

func FetchBashrc(w http.ResponseWriter, r *http.Request) {
	url := "https://raw.githubusercontent.com/hwakabh/dotfiles/refs/heads/main/bash/.bashrc"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to invoke URL [ %s ]\n", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get responses. (Code: %v)\n", resp.StatusCode)
	}

	body, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		fmt.Println("Error reading response body:", err)
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(body))
}

func FetchGitConfig(w http.ResponseWriter, r *http.Request) {
	url := "https://raw.githubusercontent.com/hwakabh/dotfiles/refs/heads/main/gitconf/.gitconfig"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to invoke URL [ %s ]\n", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get responses. (Code: %v)\n", resp.StatusCode)
	}

	body, ioerr := io.ReadAll(resp.Body)
	if ioerr != nil {
		fmt.Println("Error reading response body:", err)
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(body))
}
