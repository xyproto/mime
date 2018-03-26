// Package mime helps retrieving mimetypes given extensions.
// This is an alternative to the "mime" package, and has fallbacks for the most common types.
package mime

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sync"
)

var fallback = map[string]string{
	"tar.gz":  "application/x-gzip-compressed-tar",
	"tar.bz":  "application/x-bzip-compressed-tar",
	"tar.bz2": "application/x-bzip-compressed-tar",
	"tar.xz":  "application/x-xz-compressed-tar",
	"tgz":     "application/x-gzip-compressed-tar",
	"tbz":     "application/x-bzip-compressed-tar",
	"tbz2":    "application/x-bzip-compressed-tar",
	"txz":     "application/x-xz-compressed-tar",
	"gz":      "application/x-gzip",
	"bz2":     "application/x-bzip2",
	"xz":      "application/x-xz",
	"html":    "text/html",
	"css":     "text/css",
	"js":      "application/javascript",
	"txt":     "text/plain",
	"png":     "image/png",
	"jpg":     "image/jpg",
	"json":    "application/javascript",
	"svg":     "image/svg+xml",
	"xml":     "text/xml",
	"rss":     "application/rss+xml",
	"zip":     "application/zip",
	"tar":     "application/x-tar",
}

// Reader caches the contents of a mime info text file
type Reader struct {
	filename  string
	utf8      bool
	mimetypes map[string]string
	mu        sync.Mutex
}

// New creates a new Reader. The filename is a list of mimetypes and extensions.
// If utf8 is true, "; charset=utf-8" will be added when setting http headers.
func New(filename string, utf8 bool) *Reader {
	return &Reader{filename, utf8, nil, sync.Mutex{}}
}

// Read a mimetype text file. Return a hash map from ext to mimetype.
func readMimetypes(filename string) (map[string]string, error) {
	mimetypes := make(map[string]string)
	// Read the mimetype file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// For each line, store extensions and mimetypes in the hash map
	for _, line := range bytes.Split(data, []byte("\n")) {
		fields := bytes.Fields(line)
		if len(fields) > 1 {
			for _, ext := range fields[1:] {
				mimetypes[string(ext)] = string(fields[0])
			}
		}
	}
	return mimetypes, nil
}

// Get returns the mimetype, or an empty string if no mimetype or mimetype source is found
func (mr *Reader) Get(ext string) string {
	var err error
	// No extension, suggest text/plain (README, LICENSE etc)
	if len(ext) == 0 {
		return "text/plain"
	}
	// Strip the leading dot
	if ext[0] == '.' {
		ext = ext[1:]
	}
	mr.mu.Lock()
	defer mr.mu.Unlock()
	if mr.mimetypes == nil {
		mr.mimetypes, err = readMimetypes(mr.filename)
		if err != nil {
			// Using the fallback hash map
			if mime, ok := fallback[ext]; ok {
				return mime
			}
			// Unable to find the mime type for the given extension
			return ""
		}
	}
	// Use the value from the hash map
	if mime, ok := mr.mimetypes[ext]; ok {
		return mime
	}
	// Using the fallback hash map
	if mime, ok := fallback[ext]; ok {
		return mime
	}
	// Unable to find the mime type for the given extension
	return ""
}

// SetHeader sets the Content-Type for a given ResponseWriter and filename extension
func (mr *Reader) SetHeader(w http.ResponseWriter, ext string) {
	mimestring := mr.Get(ext)
	if mimestring == "" {
		// Default mime type
		mimestring = "application/octet-stream"
	}
	if mr.utf8 {
		mimestring += "; charset=utf-8"
	}
	w.Header().Add("Content-Type", mimestring)
}
