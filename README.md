# Url Shortener

Code an http.Handler that forwards paths to other URLs (similar to Bitly).

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Command:**

- go run main\main.go
- go run main\main.go -h (--help) (to get information regarding flags)
- go run main\main.go --yaml file-name.yaml -json=file-name.json

**Urls:**

- /golang : <https://github.com/hauntarl/golang>
- /gophercises : <https://courses.calhoun.io/courses/cor_gophercises>
* /yaml : <https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc>
* /yaml_github : <https://github.com/go-yaml/yaml>
- /quiz : <https://github.com/hauntarl/go_001_quiz>
- /quiz_readme : <https://github.com/hauntarl/go_001_quiz/blob/master/README.md>
* /json : <https://pkg.go.dev/encoding/json?tab=doc>
* /json_intro :<https://blog.golang.org/json>
- /urlshort : <https://github.com/hauntarl/go_002_urlshort>
- /urlshort_readme : https://github.com/hauntarl/go_002_urlshort/blob/master/README.md

**Features:**

- grouping packages using go.mod
- setting up a basic http server
- redirecting urls
- parsing yaml input format
* YamlFileHandler to parse file.yaml and return http.HandleFunc

**Packages explored:**

- fmt
- net/http - to setup a basic http server and redirect requests
- gopkg.in/yaml.v2 - to work with yaml data
* flag package: to get yaml file name
* os package: to open and close the file
