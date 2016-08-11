package ldb

import (
	"log"
	"os"

	"github.com/bmatsuo/lmdb-go/lmdb"
)

var env *lmdb.Env

// Ldb struct
type Ldb struct {
	Name string
	dbi  lmdb.DBI
}

func openDatabase(dir string, name string, dbi *lmdb.DBI, flags uint) error {
	var err error
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModeDir|os.ModePerm)
		if err != nil {
			log.Print("Create fail")
			return err
		}
		log.Println("Create db dir: ", dir)
	}

	if env == nil {
		env, err = lmdb.NewEnv()
		if err != nil {
			return err
		}

		if err = env.SetMapSize(1 << 20); err != nil {
			return err
		}

		// Must set, but need to fix value.
		if err = env.SetMaxDBs(4); err != nil {
			return err
		}

		if err = env.Open(dir, 0, 0600); err != nil {
			return err
		}
	}

	err = env.Update(func(txn *lmdb.Txn) (err error) {
		*dbi, err = txn.OpenDBI(name, lmdb.Create|flags)
		return err
	})
	return err
}

// NewDB new one
func NewDB(name string, dir string, flags uint) *Ldb {
	var dbi lmdb.DBI
	if err := openDatabase(dir, name, &dbi, flags); err != nil {
		log.Panicln(err)
	}
	return &Ldb{Name: name, dbi: dbi}
}

// Put data.
func (l Ldb) Put(key []byte, val []byte) error {
	return env.Update(func(txn *lmdb.Txn) (err error) {
		return txn.Put(l.dbi, key, val, 0)
	})
}

// Get data
func (l Ldb) Get(key []byte) ([]byte, error) {
	var val []byte
	var err error

	err = env.View(func(txn *lmdb.Txn) (err error) {
		val, err = txn.Get(l.dbi, key)
		return err
	})
	if err != nil && !lmdb.IsNotFound(err) {
		return nil, err
	}
	return val, err
}

// CGet is Cursor get
func (l Ldb) CGet(setkey, setval []byte, data chan []byte) {
	var cur *lmdb.Cursor

	//data = make([][]byte, 0)

	env.View(func(txn *lmdb.Txn) (err error) {
		cur, err = txn.OpenCursor(l.dbi)
		if err != nil {
			return err
		}

		go func(cur *lmdb.Cursor) {
			for {
				_, val, err := cur.Get(setkey, setval, lmdb.NextNoDup)

				if err != nil {
					if !lmdb.IsNotFound(err) {
						log.Println(err)
					}
					close(data)
					break
				} else {
					data <- val
				}
				for {
					_, val, err := cur.Get(setkey, setval, lmdb.NextDup)
					if lmdb.IsNotFound(err) {
						log.Println(err)
						break
					} else if err == nil {
						data <- val
					}
				}
			}
		}(cur)
		return err
	})
}
