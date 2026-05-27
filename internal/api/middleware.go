package api

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
	"github.com/resonatehq/resonate/internal/kernel/t_api"
)

type Middleware interface {
	Process(r *t_api.Request) *t_api.Error
}

// TODO(avillega): This authenticator is the first middleware that we currently have
// as we have more middleware we should rethink where to put them
type JwtAuthenticator struct {
	publicKey *rsa.PublicKey
}

type Claims struct {
	Prefix *string `json:"prefix"`
	Role   string  `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTAuthenticator(publicKeyPEM []byte) (*JwtAuthenticator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *JwtAuthenticator) Process(req *t_api.Request) *t_api.Error {
	_ = "STUB: not implemented"
	return nil
}

func (a *JwtAuthenticator) authenticate(req *t_api.Request) (*Claims, error) {
	_ = "STUB: not implemented"
	// Assume what ever is in the metadata["authorization"] is just the token without
	// 'Bearer' or other prefixes
	return nil, nil
}

// ParseWithClaims also checks registered claims like expiry time

// Only support RSA (private/public key) signed method

func (a *JwtAuthenticator) authorize(claims *Claims, req *t_api.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// Admins have access to all promise prefixes

// "aaabb" is more restrictive than "aaa" which is more restrictive than "a" which is more restrictive that ""
// Empty string gets access to all promises

func matchPromisePrefix(req *t_api.Request, prefix string) error {
	_ = "STUB: not implemented"
	return nil
}

// Tasks have their own way of matching prefix to their ids

func matchTaskId(taskId, prefix string) error {
	_ = "STUB: not implemented"
	// We expect the taskId to have the one of following formats
	// __resume:{promiseId}:{another}
	// __notify:{promiseId}:{another}
	// __invoke:{promiseId}
	// if that changes we need to change this code
	return nil
}
