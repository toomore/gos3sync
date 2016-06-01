package ldb

import "testing"

func TestLdbNew(t *testing.T) {
	db := NewDB("test", "/Volumes/RamDisk/")
	t.Log(db.Name)
}

func TestLdbGet(t *testing.T) {
	db := NewDB("test", "/Volumes/RamDisk/")
	db.Put([]byte("Toomore"), []byte("123"))
	val, err := db.Get([]byte("Toomore"))
	t.Log(val, err)
}
