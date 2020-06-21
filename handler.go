package urlshort

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := urlMap[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type pathURL struct {
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
	URL  string `yaml:"url,omitempty" json:"url,omitempty"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlSlice []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlSlice)
	if err != nil {
		return nil, err
	}
	urlMap := buildMap(pathUrls)
	return MapHandler(urlMap, fallback), nil
}

func parseYaml(yamlSlice []byte) ([]pathURL, error) {
	var pathURLs []pathURL
	// Unmarshal function accepts []byte and converts it to given structure
	err := yaml.Unmarshal(yamlSlice, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

// YAMLFileHandler opens, reads and parses the given yaml file
func YAMLFileHandler(yamlFileName string, fallback http.Handler) (http.HandlerFunc, error) {
	yamlFile, err := os.Open(yamlFileName)
	if err != nil {
		return nil, err
	}
	defer yamlFile.Close()

	pathUrls, err := parseYamlFile(yamlFile)
	if err != nil {
		return nil, err
	}

	urlMap := buildMap(pathUrls)
	return MapHandler(urlMap, fallback), nil
}

func parseYamlFile(yamlFileName io.Reader) ([]pathURL, error) {
	var pathUrls []pathURL
	// yaml Decoder accepts io.Reader, read bytes from it and converts it to
	// given data structure
	err := yaml.NewDecoder(yamlFileName).Decode(&pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func buildMap(pathURLs []pathURL) map[string]string {
	// iterates through parsed yaml and converts them to map[string]string
	urlMap := make(map[string]string)
	for _, pu := range pathURLs {
		urlMap[pu.Path] = pu.URL
	}
	return urlMap
}

// JSONHandler parses data to map and returns http.HandleFunc
func JSONHandler(jsonSlice []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJSON(jsonSlice)
	if err != nil {
		return nil, err
	}

	urlMap := buildMap(pathUrls)
	return MapHandler(urlMap, fallback), nil
}

func parseJSON(jsonSlice []byte) ([]pathURL, error) {
	var pathUrls []pathURL
	err := json.Unmarshal(jsonSlice, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

// JSONFileHandler accepts a file name, opens, parses it and returns http.HandlerFunc
func JSONFileHandler(jsonFileName string, fallback http.Handler) (http.HandlerFunc, error) {
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	pathUrls, err := parseJSONFile(jsonFile)
	if err != nil {
		return nil, err
	}

	urlMap := buildMap(pathUrls)
	return MapHandler(urlMap, fallback), nil
}

func parseJSONFile(jsonFile io.Reader) ([]pathURL, error) {
	var pathUrls []pathURL
	err := json.NewDecoder(jsonFile).Decode(&pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}
