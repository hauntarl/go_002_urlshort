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

func CreateDB(name string) (*bolt.DB, error) {
	DBName = name
	return bolt.Open(name, fs.ModePerm, nil)
}

func CreateBucket(db *bolt.DB, name string) error {
	BucketName = []byte(name)
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(BucketName)
		return err
	})
}

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

type pathURL struct {
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
	URL  string `yaml:"url,omitempty" json:"url,omitempty"`
}

func buildMap(input []pathURL) map[string]string {
	data := make(map[string]string, len(input))
	for _, item := range input {
		data[item.Path] = item.URL
	}
	return data
}

func ReadData(
	input string, db *bolt.DB,
	unmarshal func([]byte, interface{}) error) error {
	var data []pathURL
	if err := unmarshal([]byte(input), &data); err != nil {
		return err
	}
	return UpdateBucket(db, buildMap(data))
}

type Decoder interface{ Decode(interface{}) error }

func parseFile(file io.Reader, dec Decoder) (data []pathURL, err error) {
	err = dec.Decode(&data)
	return
}

func ReadFile(
	name string, db *bolt.DB,
	newDecoder func(io.Reader) Decoder) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := parseFile(file, newDecoder(file))
	if err != nil {
		return err
	}
	return UpdateBucket(db, buildMap(data))
}
