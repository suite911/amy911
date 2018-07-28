# onfail
The `onfail` package allows specifiable behavior when an error is encountered.

## Imports
```go
import "github.com/amyadzuki/amystuff/onfail"
```

## Usage Examples
```go
package main

import (
	"errors"
	"github.com/amyadzuki/amystuff/onfail"
)

func gowrong(onFail onfail.OnFail) {
	onFail.Fail(errors.New("Oh noes, something went wrong! >~<"))
}

func main() {
	gowrong(onfail.Panic)
}
```
