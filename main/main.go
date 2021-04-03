package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	boltdb "hauntarl/gophercises/url-short"

	"github.com/boltdb/bolt"
	"gopkg.in/yaml.v2"
)

const (
	usageYAML = `
accepts a yaml file format: 
- path: /path
  url: url-redirect
...
`
	usageJSON = `
accepts a json file format:
[
	{"path": "/path", "url": "url-redirect"},
	...
]
`
)

var flagYAML, flagJSON *string

func init() {
	flagYAML = flag.String("yaml", "yaml_urls.yaml", usageYAML)
	flagJSON = flag.String("json", "json_urls.json", usageJSON)
	flag.Parse()
	log.Printf("%-15s: -yaml=%s, -json=%s\n", "flags parsed", *flagYAML, *flagJSON)
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%-15s: data bucket created\n", boltdb.BucketName)

	if err = insertData(db); err != nil {
		log.Fatalln(err)
	}
	log.Println("starting server at http://localhost:8080/")
	log.Fatalln(http.ListenAndServe(":8080", makeHandler(db, defaultMux())))
}

func setupDB() (*bolt.DB, error) {
	db, err := boltdb.CreateDB("urlshort.db")
	if err != nil {
		return nil, err
	}
	log.Printf("%-15s: database created\n", boltdb.DBName)

	err = boltdb.CreateBucket(db, "DB")
	return db, err
}

var (
	dataMap = map[string]string{
		"/golang":      "https://github.com/hauntarl/golang",
		"/gophercises": "https://courses.calhoun.io/courses/cor_gophercises",
	}
	dataYAML = `
- path: /yaml
  url: https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc
- path: /yaml_github
  url: https://github.com/go-yaml/yaml
`
	dataJSON = `
[
	{"path": "/json", "url": "https://pkg.go.dev/encoding/json?tab=doc"},
	{"path": "/json_intro", "url": "https://blog.golang.org/json"}
]
`
)

func insertData(db *bolt.DB) (err error) {
	if err = boltdb.UpdateBucket(db, dataMap); err != nil {
		return
	}
	if err = boltdb.ReadData(dataYAML, db, yaml.Unmarshal); err != nil {
		return
	}
	if err = boltdb.ReadFile(*flagYAML, db,
		func(file io.Reader) boltdb.Decoder {
			return yaml.NewDecoder(file)
		}); err != nil {
		return
	}
	if err = boltdb.ReadData(dataJSON, db, json.Unmarshal); err != nil {
		return
	}
	if err = boltdb.ReadFile(*flagJSON, db,
		func(file io.Reader) boltdb.Decoder {
			return json.NewDecoder(file)
		}); err != nil {
		return
	}
	return
}

func makeHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			path = r.URL.Path
			url  string
			err  error
		)
		url, err = boltdb.ReadBucket(db, path)
		if err != nil {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, string(url), http.StatusFound)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the home page of url shortner, where all unregistered paths gets redirected to.")
}
