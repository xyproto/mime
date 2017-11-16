package mime

import (
	"github.com/bmizerany/assert"

	"testing"
)

var m = New("/etc/mime.types", true)

func TestSVG(t *testing.T) {
	assert.Equal(t, m.Get("svg"), "image/svg+xml")
}

func TestTarXZ(t *testing.T) {
	assert.Equal(t, m.Get("tar.xz"), "application/x-gtar-compressed")
}

func TestXZ(t *testing.T) {
	assert.Equal(t, m.Get("xz"), "application/x-xz")
}

func TestTar(t *testing.T) {
	assert.Equal(t, m.Get("tar"), "application/x-tar")
}

func TestTgz(t *testing.T) {
	assert.Equal(t, m.Get("tgz"), "application/x-gtar-compressed")
}

func TestGz(t *testing.T) {
	assert.Equal(t, m.Get("gz"), "application/x-gzip")
}

func TestBz2(t *testing.T) {
	assert.Equal(t, m.Get("bz2"), "application/x-bzip2")
}

func TestTbz2(t *testing.T) {
	assert.Equal(t, m.Get("tbz2"), "application/x-gtar-compressed")
}
