package db

// Local implementation of key-value DB using bolt

// Connection is opened automatically on each call
// Buckets are created if they don't exist, when using Put

import (
	"encoding/json"
	"fmt"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

const (
	path        = "/tmp/cedana.db"
	openROPerms = 0o600
	openRWPerms = 0o400
)

type LocalDB struct{}

func NewLocalDB() DB {
	return &LocalDB{}
}

func openRO() (*bolt.DB, error) {
	return bolt.Open(path, openROPerms, &bolt.Options{ReadOnly: true})
}

func openRW() (*bolt.DB, error) {
	return bolt.Open(path, openRWPerms, nil)
}

/////////////
// Getters //
/////////////

func (db *LocalDB) Get(path [][]byte, key []byte) ([]byte, error) {
	conn, err := openRO()
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}

	var v []byte
	err = conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(path[0])
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		for i, name := range path {
			if i == 0 {
				continue
			}
			b = b.Bucket(name)
			if b == nil {
				return fmt.Errorf("bucket not found")
			}
		}

		v = b.Get(key)
		if v == nil {
			return nil
		}

		return nil
	})
	defer conn.Close()

	return v, err
}

func (db *LocalDB) GetString(path [][]byte, key []byte) (string, error) {
	v, err := db.Get(path, key)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (db *LocalDB) GetInt(path [][]byte, key []byte) (int, error) {
	v, err := db.Get(path, key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(v))
}

func (db *LocalDB) GetBool(path [][]byte, key []byte) (bool, error) {
	v, err := db.Get(path, key)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(string(v))
}

func (db *LocalDB) GetAny(path [][]byte, key []byte) (any, error) {
	v, err := db.Get(path, key)
	if err != nil {
		return nil, err
	}
	var value any
	err = json.Unmarshal(v, &value)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal value: %v", err)
	}
	return v, nil
}

/////////////
// Setters //
/////////////

// Recursively create buckets if they don't exist

func (db *LocalDB) Put(path [][]byte, key []byte, value []byte) error {
	conn, err := openRW()
	if err != nil {
		return fmt.Errorf("could not open db: %v", err)
	}

	err = conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(path[0])
		if err != nil {
			return fmt.Errorf("could not create bucket: %v", err)
		}

		for i, name := range path {
			if i == 0 {
				continue
			}
			b, err = b.CreateBucketIfNotExists(name)
			if err != nil {
				return fmt.Errorf("could not create bucket: %v", err)
			}
		}

		return b.Put(key, value)
	})
	defer conn.Close()

	return err
}

func (db *LocalDB) PutString(path [][]byte, key []byte, value string) error {
	return db.Put(path, key, []byte(value))
}

func (db *LocalDB) PutInt(path [][]byte, key []byte, value int) error {
	return db.Put(path, key, []byte(strconv.Itoa(value)))
}

func (db *LocalDB) PutBool(path [][]byte, key []byte, value bool) error {
	return db.Put(path, key, []byte(strconv.FormatBool(value)))
}

/////////////
// Listers //
/////////////

func (db *LocalDB) List(path [][]byte) ([][]byte, error) {
	conn, err := openRO()
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}

	var list [][]byte
	err = conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(path[0])
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		for i, name := range path {
			if i == 0 {
				continue
			}
			b = b.Bucket(name)
			if b == nil {
				return fmt.Errorf("bucket not found")
			}
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			list = append(list, v)
		}

		return nil
	})
	defer conn.Close()

	return list, err
}

func (db *LocalDB) ListKeys(path [][]byte) ([][]byte, error) {
	conn, err := openRO()
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}

	var list [][]byte
	err = conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(path[0])
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		for i, name := range path {
			if i == 0 {
				continue
			}
			b = b.Bucket(name)
			if b == nil {
				return fmt.Errorf("bucket not found")
			}
		}

		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			list = append(list, k)
		}

		return nil
	})
	defer conn.Close()

	return list, err
}
