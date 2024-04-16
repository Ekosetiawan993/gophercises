package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// handler for the url in map
func MapHandler(pathsToURL map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToURL[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		// if the short link not found
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the yaml data
	pathUrls, err := parseYAML(yamlBytes)
	// var pathUrls []pathUrl
	// err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	// 2. convert the data to url map
	pathToUrls := buildMap(pathUrls)
	// pathToUrls := make(map[string]string)

	// for _, pu := range pathUrls {
	// 	pathToUrls[pu.Path] = pu.URL
	// }

	// 3. return a map handler
	return MapHandler(pathToUrls, fallback), nil

}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

func parseYAML(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

// struct for parsing yaml data
type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

type pathsToURLJSON struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func JSONHandler(jsondata []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToURL []pathsToURLJSON

	if err := json.Unmarshal(jsondata, &pathsToURL); err != nil {
		return nil, err
	}
	fmt.Printf("%v", pathsToURL)
	// map url to paths
	jsonPathsURL := make(map[string]string)

	for _, puJSON := range pathsToURL {
		jsonPathsURL[puJSON.Path] = puJSON.URL
	}

	return MapHandler(jsonPathsURL, fallback), nil
}

func main() {
	mux := defaultMux()

	// map contain URLs
	pathsToUrls := map[string]string{
		"/porto-eko":   "https://ekosetiawan993.github.io/",
		"/gophercises": "https://gophercises.com/",
	}

	// if the urls list store in map
	mapHandler := MapHandler(pathsToUrls, mux)

	// if we store the data using yaml
	yaml_urls := `
- path: /porto-eko-yaml
  url: https://ekosetiawan993.github.io/
- path: /gophercises-yaml
  url: https://gophercises.com/
`
	yamlHandler, err := YAMLHandler([]byte(yaml_urls), mapHandler)

	if err != nil {
		panic(err)
	}

	// JSON
	json_urls := `[{"path": "/porto-eko-json", "url": "https://ekosetiawan993.github.io/"}, {"path": "/gophercises-json", "url": "https://gophercises.com/"}]`
	// ,
	// 		{"path": "/gophercises-json", "url": "https://gophercises.com/"}
	fmt.Println(json.Valid([]byte(json_urls)))

	jsonHandler, err := JSONHandler([]byte(json_urls), mapHandler)

	if err != nil {
		panic(err)
	}

	fmt.Println("Serve http on port : 8084")
	serve := "json"

	if serve == "json" {
		http.ListenAndServe(":8084", jsonHandler)
	} else {
		http.ListenAndServe(":8084", yamlHandler)
	}

}

// default server route
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello worldd...")
}
