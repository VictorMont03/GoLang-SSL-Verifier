package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"ssl-verifier/internal/config"
	"time"
)

type VerifierLine struct {
	Server   string
	Port     string
	Rasnking string
	DaysLeft string
}

func CheckCertificate(server string, port string, ranking string, cfg config.Config) (VerifierLine, error) {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: cfg.SSLCertTimeout}, "tcp", "www."+server+":"+port, nil)

	if err != nil {
		return VerifierLine{}, err
	}

	defer conn.Close()

	expire := conn.ConnectionState().PeerCertificates[0].NotAfter

	currentTime := time.Now()

	daysLeft := expire.Sub(currentTime).Hours() / 24

	log.Printf("Domain: %s, Port: %s, Ranking: %s, Days left: %1.f", server, port, ranking, daysLeft)

	line := VerifierLine{
		Server:   server,
		Port:     port,
		Rasnking: ranking,
		DaysLeft: fmt.Sprintf("%d", int(daysLeft)),
	}

	return line, nil
}
