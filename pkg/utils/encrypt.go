package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func RSA_OAEP_Encrypt(secretMessage string, public rsa.PublicKey) (string, error) {
	label := []byte("OAEP Encrypted")
	hash := sha256.New()
	random := rand.Reader

	msg := []byte(secretMessage)
	msgLen := len(msg)
	step := public.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash, random, &public, msg[start:finish], label)
		if err != nil {
			return "", err
		}

		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), nil

}

func RSA_OAEP_Decrypt(cipherText string, private rsa.PrivateKey) (string, error) {
	label := []byte("OAEP Encrypted")
	hash := sha256.New()
	random := rand.Reader

	msg, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	msgLen := len(msg)
	step := private.PublicKey.Size()
	var decryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		decryptedBlockBytes, err := rsa.DecryptOAEP(hash, random, &private, msg[start:finish], label)
		if err != nil {
			return "", err
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}

	return string(decryptedBytes), nil
}
