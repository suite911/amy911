# str
The `str` package contains some mid-level string operations.

## Imports
```go
import "github.com/amyadzuki/amystuff/str"
```

## Usage Examples
```go
package main

import (
	"fmt"

	"github.com/amyadzuki/amystuff/str"
)

func main() {
	match, tail := str.CaseHasPrefix("Hello, World!", "heLLo")
	simp := str.Simp("@Amy.Adzuki.1234!")
	fmt.Printf("%v: \"%s\"; \"%s\"\n", match, tail, simp)
	// Prints: `true: ", World!"; "amyadzuki1234"`
}
```
