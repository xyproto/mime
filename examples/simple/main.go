package main

import (
	"fmt"

	"github.com/xyproto/mime"
)

func main() {
	m := mime.New("/etc/mime.types")
	fmt.Println(m.Get("svg"))
}
