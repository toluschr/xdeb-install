package xdeb

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(path string, url string) (string, error) {
	client := &http.Client{}
	resp, err := client.Get(url)

	if err != nil {
		return "", fmt.Errorf("Could not download file %s", url)
	}

	finalUrl := resp.Request.URL.String()
	resp, err = client.Get(finalUrl)

	if err != nil {
		return "", fmt.Errorf("Could not download file %s", url)
	}

	defer resp.Body.Close()

	err = os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return "", fmt.Errorf("Could not create path %s", path)
	}

	fullPath := filepath.Join(path, filepath.Base(finalUrl))
	out, err := os.Create(fullPath)

	if err != nil {
		return "", fmt.Errorf("Could not create file %s", fullPath)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return fullPath, err
}
