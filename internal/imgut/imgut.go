package imgut

import (
	"fmt"
	"github.com/floholz/imgut/internal/pattern"
	"github.com/floholz/imgut/internal/utils"
	"path"
	"path/filepath"
)

func DownloadImages(url string, outDir string) error {
	urls, err := pattern.ResolveUrl(url)
	if err != nil {
		return err
	}
	for _, imgUrl := range urls {
		data, errImg := utils.DownloadFile(imgUrl)
		if errImg != nil {
			fmt.Printf("Error downloading image: %s\n", errImg)
			continue
		}
		outPath := path.Join(outDir, filepath.Base(imgUrl))
		errSave := utils.SaveFile(outPath, data)
		if errSave != nil {
			fmt.Printf("Error saving image: %s\n", errSave)
		}
	}
	return nil
}

func FuzzUrl(url string, outPath string) error {
	urls, err := pattern.ResolveUrl(url)
	if err != nil {
		return err
	}

	if filepath.Ext(outPath) == "" {
		outPath = filepath.Join(outPath, "fuzz.json")
	}
	var validUrls []string
	for _, testUrl := range urls {
		_, errImg := utils.DownloadFile(testUrl)
		if errImg != nil {
			continue
		}
		validUrls = append(validUrls, testUrl)
	}

	errSave := utils.SaveJson(outPath, validUrls)
	if errSave != nil {
		fmt.Printf("Error saving json: %s\n", errSave)
	}
	return nil
}
