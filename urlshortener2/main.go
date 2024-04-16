package main

import (
	"fmt"
	"net/http"

	"github.com/Ekosetiawan993/urlshortener2/handler"
)

func main() {
	mux := defaultMux()
	// map of urls
	pathToURLs := map[string]string{
		"/porto-eko":   "https://ekosetiawan993.github.io/",
		"/gophercises": "https://gophercises.com/",
	}

	// handler that consume mux and pathtoURLs
	mapHandler := handler.MapHandler(pathToURLs, mux)

	// http.HandleFunc("/", mapHandler)

	// YAML
	yaml_urls := `
- path: /porto-eko-yaml
  url: https://ekosetiawan993.github.io/
- path: /gophercises-yaml
  url: https://gophercises.com/
`

	yamlHandler, err := handler.YAMLHandler([]byte(yaml_urls), mapHandler)

	if err != nil {
		panic(err)
	}

	// JSON
	json_urls := `[{"path": "/porto-eko-json", "url": "https://ekosetiawan993.github.io/"}, {"path": "/gophercises-json", "url": "https://gophercises.com/"}]`

	jsonHandler, err := handler.JSONHandler([]byte(json_urls), mapHandler)

	if err != nil {
		panic(err)
	}

	serve := "json"
	fmt.Println("Serve http on port: 8084")
	if serve == "yaml" {
		http.ListenAndServe(":8084", yamlHandler)
	} else {
		http.ListenAndServe(":8084", jsonHandler)
	}

	// http.ListenAndServe(":8084", nil)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.TryHandler)
	return mux
}
