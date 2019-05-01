package pwd

import (
	"bytes"
	"log"

	"golang.org/x/crypto/scrypt"
)

// EncryptedPassword type
type EncryptedPassword []byte

// Match type
type Match bool

// EncryptPassword returns encrypted password
func EncryptPassword(pass, salt []byte) *EncryptedPassword {
	hash, err := scrypt.Key(pass, salt, 1<<14, 8, 1, 64)
	if err != nil {
		log.Fatal(err)
	}
	ep := EncryptedPassword(hash)
	return &ep
}

// MatchPassword matches is passwords valid
func MatchPassword(pass, salt []byte) *Match {
	ep := EncryptPassword(pass, salt)

	res := Match(bytes.Equal(pass, []byte(*ep)))
	return &res
}
