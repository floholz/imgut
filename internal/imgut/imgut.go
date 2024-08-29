package imgut

import (
	"fmt"
	"github.com/floholz/imgut/internal/pattern"
	"github.com/floholz/imgut/internal/utils"
	nurl "net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func DownloadImages(url string, outDir string, jobs int) error {
	urls, err := pattern.ResolveUrl(url)
	if err != nil {
		return err
	}

	urlObj, err := nurl.Parse(url)
	if err != nil {
		return err
	}
	extendFilename := false
	if !strings.Contains(path.Base(urlObj.Path), "}") {
		extendFilename = true
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, jobs)

	for _, imgUrl := range urls {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(u, o string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			errJob := downloadAndSaveImage(u, o, extendFilename)
			if errJob != nil {
				fmt.Printf("%v", errJob)
			}
		}(imgUrl, outDir)
	}
	wg.Wait()
	return nil
}

func FuzzUrl(url string, outPath string, jobs int) error {
	urls, err := pattern.ResolveUrl(url)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	semaphore := make(chan struct{}, jobs)

	if filepath.Ext(outPath) == "" {
		outPath = filepath.Join(outPath, "fuzz.json")
	}
	var validUrls []string
	for _, testUrl := range urls {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(u string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			_, errImg := utils.DownloadFile(u)
			if errImg == nil {
				mu.Lock()
				validUrls = append(validUrls, u)
				mu.Unlock()
			}
		}(testUrl)
	}
	wg.Wait()

	errSave := utils.SaveJson(outPath, validUrls)
	if errSave != nil {
		fmt.Printf("Error saving json: %s\n", errSave)
	}
	return nil
}

func downloadAndSaveImage(url, outDir string, extendedFilename bool) error {
	data, errImg := utils.DownloadFile(url)
	if errImg != nil {
		return fmt.Errorf("Error downloading image: %s\n", errImg)
	}

	fileName := filepath.Base(url)
	if extendedFilename {
		fileName = strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + fileName
	}
	outPath := path.Join(outDir, fileName)
	errSave := utils.SaveFile(outPath, data)
	if errSave != nil {
		fmt.Printf("Error saving image: %s\n", errSave)
	}
	return nil
}
