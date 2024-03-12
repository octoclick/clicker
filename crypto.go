package clicker

import (
	"crypto/aes"
	"crypto/cipher"
)

// CryptoClick ...
type CryptoClick struct {
	passphrase string
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// NewCryptoClick ...
func NewCryptoClick(
	passphrase string,
) *CryptoClick {
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
