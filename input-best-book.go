package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    _ = scanner.Scan()
    book := scanner.Text()
    fmt.Println(book + " - лучшая книга!")
}
