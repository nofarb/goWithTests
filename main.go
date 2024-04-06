package main

import (
    "fmt"
    "go-sample-app/pkg"
)

func main() {
    result := pkg.Add(2, 3)
    fmt.Printf("The result is: %d\n", result)
}
