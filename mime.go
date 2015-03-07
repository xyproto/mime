// Package from retrieving mimetypes for extensions
package mime

import (
	"bytes"
	"io/ioutil"
)

type MimetypeReader struct {
	filename  string
	mimetypes map[string]string
}

func New(filename string) *MimetypeReader {
	return &MimetypeReader{filename, nil}
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

// Returns the mimetype or an empty string if no mimetype or mimetype source is found
func (m *MimetypeReader) Get(ext string) string {
	var err error
	if m.mimetypes == nil {
		m.mimetypes, err = readMimetypes(m.filename)
		if err != nil {
			return ""
		}
	}
	return m.mimetypes[ext]
}
