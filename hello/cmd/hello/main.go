package main

import (
	"os"

	"github.com/devrewoh/tpg-tools/hello"
)

func main() {
	hello.PrintTo(os.Stdout) // prints hello message to stdout
}
