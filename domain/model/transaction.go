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
	Transaction []*Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;" valid:"-"`
	Status            string   `gorm:"type:varchar(28)" json:"status" valid:"notnull"`
	Description       string   `gorm:"type:varchar(255)" json:"description" valid:"notnull"`
	CancelDescription string   `gorm:"type:varchar(255)" json:"cancel_description" valid:"notnull"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if transaction.Status != TransactionPending && transaction.Status != TransactionCompleted && transaction.Status != TransactionError {
		return errors.New("invalid status of this Transaction")
	}

	if transaction.PixKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewTransaction(accountfrom *Account, amount float64, pixkeyto *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountfrom,
		Amount:      amount,
		PixKeyTo:    pixkeyto,
		Description: description,
		Status:      TransactionPending,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (transaction *Transaction) Completed() error {
	transaction.Status = TransactionCompleted
	transaction.UpDatedAt = time.Now()
	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.UpDatedAt = time.Now()
	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TransactionError
	transaction.UpDatedAt = time.Now()
	transaction.Description = description
	err := transaction.isValid()
	return err
}
