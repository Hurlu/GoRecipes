package main

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/argon2"
)

type ArgonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateFromPassword(password string, p *ArgonParams) (hash []byte, err error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return nil, err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash = argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return hash, nil
}

func create_user(username string, password string) {
//	p := &ArgonParams{
//	memory:      64 * 1024,
//	iterations:  3,
//	parallelism: 2,
//	saltLength:  16,
//	keyLength:   32,
//}
//	hash, err := generateFromPassword("password123", p)
//	if err != nil {
//		log.Fatal(err)
//	}

	return
}

func auth_routes(router *gin.Engine) {

}
