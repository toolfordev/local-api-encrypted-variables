package application

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"

	"github.com/toolfordev/local-api-encrypted-variables/models"
	"github.com/toolfordev/local-api-encrypted-variables/persistence"
)

type PasswordService struct {
	passwordHashRepository *persistence.PasswordHashRepository
	password               string
}

func (service *PasswordService) passwordToHash(password string) (hash string) {
	hashByte := sha512.Sum512([]byte(password))
	hash = hex.EncodeToString(hashByte[:])
	return
}

func (service *PasswordService) Init(repository *persistence.PasswordHashRepository) {
	service.passwordHashRepository = repository
}

func (service *PasswordService) Set(model *models.Password) (err error) {
	passwordHashes, err := service.passwordHashRepository.GetAll()
	if err != nil {
		return
	}
	if len(passwordHashes) != 0 {
		err = errors.New("password is seted")
		return
	}
	hashModel := &models.PasswordHash{Value: service.passwordToHash(model.Value)}
	err = service.passwordHashRepository.Create(hashModel)
	return
}

func (service *PasswordService) Unlock(model *models.Password) (err error) {
	hashes, err := service.passwordHashRepository.GetAll()
	if err != nil {
		return
	}
	passwordHash := service.passwordToHash(model.Value)
	for _, hash := range hashes {
		if hash.Value == passwordHash {
			service.password = model.Value
			break
		}
	}
	if service.password == "" {
		err = errors.New("invalid password")
	}
	return
}

func (service *PasswordService) Lock() {
	service.password = ""
}

func (service *PasswordService) Encrypt(variable *models.EncryptedVariable) (err error) {
	if service.password == "" {
		err = errors.New("toolfordev locked please unlock")
		return
	}
	text := []byte(variable.Value)
	key := []byte(service.password)
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}
	variable.EncryptedValue = gcm.Seal(nonce, nonce, text, nil)
	return
}

func (service *PasswordService) Decrypt(variable *models.EncryptedVariable) (err error) {
	if service.password == "" {
		err = errors.New("toolfordev locked please unlock")
		return
	}
	key := []byte(service.password)
	ciphertext := variable.EncryptedValue
	if err != nil {
		return
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}
	variable.Value = string(plaintext)
	return
}

func (service *PasswordService) DecryptAll(encryptedVariables []models.EncryptedVariable) (decryptedVariables []models.EncryptedVariable, err error) {
	decryptedVariables = make([]models.EncryptedVariable, len(encryptedVariables))
	for i, variable := range encryptedVariables {
		err = service.Decrypt(&variable)
		if err != nil {
			return
		}
		decryptedVariables[i] = variable
	}
	return
}
