package clicker

import (
	"crypto/aes"
	"crypto/cipher"
)

// CryptoClick ...
type CryptoClick struct {
	passphrase string
}

// NewCryptoClick ...
func NewCryptoClick(passphrase string) *CryptoClick {
	return &CryptoClick{
		passphrase: passphrase,
	}
}

func (c CryptoClick) Encrypt(plainText []byte) (string, error) {
	block, err := aes.NewCipher([]byte(c.passphrase))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return base64Encode(cipherText), nil
}

// Decrypt ...
func (c CryptoClick) Decrypt(data string) (string, error) {
	block, err := aes.NewCipher([]byte(c.passphrase))
	if err != nil {
		return "", err
	}
	cipherText := base64Decode(data)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
