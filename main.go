package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	text := "123456"
	key := "passphrasewhichneedstobe32bytes!"

	fmt.Println(Encrypt(text, key))
}

func Encrypt(textForEncryption string, encryptionKey string) []byte {
	text := []byte(textForEncryption)
	key := []byte(encryptionKey)

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	encryptedText := gcm.Seal(nonce, nonce, text, nil)
	return encryptedText
}
