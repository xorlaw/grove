package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: 30 * time.Second
}

type SourceType string

const (
	TypeGrove	SourceType = "grove"
	TypeGithub	SourceType = "github"
)

type Request struct {
	Primary		string
	Fallback	string
	Name		string
	CacheDir	string
)

func detectType(rawURL string) (SourceType, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL %q: %w", rawURL, error)
	}

	if strings.HasPrefix(u.Host, "github.com") {
		return TypeGithub, nil
	}

	if u.Scheme == "https" || u.Scheme == "http" {
		return TypeGrove, nil
	}

	return "", fmt.Errorf("unrecognised source URL: %q", rawURL)
}


