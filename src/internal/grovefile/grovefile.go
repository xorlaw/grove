package grovefile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

const Filename = "Grovefile"

type Grovefile struct {
	Package 	PackageMeta `toml:"package"`
	Source		Source		`toml:"source"`
	Build		Build		`toml:"build"`
	Deps		Deps		`toml:"deps"`
}

type PackageMeta struct {
	Name		string `toml:"name"`
	Version 	string `toml:"version"`
	Desc		string `toml:"desc"`
	Author		string `toml:"author"`
	License		string `toml:"license"`
}

type Source struct {
	Primary		string `toml:"primary"`
	Fallback 	string `toml:"fallback"`
}

type Build struct {
	Cmd			string `toml:"cmd"`
	Install 	string `toml:"install"`
	Prefix		string `toml:"prefix"`
}

type Deps struct {
	Requires	[]string `toml:"requires"`
}


func Parse(path string) (*Grovefile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading Grovefile: %w", err)
	}

	var gf Grovefile
	if _, err := toml.Decode(string(data), &gf); err != nil {
		return nil, fmt.Errorf("parsing Grovefile: %w", err)
	}

	return &gf, nil
}

func FindAndParse(dir string) (*Grovefile, error) {
	path := filepath.Join(dir, Filename)
	return Parse(path)
}










