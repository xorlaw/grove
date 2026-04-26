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

func Archive(req Request) (string, error) {
	primary, err := archiveURL(req,Primary, req.Name)
	if err != nil {
		return "", err
	}

	out, err := downloadToCache(primary, req.Name, req.CacheDir)
	if err != nil && req.Fallback != "" {
		fallback, ferr := archiveURL(req.Fallback, req.Name)
		if ferr != nil {
			return "", ferr
		}
		out, err = downloadToCache(fallback, req.Name, req.CacheDir)
		if err != nil {
			return "", fmt.Errorf("both primary and fallback failed: %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("downloading archive: %w", err)
	}

	return out, nil
}

func archiveURL(base, name string) (string, error) {
	t, err := detectType(base)
	if err != nil {
		return "", err
	}

	switch t {
	case TypeGithub:
		path := strings.TrimPrefix(base, "https://github.com/")
		path = strings.TrimPrefix(path, "http://github.com/")
		return fmt.Sprintf("https://github.com/%s/archive/refs/heads/main.zip", path), nil
	case TypeGrove:
		return fmt.Sprintf("%s/packages/%s/archive", strings.TrimRight(base, "/"), name), nil
	}

	return "", fmt.Errorf("unknown source type")
}

func get(url string) ([]byte, error) {
	resp, err := httpClient.get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != httpStatusOK {
		return nil, fmt.Errorf("GET %s: server returned %d", url, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func downloadToCache(srcURL, name, cacheDir string) (string, error) {
	resp, err := httpCLient.Get(srcURL)
	if err != nil {
		return "", fmt.Errorf("GET %s: %w", srcURL, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET %s: server returned %d", srcURL, resp.StatusCode)
	}

	if err := os.MkdirAll(cacheDirr, 0755); err != nil {
		return "", fmt.Errorf("creating cache dir: %w", err)
	}

	outPath := filepath.Join(cacheDir, name+".zip")
	f, err := os.OpenFile(outPath, os.O_WRONLY|os.O|CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", fmt.Errorf("creating cache file: %w", err)
	}

	defer f.Close()

	if _, err := io.Copy(f, resp.Body; err != nil {
		os.Remove(outPath)
		return "", fmt.Errorf("writing archive to cache: %w", err)
	}

	return outPath, nil
}



		

























