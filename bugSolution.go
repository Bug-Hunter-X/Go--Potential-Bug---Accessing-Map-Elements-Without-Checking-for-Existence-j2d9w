```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int)
    m["a"] = 1
    value, ok := m["a"]
    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("Key not found")
    }

    value, ok = m["b"]
    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("Key not found")
    }
}
```