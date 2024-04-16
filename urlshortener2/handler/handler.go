package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func TryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func MapHandler(pathToURL map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToURL[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		// of the short link not found
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. parse yaml
	yamlPathURLs, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}

	// convert yaml data to url map
	URLMapResult := buildMap(yamlPathURLs)

	return MapHandler(URLMapResult, fallback), nil

}

type pathURLYAML struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(data []byte) ([]pathURLYAML, error) {
	var pathURLs []pathURLYAML
	err := yaml.Unmarshal(data, &pathURLs)

	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

func buildMap(yamlPath []pathURLYAML) map[string]string {
	yamlPathToURLMap := make(map[string]string)

	for _, pu := range yamlPath {
		yamlPathToURLMap[pu.Path] = pu.URL
	}

	return yamlPathToURLMap
}

type pathJSONURL struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathToURLJSON []pathJSONURL

	if err := json.Unmarshal(jsonBytes, &pathToURLJSON); err != nil {
		return nil, err
	}

	jsonURLMap := make(map[string]string)

	for _, puJSON := range pathToURLJSON {
		jsonURLMap[puJSON.Path] = puJSON.URL
	}

	return MapHandler(jsonURLMap, fallback), nil

}
