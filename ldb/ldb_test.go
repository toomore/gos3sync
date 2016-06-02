package ldb

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLdbNew(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "ldb_test")
	defer os.RemoveAll(tempDir)
	t.Log("Temp: ", tempDir)
	if err == nil {
		db := NewDB("test", tempDir)
		t.Log(db.Name)
	} else {
		t.Log(err)
	}
}

func TestLdbGet(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "ldb_test")
	defer os.RemoveAll(tempDir)
	t.Log("Temp: ", tempDir)
	if err == nil {
		db := NewDB("test", tempDir)
		db.Put([]byte("Toomore"), []byte("123"))
		val, err := db.Get([]byte("Toomore"))
		t.Log(val, err)
	} else {
		t.Log(err)
	}
}
