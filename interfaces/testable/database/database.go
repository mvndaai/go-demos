//go:generate mockgen -destination=./mocks/database.go -package=mocks -source=./database.go

package database

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type IDatabase interface {
	io.Closer // interfaces can include other interfaces
	Save(s string) error
}

type Database struct {
	filepath string
	file     *os.File
}

func NewDatabase() (IDatabase, error) {
	fp, err := filepath.Abs("./dbfile.txt")
	if err != nil {
		return nil, err
	}

	return &Database{
		filepath: fp,
	}, nil
}

func (db *Database) Close() error {
	if db.file == nil {
		return nil
	}
	return db.file.Close()
}

func (db *Database) Save(s string) error {
	if db.file == nil {
		f, err := os.Create(db.filepath)
		if err != nil {
			return fmt.Errorf("could not create file: %w", err)
		}
		db.file = f
	}

	_, err := db.file.WriteString(s + "\n")
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	err = db.file.Sync()
	if err != nil {
		return fmt.Errorf("could not sync file: %w", err)
	}
	return nil
}
