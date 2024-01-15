package farm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

const encryptionKey = "F6FFAF11E0BC4B27ADDA7E4F4B91B4BA"

type encodeableFarm struct {
	Name         string
	W, H         int
	TimeScale    time.Duration
	Field        [][]*Crop
	Money        int
	LastModified time.Time
}

func (f *Farm) Marshal() ([]byte, error) {
	jsonData, err := json.Marshal(encodeableFarm{
		f.name,
		f.w, f.h,
		f.timeScale,
		f.field,
		f.money,
		time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling json: %w", err)
	}

	encrypted, err := encrypt(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error encrypting: %w", err)
	}

	return encrypted, nil
}

func createCipher() (cipher.AEAD, error) {
	c, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return nil, fmt.Errorf("error creating cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, fmt.Errorf("error creating gcm: %w", err)
	}

	return gcm, nil
}

func decrypt(in []byte) ([]byte, error) {
	gcm, err := createCipher()
	if err != nil {
		return nil, fmt.Errorf("error creating cipher: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(in) < nonceSize {
		return nil, fmt.Errorf("wrong nonce size")
	}

	nonce, in := in[:nonceSize], in[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, in, nil)
	if err != nil {
		return nil, fmt.Errorf("error decrypting: %w", err)
	}

	return plaintext, nil
}

func encrypt(in []byte) ([]byte, error) {
	gcm, err := createCipher()
	if err != nil {
		return nil, fmt.Errorf("error creating cipher: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, fmt.Errorf("error reading nonce: %w", err)
	}

	return gcm.Seal(nonce, nonce, in, nil), nil
}
