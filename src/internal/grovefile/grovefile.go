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


