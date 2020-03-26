package bufio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBufReader(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	b, _ := br.Peek(5)
	t.Log(string(b))
	b = append(b, 'o') // underlying bytes maybe changed by append

	br.Discard(5)
	line, _, _ := br.ReadLine()
	t.Log(string(line))
	rb := make([]byte, 7)
	_, _ = br.Read(rb)
	t.Log(string(rb))

	t.Log(br.Buffered())
}

func TestBufferWrite(t *testing.T) {
	bw := bufio.NewWriterSize(os.Stdout, 5)

	bw.WriteString("hello\n")
	bw.WriteString("world\n")
	fmt.Println("after bf write")
	time.Sleep(time.Second)
	bw.Flush()
}

func TestScanner(t *testing.T) {
	buf := strings.NewReader("hello world")
	sc := bufio.NewScanner(buf)
	sc.Split(bufio.ScanBytes)

	for sc.Scan() {
		t.Log(sc.Text())
	}
}
