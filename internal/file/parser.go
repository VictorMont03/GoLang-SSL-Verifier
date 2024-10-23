package file

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type TopURL struct {
	GlobalRanking string
	DomainRanking string
	Address       string
	Country       string
}

func CreateUrlList(urlList *os.File) []TopURL {
	reader := csv.NewReader(urlList)
	defer urlList.Close()

	var urls []TopURL

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the file: %v", err)
		}

		if record[3] == "br" {
			urls = append(urls, TopURL{
				GlobalRanking: record[0],
				DomainRanking: record[1],
				Address:       record[2],
				Country:       record[3],
			})
		}
	}

	return urls
}
