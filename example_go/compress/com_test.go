package compress

import (
	"archive/tar"
	"compress/bzip2"
	"compress/gzip"
	"os"
	"testing"
)

//go:generate tar cjvf test.bz.tar com_test.go
func TestBzip(t *testing.T) {
	f, err := os.Open("test.bz.tar")
	if err != nil {
		t.Fatal("open file failed: ", err)
	}
	defer f.Close()
	tr := tar.NewReader(bzip2.NewReader(f))

	for {
		var b []byte
		if h, err := tr.Next(); err != nil {
			t.Logf("read next failed: %v", err)
			return
		} else {
			t.Logf("read header %v", h)
			b = make([]byte, h.Size)
		}

		n, err := tr.Read(b)
		t.Log("actually size is ", n)
		t.Log(string(b))
		if err != nil {
			t.Logf("read body failed: %v", err)
			return
		}
	}
}

//go:generate tar czvf test.gz.tar com_test.go
func TestGzip(t *testing.T) {
	f, err := os.Open("test.gz.tar")
	if err != nil {
		t.Fatal("open file failed: ", err)
	}
	defer f.Close()
	gzr, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	tr := tar.NewReader(gzr)

	for {
		var b []byte
		if h, err := tr.Next(); err != nil {
			t.Logf("read next failed: %v", err)
			return
		} else {
			t.Logf("read header %v", h)
			b = make([]byte, h.Size)
		}

		n, err := tr.Read(b)
		t.Log("actually size is ", n)
		t.Log(string(b))
		if err != nil {
			t.Logf("read body failed: %v", err)
			return
		}
	}
}
