package main

import (
	"crypto/tls"
)

// load the specified tls certificate and key and return the corresponding tls config
func LoadTLSCertificate() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(Config.Verteilzentrum.TlsCertFile, Config.Verteilzentrum.TlsKeyFile)
	if err != nil {
		return &tls.Config{}, err
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}
