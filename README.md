# Url Shortener

Code an http.Handler that forwards paths to other URLs (similar to Bitly).

Implementation of URL Shortener from gophercises, including the bonus section.

**[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun

**Run Commands:**

- go run main\main.go
- go run main\main.go -h (--help) (to get information regarding flags)
- go run main\main.go --yaml file-name.yaml -json=file-name.json

**Test Urls:**

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

**Test Output:**

                Flags set...
                        yaml = yaml_urls.yaml
                        json = json_urls.json

                Database initialized...

                Fallback http.HandleFunc created...

                From MapHandler...
                        Inserted /golang...
                        Inserted /gophercises...

                From YamlHandler...
                        Inserted /yaml...
                        Inserted /yaml_github...

                From YamlFileHandler...
                        Inserted /quiz...
                        Inserted /quiz_readme...

                From JSONHandler...
                        Inserted /json...
                        Inserted /json_intro...

                From JSONFileHandler...
                        Inserted /urlshort...
                        Inserted /urlshort_readme...

                Starting the server on localhost:8080

                /                    : path not registered
                /favicon.ico         : path not registered
                /golang              : https://github.com/hauntarl/golang
                /favicon.ico         : path not registered
                /gophercises         : https://courses.calhoun.io/courses/cor_gophercises
                /favicon.ico         : path not registered

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
- gopkg.in/yaml.v2 - to work with yaml data
- flag - to get yaml/json file name
- os - to open and close the file
- io - to read from file whic satisfies io.Reader interface
- encoding/json - to work with json data
- github.com/boltdb/bolt - to store and retrieve urls for specified path
