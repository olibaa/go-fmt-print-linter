# go-fmt-print-linter
`go-fmt-print-linter` is a static analysis tool to detect `fmt.Print` or `fmt.Println` or `fmt.Printf`.


```go

package main

import "fmt"

func main() {
	fmt.Print("Hello, world!") // detect this

	fmt.Println("Hello, world!") // detect this

	fmt.Printf("Hello, world!") // detect this
}

```


You can add an ignore comment so that 'go-fmt-print-linter' is not detected

```go

package main

import "fmt"

func main() {
    //lint:ignore go_fmt_print_linter ignore
	fmt.Print("Hello, world!") // don't detect this by ignore comment
}

