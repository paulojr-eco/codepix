package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transactions []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid" valid:"-"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;not null" valid:"-"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if transaction.Status != TransactionPending && transaction.Status != TransactionCompleted && transaction.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	if transaction.PixKeyTo.Account.ID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (transaction *Transaction) Complete() error {
	err := transaction.isValid()
	if err != nil {
		return err
	}
	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()
	return nil
}

func (transaction *Transaction) Confirm() error {
	err := transaction.isValid()
	if err != nil {
		return err
	}
	transaction.Status = TransactionConfirmed
	transaction.UpdatedAt = time.Now()
	return nil
}

func (transaction *Transaction) Cancel(cancelDescription string) error {
	err := transaction.isValid()
	if err != nil {
		return err
	}
	transaction.Status = TransactionError
	transaction.UpdatedAt = time.Now()
	transaction.CancelDescription = cancelDescription
	return nil
}
