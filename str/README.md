# str
The `str` package contains some mid-level string operations.

## Imports
```go
import "github.com/amy911/amy911/str"
```

## Usage Examples
```go
package main

import (
	"fmt"

	"github.com/amy911/amy911/str"
)

func main() {
	match, tail := str.CaseHasPrefix("Hello, World!", "heLLo")
	simp := str.Simp("@Amy.Adzuki#1234!")
	fmt.Printf("%v: \"%s\"; \"%s\"\n", match, tail, simp)
	// Prints: `true: ", World!"; "amy.adzuki1234"`
}
```
