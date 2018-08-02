# onfail
The `onfail` package allows specifiable behavior when an error is encountered.

## Imports
```go
import "github.com/amyadzuki/amystuff/onfail"
```

## Usage Examples
```go
package main

import "github.com/amyadzuki/amystuff/onfail"

func gowrong(onFail ...onfail.OnFail) {
	onfail.Fail("Oh noes, something went wrong! >~<", nil, onfail.Fatal, onFail...)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			gowrong() // log.Fatalln(err)
		}
	}
	gowrong(onfail.Panic) // panic(err)
}
```
