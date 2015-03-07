# Mime [![GoDoc](https://godoc.org/github.com/xyproto/mime?status.svg)](http://godoc.org/github.com/xyproto/mime)

<!--#Mime [![Build Status](https://travis-ci.org/xyproto/mime.svg?branch=master)](https://travis-ci.org/xyproto/mime) [![Build Status](https://drone.io/github.com/xyproto/mime/status.png)](https://drone.io/github.com/xyproto/mime/latest) [![GoDoc](https://godoc.org/github.com/xyproto/mime?status.svg)](http://godoc.org/github.com/xyproto/mime)-->

Package for retrieving the mime type given an extension.

Features and limitations
------------------------

* Must be given a filename that contains a list of mimetypes followed by extensions. Typically `/etc/mime.types`.
* Will only read the file once, then store the lookup table in memory. This results in fast lookups.

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

* Will output: `image/svg+xml`


General information
-------------------

* Version: 0.1
* License: MIT
* Alexander F Rødseth

