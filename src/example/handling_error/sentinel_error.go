package handling_error

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func SentinelError() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat { // zip.ErrFormat is a sentinel error
		fmt.Println("Told you so")
	}
}
