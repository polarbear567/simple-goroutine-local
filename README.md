Simple-goroutine-local is a cross goroutine storage tool with very simple implementation and function (the concept is similar to Java ThreadLocal).

Getting Started
===============

## Installing

To start using GJSON, install Go and run `go get`:

```sh
$ go https://github.com/polarbear567/simple-goroutine-local
```

This will retrieve the library.

## How to use

```go
package main

import (
    "fmt"
    sgl "github.com/polarbear567/simple-goroutine-local"
    "sync"
)

func main() {
    gl := sgl.NewGoRoutineLocal()
    var wg sync.WaitGroup
    wg.Add(5)
    gl.Set("value", "main")
    for i := 0; i < 5; i++ {
        go func(i int) {
            defer func() {
                wg.Done()
                gl.Del(i)
            }()
            gl.Set("value", i)
            v, _ := gl.Get("value")
            fmt.Printf("i: %d, v: %d\n", i, v)
        }(i)
    }
    wg.Wait()
    v, _ := gl.Get("value")
    fmt.Printf("main, value: %s\n", v)
    gl.Del("value")
    v, ok := gl.Get("value")
    fmt.Printf("main, value after delete: %v, value exist: %t\n", v, ok)
    gl.DelMap()
    if v, ok = gl.Get("value"); !ok {
        fmt.Printf("main, map exist: %t\n", ok)
    }
}
```

The output:

```
i: 4, v: 4
i: 2, v: 2
i: 1, v: 1
i: 0, v: 0
i: 3, v: 3
main, value: main
main, value after delete: <nil>, value exist: false
main, map exist: false
```

### Tip
Since goroutine may be reused, please make sure to explicitly set in the current goroutine before get, or delete the key before the end of the current goroutine.

