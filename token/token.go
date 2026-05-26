package token

import (
	"crypto/ecdsa"
	"errors"
	"sync"
)

const (
	// TokenTimeout is the period of time in seconds that a token is valid for.
	// If the timestamp for token issue is not within the last hour, APNs
	// rejects subsequent push messages. This is set to under an hour so that
	// we generate a new token before the existing one expires.
	TokenTimeout = 3000
)

// Possible errors when parsing a .p8 file.
var (
	ErrAuthKeyNotPem   = errors.New("token: AuthKey must be a valid .p8 PEM file")
	ErrAuthKeyNotECDSA = errors.New("token: AuthKey must be of type ecdsa.PrivateKey")
	ErrAuthKeyNil      = errors.New("token: AuthKey was nil")
)

// Token represents an Apple Provider Authentication Token (JSON Web Token).
type Token struct {
	sync.Mutex
	AuthKey  *ecdsa.PrivateKey
	KeyID    string
	TeamID   string
	IssuedAt int64
	Bearer   string
}

// AuthKeyFromFile loads a .p8 certificate from a local file and returns a
// *ecdsa.PrivateKey.
func AuthKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AuthKeyFromBytes loads a .p8 certificate from an in memory byte array and
// returns an *ecdsa.PrivateKey.
func AuthKeyFromBytes(bytes []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateIfExpired checks to see if the token is about to expire and
// generates a new token.
func (t *Token) GenerateIfExpired() (bearer string) { _ = "STUB: not implemented"; return "" }

// Expired checks to see if the token has expired.
func (t *Token) Expired() bool { _ = "STUB: not implemented"; return false }

// Generate creates a new token.
func (t *Token) Generate() (bool, error) { _ = "STUB: not implemented"; return false, nil }
