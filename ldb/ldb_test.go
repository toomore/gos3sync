package ldb

import (
	"bytes"
	"io/ioutil"
	"log"
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

func TestLdbGet_two(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "ldb_test")
	defer os.RemoveAll(tempDir)
	t.Log("Temp: ", tempDir)
	if err == nil {
		db1 := NewDB("test1", tempDir)
		db1.Put([]byte("db1"), []byte("val1"))

		db2 := NewDB("test2", tempDir)
		db2.Put([]byte("db2"), []byte("val2"))

		if val1, err := db1.Get([]byte("db1")); err != nil {
			log.Fatalln("err: ", err)
		} else {
			log.Println(bytes.Equal(val1, []byte("val1")))
		}
		if val2, err := db2.Get([]byte("db2")); err != nil {
			log.Fatalln("err: ", err)
		} else {
			log.Println(bytes.Equal(val2, []byte("val2")))
		}
	}
}
