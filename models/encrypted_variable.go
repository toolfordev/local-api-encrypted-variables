package models

type EncryptedVariable struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Value          string `json:"value"`
	EncryptedValue string `json:"-"`
}
