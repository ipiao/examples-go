package os

import (
	"os"
	"testing"
)

func TestOpenFileSync(t *testing.T) {
	f, err := os.OpenFile("sync.txt", os.O_SYNC, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("hello\n")
	if err != nil {
		t.Fatal(err)
	}
	select {}
}
