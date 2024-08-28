package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url string) ([]byte, error) {
	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the download was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: %s (%s)", url, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func SaveFile(outPath string, data []byte) error {
	// Get the directory part from the output path
	dir := filepath.Dir(outPath)

	// Check if the directory exists, if not, create it
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// Write the data to the file
	err := os.WriteFile(outPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	// Successfully saved the image
	fmt.Printf("File saved as %s\n", outPath)
	return nil
}

func SaveJson(outPath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to save json: %w", err)
	}
	return SaveFile(outPath, jsonData)
}
