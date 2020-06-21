package main

import (
	"flag"
	"fmt"
	"net/http"

	urlshort "hauntarl/gophercises/url-short"

	"github.com/boltdb/bolt"
)

var yamlFlag, jsonFlag *string

func main() {
	setFlags()
	db, err := setupDB()
	if err != nil {
		panic(err)
	}
	handleFunc := getHandlerFunc(db)

	fmt.Printf("\nStarting the server on localhost:8080\n\n")
	http.ListenAndServe(":8080", handleFunc)
}

func getHandlerFunc(db *bolt.DB) http.HandlerFunc {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/golang":      "https://github.com/hauntarl/golang",
		"/gophercises": "https://courses.calhoun.io/courses/cor_gophercises",
	}
	fmt.Println("\nFrom MapHandler...")
	mapHandler := urlshort.MapHandler(pathsToUrls, db, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := `
- path: /yaml
  url: https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc
- path: /yaml_github
  url: https://github.com/go-yaml/yaml
`
	_, err := urlshort.YAMLHandler([]byte(yaml), db, mux)
	if err != nil {
		panic(err)
	}

	// Build the YamlFileHandler using yamlHandler as fallback
	_, err = urlshort.YAMLFileHandler(*yamlFlag, db, mux)
	if err != nil {
		panic(err)
	}

	// Build the JsonHandler using yamlFileHandler as fallback
	json := `
[
	{"path": "/json", "url": "https://pkg.go.dev/encoding/json?tab=doc"},
	{"path": "/json_intro", "url": "https://blog.golang.org/json"}
]
`
	_, err = urlshort.JSONHandler([]byte(json), db, mux)
	if err != nil {
		panic(err)
	}

	// Build the JsonFileHandler using jsonHandler as fallback
	_, err = urlshort.JSONFileHandler(*jsonFlag, db, mux)
	if err != nil {
		panic(err)
	}
	return mapHandler
}

func setFlags() {
	yamlFlag = flag.String("yaml", "yaml_urls.yaml", `accepts a yaml file format: 
- path: /path
  url: url-redirect
  ...
	`)
	jsonFlag = flag.String("json", "json_urls.json", `accepts a json file format:
[
	{"path": "/path", "url": "url-redirect"},
	...
]
	`)
	flag.Parse()
	fmt.Printf("Flags set...\n\tyaml = %s\n\tjson = %s\n\n", *yamlFlag, *jsonFlag)
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("urlshort.db", 0666, nil)
	if err != nil {
		return nil, err
	}
	if err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	fmt.Print("Database initialized...\n\n")
	return db, nil
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	fmt.Println("Fallback http.HandleFunc created...")
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
