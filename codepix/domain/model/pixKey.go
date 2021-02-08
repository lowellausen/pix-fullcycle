package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	PixKeyActive    string = "active"
	PixKeyInactive  string = "inactive"
	PixKeyCpfKind   string = "cpf"
	PixKeyEmailKind string = "email"
)

type PixKeyRepositoryInterface interface {
	RegisterPixKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}

	if pixKey.Kind != PixKeyCpfKind && pixKey.Kind != PixKeyEmailKind {
		return errors.New("invalid key type")
	}

	if pixKey.Status != PixKeyActive && pixKey.Status != PixKeyInactive {
		return errors.New("invalid status")
	}

	return nil
}

func NewPixKey(kind string, key string, account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  PixKeyActive,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
