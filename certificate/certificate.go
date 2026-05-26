// Package certificate contains functions to load an Apple APNs PKCS#12
// or PEM certificate from either an in memory byte array or a local file.
package certificate

import (
	"crypto"
	"crypto/tls"
	"encoding/pem"
	"errors"
)

// Possible errors when parsing a certificate.
var (
	ErrFailedToParsePrivateKey  = errors.New("failed to parse private key")
	ErrFailedToParseCertificate = errors.New("failed to parse certificate PEM data")
	ErrNoPrivateKey             = errors.New("no private key")
	ErrNoCertificate            = errors.New("no certificate")
)

// FromP12File loads a PKCS#12 certificate from a local file and returns a
// tls.Certificate.
//
// Use "" as the password argument if the PKCS#12 certificate is not password
// protected.
func FromP12File(filename string, password string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// FromP12Bytes loads a PKCS#12 certificate from an in memory byte array and
// returns a tls.Certificate.
//
// Use "" as the password argument if the PKCS#12 certificate is not password
// protected.
func FromP12Bytes(bytes []byte, password string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// FromPemFile loads a PEM certificate from a local file and returns a
// tls.Certificate. This function is similar to the crypto/tls LoadX509KeyPair
// function, however it supports PEM files with the cert and key combined
// in the same file. It does not support password-protected key files due
// to security concerns with the deprecated PEM encryption method.
//
// The password argument is kept for backwards compatibility but is no longer used.
func FromPemFile(filename string, password string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// FromPemBytes loads a PEM certificate from an in memory byte array and
// returns a tls.Certificate. This function is similar to the crypto/tls
// X509KeyPair function, however it supports PEM files with the cert and
// key combined, as well as password protected keys which are both common with
// APNs certificates.
//
// Use "" as the password argument if the PEM certificate is not password
// protected.
func FromPemBytes(bytes []byte, password string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

func unencryptPrivateKey(block *pem.Block, password string) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func parsePrivateKey(bytes []byte) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}
