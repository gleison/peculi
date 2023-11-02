package web

import (
	"embed"
	"io/fs"
)

//go:embed static/*
var StaticFiles embed.FS

func GetHtmlFile(name string) []byte {
	staticFS, err := fs.Sub(StaticFiles, "static")
	if err != nil {
		panic(err)
	}
	htmls := fs.FS(staticFS)
	bytes, err := fs.ReadFile(htmls, name)
	if err != nil {
		panic(err)
	}
	return bytes
}
