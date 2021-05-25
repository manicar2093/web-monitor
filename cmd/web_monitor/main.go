package main

import (
	"embed"
	"io"
	"log"
	"os"
)

// go:embed templates/*
var tpl embed.FS

func main() {

	log.Println(os.Getwd())

	f, err := tpl.Open("index.html")
	if err != nil {
		log.Println("no se pudo abrir embed", err)
		os.Exit(1)
	}

	c, err := io.ReadAll(f)
	if err != nil {
		log.Println("no se pudo leer embed")
		os.Exit(1)
	}
	log.Println(string(c))

}
