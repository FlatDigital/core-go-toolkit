package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(secretKey string, plainText string) (encryptedText string, err error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	b := aesgcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(b), nil
}

func Decrypt(secretKey string, encryptedText string) (plainText string, err error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return
	}

	if data == nil || len(data) < aesgcm.NonceSize() {
		return "", fmt.Errorf("invalid encrypted text format")
	}

	nonce, cipherText := data[:aesgcm.NonceSize()], data[aesgcm.NonceSize():]
	b, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return
	}

	return string(b), nil
}
