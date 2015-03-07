package main

import (
	"fmt"

	"github.com/xyproto/mime"
)

func main() {
    fmt.Println(mime.New("/etc/mime.types").Get("svg"))
}

