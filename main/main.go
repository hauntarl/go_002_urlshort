package main

import (
	"flag"
	"fmt"
	"net/http"

	urlshort "hauntarl/gophercises/url-short"
)

func main() {
	yamlFlag := flag.String("yaml", "yaml_urls.yaml", `accepts a yaml file format: 
- path: /path
  url: url-redirect
  ...
	`)
	jsonFlag := flag.String("json", "json_urls.json", `accepts a json file format:
[
	{"path": "/path", "url": "url-redirect"},
	...
]
	`)
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/golang":      "https://github.com/hauntarl/golang",
		"/gophercises": "https://courses.calhoun.io/courses/cor_gophercises",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /yaml
  url: https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc
- path: /yaml_github
  url: https://github.com/go-yaml/yaml
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	// yamlFileHandler using yamlHandler as fallback
	yamlFileHandler, err := urlshort.YAMLFileHandler(*yamlFlag, yamlHandler)
	if err != nil {
		panic(err)
	}

	// jsonHandler using yamlFileHandler as fallback
	json := `
[
	{"path": "/json", "url": "https://pkg.go.dev/encoding/json?tab=doc"},
	{"path": "/json_intro", "url": "https://blog.golang.org/json"}
]
`
	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlFileHandler)
	if err != nil {
		panic(err)
	}

	// jsonFileHandler using jsonHandler as fallback
	jsonFileHandler, err := urlshort.JSONFileHandler(*jsonFlag, jsonHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonFileHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
