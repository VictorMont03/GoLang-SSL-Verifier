package main

import (
	"log"
	"os"
	"ssl-verifier/internal/config"
	"ssl-verifier/internal/file"
	"ssl-verifier/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error while loading the config: %v", err)
	}

	errCsv := file.DownloadDomains(cfg)

	if errCsv != nil {
		log.Fatalf("Failed to download Majestic Million CSV: %v", errCsv)
	}

	urlsList, err := os.Open("internal/file/temp/majuestic_million.csv")

	if err != nil {
		log.Fatalf("Error while opening the file: %v", err)
	}

	urls := file.CreateUrlList(urlsList)

	services.ProcessList(urls, cfg)

	os.Exit(0)
}
