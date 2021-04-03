# Url Shortener

Code an http.Handler that forwards paths to other URLs (similar to Bitly).

Implementation of URL Shortener from gophercises, including the bonus section.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Commands:**

- go run main\main.go
- go run main\main.go -h (--help) (to get information regarding flags)
- go run main\main.go --yaml file-name.yaml -json=file-name.json

**Features:**

- grouping packages using go.mod
- using command-line flags
- parsing yaml bytes and files
- parsing json bytes and files
- setting up a basic http server
- redirecting requests using http
- persisting the url mapping in database

**Packages explored:**

- fmt
- net/http - to setup a basic http server and redirect requests
- [gopkg.in/yaml.v2](gopkg.in/yaml.v2) - to work with yaml data
- flag - to get yaml/json file name
- os - to open and close the file
- io - to read from file which satisfies io.Reader interface
- encoding/json - to work with json data
- [github.com/boltdb/bolt](github.com/boltdb/bolt) - to store and retrieve urls for specified path

**Test Output:**

``` terminal
D:\gophercises\url-short>go run main\main.go
2021/04/03 18:28:23 flags parsed   : -yaml=yaml_urls.yaml, -json=json_urls.json
2021/04/03 18:28:23 urlshort.db    : database created
2021/04/03 18:28:23 DB             : data bucket created
2021/04/03 18:28:23 inserted       : /gophercises        , https://courses.calhoun.io/courses/cor_gophercises
2021/04/03 18:28:23 inserted       : /golang             , https://github.com/hauntarl/golang
2021/04/03 18:28:23 inserted       : /yaml               , https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc
2021/04/03 18:28:23 inserted       : /yaml_github        , https://github.com/go-yaml/yaml
2021/04/03 18:28:24 inserted       : /quiz               , https://github.com/hauntarl/go_001_quiz
2021/04/03 18:28:24 inserted       : /quiz_readme        , https://github.com/hauntarl/go_001_quiz/blob/master/README.md
2021/04/03 18:28:24 inserted       : /json               , https://pkg.go.dev/encoding/json?tab=doc
2021/04/03 18:28:24 inserted       : /json_intro         , https://blog.golang.org/json
2021/04/03 18:28:24 inserted       : /urlshort           , https://github.com/hauntarl/go_002_urlshort
2021/04/03 18:28:24 inserted       : /urlshort_readme    , https://github.com/hauntarl/go_002_urlshort/blob/master/README.md
2021/04/03 18:28:24 starting server at http://localhost:8080/
2021/04/03 18:28:38 /                   : unregistered path
2021/04/03 18:28:38 /favicon.ico        : unregistered path
2021/04/03 18:28:59 /golang             : https://github.com/hauntarl/golang
2021/04/03 18:29:17 /favicon.ico        : unregistered path
2021/04/03 18:29:38 /yaml               : https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc
2021/04/03 18:29:41 /favicon.ico        : unregistered path
2021/04/03 18:29:58 /quiz_readme        : https://github.com/hauntarl/go_001_quiz/blob/master/README.md
2021/04/03 18:30:00 /favicon.ico        : unregistered path
2021/04/03 18:30:11 /json_intro         : https://blog.golang.org/json
2021/04/03 18:30:12 /favicon.ico        : unregistered path
2021/04/03 18:30:23 /urlshort           : https://github.com/hauntarl/go_002_urlshort
2021/04/03 18:30:25 /favicon.ico        : unregistered path
2021/04/03 18:30:30 /something          : unregistered path
2021/04/03 18:30:30 /favicon.ico        : unregistered path
exit status 3221225786
```

**Test Urls:**

- /golang : <https://github.com/hauntarl/golang>
- /gophercises : <https://courses.calhoun.io/courses/cor_gophercises>
- /yaml : <https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc>
- /yaml_github : <https://github.com/go-yaml/yaml>
- /quiz : <https://github.com/hauntarl/go_001_quiz>
- /quiz_readme : <https://github.com/hauntarl/go_001_quiz/blob/master/README.md>
- /json : <https://pkg.go.dev/encoding/json?tab=doc>
- /json_intro :<https://blog.golang.org/json>
- /urlshort : <https://github.com/hauntarl/go_002_urlshort>
- /urlshort_readme : <https://github.com/hauntarl/go_002_urlshort/blob/master/README.md>
