package urlshort

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

var (
	DBName     string
	BucketName []byte
)

// CreateDB creates and opens a database at the given path. If the file does
// not exist then it will be created automatically.
func CreateDB(name string) (*bolt.DB, error) {
	DBName = name
	return bolt.Open(name, fs.ModePerm, nil)
}

// CreateBucket creates a new bucket if it doesn't already exist.
func CreateBucket(db *bolt.DB, name string) error {
	BucketName = []byte(name)
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(BucketName)
		return err
	})
}

// UpdateBucket writes data onto database at set BucketName,
// data needs to be a key, value pair.
func UpdateBucket(db *bolt.DB, data map[string]string) error {
	return db.Update(func(tx *bolt.Tx) error {
		for k, v := range data {
			err := tx.Bucket(BucketName).Put([]byte(k), []byte(v))
			if err != nil {
				return err
			}
			log.Printf("%-15s: %-20s, %s\n", "inserted", k, v)
		}
		return nil
	})
}

// ReadBucket tries to read associated value for the given key from database
// in the set BucketName.
func ReadBucket(db *bolt.DB, key string) (string, error) {
	var (
		url []byte
		err error
	)
	err = db.View(func(tx *bolt.Tx) error {
		url = tx.Bucket(BucketName).Get([]byte(key))
		if url == nil {
			return errors.New("unregistered path")
		}
		return nil
	})
	if err != nil {
		log.Printf("%-20s: %s", key, err)
		return "", err
	}
	log.Printf("%-20s: %s", key, url)
	return string(url), nil
}

// pathURL stores parsed data from an input file into Go understandable
// structure, upon which further operations can be performed.
type pathURL struct {
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
	URL  string `yaml:"url,omitempty" json:"url,omitempty"`
}

// buildMap converts pathURL structure into key, value pair, which is the
// accepted format by our database.
func buildMap(input []pathURL) map[string]string {
	data := make(map[string]string, len(input))
	for _, item := range input {
		data[item.Path] = item.URL
	}
	return data
}

// ReadData parses the input string using the provided format unmarshaller,
// then inserts the parsed elements into database.
func ReadData(
	input string, db *bolt.DB,
	unmarshal func([]byte, interface{}) error) error {
	var data []pathURL
	if err := unmarshal([]byte(input), &data); err != nil {
		return err
	}
	return UpdateBucket(db, buildMap(data))
}

// Decoder parses the input file and stores the parsed elements in given
// interface
type Decoder interface{ Decode(interface{}) error }

// ReadFile parses the give file using the provided format decoder,
// then inserts the parsed elements into database.
func ReadFile(
	name string, db *bolt.DB,
	newDecoder func(io.Reader) Decoder) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	var data []pathURL
	err = newDecoder(file).Decode(&data)
	if err != nil {
		return err
	}
	return UpdateBucket(db, buildMap(data))
}
