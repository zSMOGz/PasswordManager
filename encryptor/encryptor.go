package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encryptor struct {
	Key string
}

func NewEncryptor() *Encryptor {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в параметры окружения")
	}
	return &Encryptor{
		Key: key,
	}
}

func (enc *Encryptor) Encrypt(plainSrt []byte) []byte {
	//Создание симметричного блока шифрования
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	//Создание GCM
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	//Создание одноразового кода
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	//Шифрование
	return aesGSM.Seal(nonce, nonce, plainSrt, nil)
}

func (enc *Encryptor) Decrypt(encryptedStr []byte) []byte {
	//Создание симметричного блока шифрования
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	//Создание GCM
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return plainText
}
