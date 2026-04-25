package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Client		Client		`toml:"client"`
	Sources		[]Source 	`toml:"source"`
}

type Client struct {
	InstallDir	string		`toml:"install_dir"`
	DBPath		string		`toml:"db_path"`
	CacheDir	string		`toml:"cache_dir"`
}

type Source struct {
	Name		string		`toml:"name"`
	URL			string		`toml:"url"`
	Type 		string		`toml:"type"` // grove or github
}


