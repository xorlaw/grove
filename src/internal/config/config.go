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

func defaults() Config {
	home, _ := os.UserHomeDir()
	return Config{
		Client: Client{
			InstallDir:	"/usr/local",
			DBPath:		filepath.Join(home, ".local", "share", "grove", "packages.toml"),
			CacheDir:	filepath.Join(home, ".cache", "grove"),
		},
	}
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("finding home directory: %w", err)
	}
	return filepath.Join(home, ".config", "grove", "grove.toml"), nil
}


