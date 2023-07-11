package core

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type Database struct {
	db *bolt.DB
}

// New create a connection to the bolt key-value database
func New(dbPath string) (*Database, error) {
	db, err := bolt.Open(dbPath, 0666, nil)
	if err != nil {
		return nil, fmt.Errorf("database open error %s", err.Error())
	}

	return &Database{db: db}, nil
}

// Interate over bolt buckets
func (d Database) IterateBucket() ([][]byte, error) {
	var res [][]byte
	err := d.db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			res = append(res, name)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create new bolt bucket
func (d Database) Create(bucketName string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket error %s", err.Error())
		}
		return nil
	})
}

// Delete bolt bucket
func (d Database) Delete(bucketName string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))

		if err != nil {
			return fmt.Errorf("delete bucket error %s", err.Error())
		}
		return nil
	})
}

// Iterate over keys in a bolt bucket
func (d Database) IterateKey(bucketName string) ([][]byte, error) {
	var res [][]byte
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			res = append(res, v)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update key to bolt bucket
func (d Database) UpdateKey(bucketName string, key string, value []byte) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		err := b.Put([]byte(key), value)
		if err != nil {
			return fmt.Errorf("add key error %s", err.Error())
		}
		return nil
	})
}

// Get key from bolt bucket
func (d Database) GetKey(bucketName string, key string) ([]byte, error) {
	var res []byte
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		v := b.Get([]byte(key))
		if v == nil {
			return fmt.Errorf("key %s not found", key)
		}

		res = make([]byte, len(v))
		copy(res, v)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete key from bolt bucket
func (d Database) DeleteKey(bucketName string, key string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		err := b.Delete([]byte(key))
		if err != nil {
			return fmt.Errorf("delete key error %s (possible deleting a key from a bucket created in read-only tx)", err.Error())
		}
		return nil
	})
}
