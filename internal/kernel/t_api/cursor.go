package t_api

import (
	"github.com/golang-jwt/jwt/v5"
)

var (
	signingMethod = jwt.SigningMethodHS256
	secretKey     = []byte("resonate") // TODO
)

type Cursor[T any] struct {
	Next *T
}

type Claims[T any] struct {
	jwt.RegisteredClaims
	Next *T
}

func NewCursor[T any](tokenString string) (*Cursor[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cursor[T]) Encode() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Cursor[T]) Decode(tokenString string) error { _ = "STUB: not implemented"; return nil }

func (c *Cursor[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Cursor[T]) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Cursor[T]) String() string { _ = "STUB: not implemented"; return "" }
