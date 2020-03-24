package archive

import (
	"archive/zip"
	"io"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	f, _ := os.Create("test.zip")
	defer f.Close()
	zw := zip.NewWriter(f)
	defer zw.Close()
	ziw, err := zw.Create("test_file1")
	if err != nil {
		t.Fatal(err)
	}
	ziw.Write([]byte("hello, world!"))

	zw.Flush()
}

//go:generate zip zip_test.zip zip_test.go
func TestReadZip(t *testing.T) {
	r, err := zip.OpenReader("zip_test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			t.Fatal(err)
		}
		io.Copy(os.Stdout, rc)
	}
}
