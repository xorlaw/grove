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


