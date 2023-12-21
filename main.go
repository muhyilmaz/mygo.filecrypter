package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	key := generateRandomKey()

	plaintext := "muhammed yılmaz"

	encryptedText, err := AesEncrypt(plaintext, key)
	if err != nil {
		fmt.Println("Şifreleme hatası:", err)
		return
	}
	fmt.Printf("Şifrelenmiş metin: %s\n", encryptedText)

	decryptedText, err := AesDecrypt(encryptedText, key)
	if err != nil {
		fmt.Println("Çözme hatası:", err)
		return
	}
	fmt.Printf("Çözülmüş metin: %s\n", decryptedText)
}

func generateRandomKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}
	return key
}

func AesEncrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// IV oluştur
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	// Şifreleme modunu oluştur
	mode := cipher.NewCBCEncrypter(block, iv)

	// Metni şifrele
	plaintextBytes := []byte(plaintext)
	plaintextBytes = pkcs7Padding(plaintextBytes, aes.BlockSize)
	ciphertext := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertext, plaintextBytes)

	// IV'yi şifreli metnin başına ekle
	ciphertextWithIV := append(iv, ciphertext...)

	// Base64 ile kodla
	return base64.StdEncoding.EncodeToString(ciphertextWithIV), nil
}

// Şifreli metni AES-256 kullanarak çöz
func AesDecrypt(encryptedText string, key []byte) (string, error) {
	// Base64'den çöz
	ciphertextWithIV, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// IV'yi al
	iv := ciphertextWithIV[:aes.BlockSize]
	ciphertext := ciphertextWithIV[aes.BlockSize:]

	// Çözme modunu oluştur
	mode := cipher.NewCBCDecrypter(block, iv)

	// Şifreli metni çöz
	plaintextBytes := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintextBytes, ciphertext)

	// PKCS7 dolgusunu kaldır
	plaintextBytes = pkcs7Unpadding(plaintextBytes)

	return string(plaintextBytes), nil
}

// PKCS7 dolgusunu ekle
func pkcs7Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// PKCS7 dolgusunu kaldır
func pkcs7Unpadding(plaintext []byte) []byte {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	return plaintext[:(length - unpadding)]
}
