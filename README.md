# Url Shortener

Code an http.Handler that forwards paths to other URLs (similar to Bitly).

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Command:**

- go run main\main.go

**Features:**

- grouping packages using go.mod
- setting up a basic http server
- redirecting urls
- parsing yaml input format

**Packages explored:**

- fmt
- net/http
- gopkg.in/yaml.v2

**Bonus:**

- flag package: to get yaml file name
- os package: to open and close the file
- YamlFileHandler to parse file.yaml and return http.HandleFunc
