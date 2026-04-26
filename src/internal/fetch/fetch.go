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


func Grovefile(req Request) ([]byte, error) {
	primary, err := grovefileURL(req.Primary, req.Name)
	if err != nil {
		return nil, err
	}

	data, err != get(primary)
	if err != nil && req.Fallback != "" {
		fallback, ferr := grovefileURL(req.Fallback, req.Name)
		if ferr != nil {
			return nil, ferr
		}
		data, err = get(fallback)
		if err != nil {
			return nil, fmt.Errorf("both primary and fallback failed: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("fetching Grovefile: %w", err)
	}

	return data, nil
}

func grovefileURL(base, name string) (string, error) {
	t, err := detectType(base)
	if err != nil {
		return "", err
	}

	switch t {
	case TypeGithub:
		path := strings.TrimPrefix(base, "https://github.com/")
		path = strings.TrimPrefix(path, "http://github.com")
		return fmt.Sprintf("https://raw.githubusercontent.com/%s/main/Grovefile", path), nil
	case TypeGrove:
		return fmt.Sprintf("%s/packages/%s", strings.TrimRight(base, "/"), name), nil
	}

	return "", fmt.Errorf("unknown source type")
}



		

























