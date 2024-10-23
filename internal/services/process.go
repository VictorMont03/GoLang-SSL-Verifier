package services

import (
	"encoding/csv"
	"log"
	"os"
	"ssl-verifier/internal/config"
	"ssl-verifier/internal/file"
	"sync"
)

func ProcessList(urls []file.TopURL, cfg config.Config) {
	var wg sync.WaitGroup

	resultFile, err := os.OpenFile("./internal/file/result/majuestic_million_ssl_result.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatalf("Error while opening the result file: %v", err)
	}

	defer resultFile.Close()
	resultFile.WriteString("Server,Port,Ranking,Days Left\n")

	csvWriter := csv.NewWriter(resultFile)

	mu := sync.Mutex{}

	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(url file.TopURL) {
			defer wg.Done()
			line, err := CheckCertificate(url.Address, "443", url.GlobalRanking, cfg)

			if err != nil {
				log.Printf("Error while checking the certificate: %v", err)
				return
			}

			lineCsv := []string{line.Server, line.Port, line.Rasnking, line.DaysLeft}

			mu.Lock()
			csvWriter.Write(lineCsv)
			mu.Unlock()
		}(urls[i])
	}

	wg.Wait()

	csvWriter.Flush()
}
