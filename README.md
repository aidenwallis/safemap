# safemap

A dead simple hashmap implementation that allows you to not worry about concurrent access/writes through the use of RWMutexes.

## Why use a lock?

This package is intended to just be a very simple and easy-to-use utility, rather than being a complex solution that doesn't use locks etc. The aim is to keep this package as bearbones as possible and just provide a simple interface for using the hashmap.

While many other solutions can be more performant, locks for many use cases are completely fine for creating a thread-safe implementation of a hashmap without experiencing any serious performance issues at fairly decent scale.


## Using the implementation

This package is incredibly easy to use, like so

```go
package main

import (
	"fmt"

	"github.com/aidenwallis/safemap"
)

func main() {
	hashmap := safemap.New()

	hashmap.Set("key", "value")

	value, ok := hashmap.Get("key")
	if ok {
		// This key exists in the hashmap.
		fmt.Println("Key value: " + value.(string))
	}

	// Delete a key from the hashmap.
	hashmap.Delete("key")

	exists := hashmap.Has("Key")
	if !exists {
		fmt.Println("Key no longer exists in hashmap")
	}
    
    count := hashmap.Count()
    fmt.Println("Only ", count, " values are in the hashmap!")
}
```
