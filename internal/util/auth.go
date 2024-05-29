package util

import (
	"golang.org/x/crypto/bcrypt"
)

var salt = "flomeiscool"

func CompareHashAndPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(salt+password))
}
