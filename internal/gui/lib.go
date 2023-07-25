package gui

// addBucket add a new bucket
func addBucket(bucket string) error {
	return db.Create(bucket)
}

// deleteBucket deletes a bucket
func deleteBucket(bucket string) error {
	return db.Delete(bucket)
}

// deleteKey deletes key from a bucket
func deleteKey(bucket, key string) error {
	return db.DeleteKey(bucket, key)
}

// updateKey add/update key new value
func updateKey(bucket, key, value string) error {
	return db.UpdateKey(bucket, key, []byte(value))
}

// updateBucket update a bucket name
func updateBucket(bucket, newName string) error {
	// Create new bucket with new name
	err := db.Create(newName)
	if err != nil {
		return err
	}

	// Iterate over old bucket keys and recreate them in new bucket
	kv, err := db.IterateKeyValues(bucket)
	if err != nil {
		return err
	}

	for _, k := range kv {
		err := db.UpdateKey(newName, k.Name, []byte(k.Value))
		if err != nil {
			return err
		}
	}

	// Delete old bucket
	err = db.Delete(bucket)
	if err != nil {
		return err
	}

	return nil
}

// bindAllBuckets retrieve all buckets from db and bind to fyne data.Binding
func bindAllBuckets() {
	// Bind buckets list of byte string
	bs, err := db.IterateBucket()
	if err != nil {
		errorHandler(err)
	}
	buckets.Set(bytesArrToStringArr(bs))

	// Clear key item list value
	selKey = ""
	keys.Set([]string{})
}

// bindAllKeys retrieve all keys in a bucket from db and bind to fyne data.Binding
func bindAllKeys(bucket string) {
	// Get bucket key list and bind
	ks, err := db.IterateKey(bucket)
	if err != nil {
		errorHandler(err)
	}
	selBucket = bucket
	keys.Set(bytesArrToStringArr(ks))

	// clear selected key value
	selKey = ""
	keyItemList.UnselectAll()
}

// bytesArrToStringArr converts array of bytes array to string array
func bytesArrToStringArr(b [][]byte) []string {
	s := []string{}

	for _, item := range b {
		s = append(s, string(item))
	}

	return s
}
