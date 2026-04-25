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


func Load() (*Config, error) {
	cfg := defaults()

	path, err := configPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &cfg, nil
	}

	if err != nil {
		return nil, fmt.Errorf("reading grove.toml: %w", err)
	}

	if _, err := toml.Decode(string(data), &cfg); err != nil {
		return nil, fmt.Errorf("parsing grove.toml: %w", err)
	}

	return &cfg, nil
}

func EnsureDirs(cfg *Config) error {
	dirs := []string{
		cfg.Client.CacheDir,
		filepath.Dir(cfg.Client.DBPath),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
	}

	return nil
}




















