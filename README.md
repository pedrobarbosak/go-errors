# go-errors
Golang Errors with line information

```go
package main

import (
	"log"

	"github.com/pedrobarbosak/go-errors"
)

func main() {
	err := errors.New("something failed with some variables:", true, 1)
	log.Println(err)
}
```

```go
example/main.go:11 @ fn:main # something failed with some variables: true 1
```