package archive

import (
	"archive/tar"
	"os"
	"testing"
	"time"
)

func TestTar(t *testing.T) {
	f, err := os.Create("test.tar")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	tw := tar.NewWriter(f)
	defer tw.Close()
	tw.WriteHeader(&tar.Header{
		Typeflag:   tar.TypeReg,
		Name:       "test.tar",
		Linkname:   "",
		Size:       20,
		Mode:       0,
		Uid:        0,
		Gid:        0,
		Uname:      "",
		Gname:      "",
		ModTime:    time.Now(),
		AccessTime: time.Now(),
		ChangeTime: time.Now(),
		Devmajor:   0,
		Devminor:   0,
		Xattrs:     nil,
		PAXRecords: nil,
		Format:     0,
	})
	i, err := tw.Write([]byte("hello, world!"))
	if err != nil {
		t.Fatalf("write err %v\n", err)
	}
	if i != 13 {
		t.Fatalf("invalid of number, want 13, got %d\n", i)
	}
	tw.Flush()
}

//go:generate tar cvf tar_test.tar tar_test.go
func TestReadTar(t *testing.T) {
	f, err := os.Open("tar_test.tar")
	if err != nil {
		t.Fatal("open file failed: ", err)
	}
	defer f.Close()
	tr := tar.NewReader(f)

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
