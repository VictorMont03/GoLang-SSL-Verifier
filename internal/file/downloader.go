package file

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"ssl-verifier/internal/config"
)

func DownloadDomains(cfg config.Config) error {
	if _, err := os.Stat("./internal/file/temp/majuestic_million.csv"); os.IsNotExist(err) {
		log.Println("Downloading Majestic Million CSV")

		out, err := os.Create("./internal/file/temp/majuestic_million.csv")

		if err != nil {
			return err
		}

		defer out.Close()

		resp, err := http.Get(cfg.CSVUrl)

		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to download Majestic Million CSV. Status code: %d", resp.StatusCode)
		}

		_, errCopy := io.Copy(out, resp.Body)

		if errCopy != nil {
			return err
		}

		log.Println("Majestic Million CSV downloaded")
	} else {
		log.Println("Majestic Million CSV already exists")
	}

	return nil
}
