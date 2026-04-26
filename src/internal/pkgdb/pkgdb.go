package pkgdb


import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type DB struct {
	Packages []Entry `toml:"package"`
}

type Entry struct {
	Name		string		`toml:"name"`
	Version		string		`toml:"version"`
	Source		string		`toml:"source"`
	InstallDir  string		`toml:"install_dir"`
	Grovefile	string		`toml:"grovefile"`
	InstalledOn	time.Time	`toml:"installed_on"`
}

func Load(path string) (*DB, error) {
	var db DB

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &db, nil
	}

	if err != nil {
		return nil, fmt.Errorf("reading packages.toml: %w", err)
	}

	if _, err := toml.Decode(string(data), &db); err != nil {
		return nil, fmt.Errorf("parsing packages.toml: %w", err)
	}

	return &db, nil
}

func Save(path string, db *DB) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("creating db directory: %w", err)
	}

	tmp := path + ".tmp"

	f, err := os.OpenFIle(tmp. os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("opening temp db file: %w", err)
	}

	if err := toml.NewEncoder(f).Encode(db); err != nil {
		f.Close()
		os.Remove(tmp)
		return fmt.Errorf("encoding packages.toml: %w", err)
	}

	f.Close()
	
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return fmt.Errorf("saving packages.toml: %w", err)
	}

	return nil
}



























