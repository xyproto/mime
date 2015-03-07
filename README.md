# Mime

<!--#Mime [![Build Status](https://travis-ci.org/xyproto/mime.svg?branch=master)](https://travis-ci.org/xyproto/mime) [![Build Status](https://drone.io/github.com/xyproto/mime/status.png)](https://drone.io/github.com/xyproto/mime/latest) [![GoDoc](https://godoc.org/github.com/xyproto/mime?status.svg)](http://godoc.org/github.com/xyproto/mime)-->

Package for retrieving the mime type given an extension.

Online API Documentation
------------------------

[godoc.org](http://godoc.org/github.com/xyproto/mime)


Example
-------

~~~ go
package main

import (
	"fmt"

	"github.com/xyproto/mime"
)

func main() {
	m := mime.New("/etc/mime.types")
	fmt.Println(m.Get("svg"))
}
~~~

Outputs:

image/svg+xml


General information
-------------------

* Version: 0.1
* License: MIT
* Alexander F Rødseth

