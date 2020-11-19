# Log-Http

This is a powerful logger for services.

### Installation

Install and start the server.

```sh
$ go run main.go
$ go get github.com/gerardoCD/log-http-go 
```

Import in your project

```goland

import (
	lgs "github.com/gerardoCD/log-http-go"
)

```


# Features!

  - Mark init service
  - Mark end service
  - Reponse Services
  - Reponde middleware service 
  
 # Examples!

```goland


func main() {
	lgs.Begin("Name Service")
	lgs.Request("{\"hola\":4}")
	lgs.ReponseService(201,"{\"hola\":4}", "http://app/getElement", "Get Element")
	lgs.Reponse(200,"{\"hola\":4}")
}
  
```
