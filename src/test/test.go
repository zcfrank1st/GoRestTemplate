package main

import (
    "fmt"
)

func main() {
    var ss interface{}
    ss = []string{"a", "b"}


    if r, ok := ss.([]string); ok {
        for i, v :=range r {
            fmt.Println(i)
            fmt.Println(v)
        }
    }
}
