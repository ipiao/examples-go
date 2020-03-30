package crypto

import (
	"crypto"
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestCrypto(t *testing.T) {
	md4 := crypto.MD4
	if md4.Available() {
		t.Log(string(md4.New().Sum([]byte("12345678"))))
	} else {
		t.Log("md4 not avaliable")
	}
}

//
func TestMD5(t *testing.T) {
	b := md5.Sum([]byte("123456"))
	t.Log(hex.EncodeToString(b[:]))
}
