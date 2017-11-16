package mime

import (
	"github.com/bmizerany/assert"

	"testing"
)

var m = New("/etc/mime.types", true)

var mimeTestTable = map[string]string{
	"tar.xz": "application/x-xz-compressed-tar",
	"xz": "application/x-xz",
	"tar": "application/x-tar",
	"tgz": "application/x-gtar-compressed",
	"gz": "application/x-gzip",
	"tbz2": "application/x-bzip-compressed-tar",
	"bz2": "application/x-bzip2",
}

func TestSVG(t *testing.T) {
	assert.Equal(t, m.Get("svg"), "image/svg+xml")
}

func TestTar(t *testing.T) {
	for ext, mimeType := range mimeTestTable {
		assert.Equal(t, m.Get(ext), mimeType)
	}
}

