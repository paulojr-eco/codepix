package model_test

import (
	"testing"

	"github.com/paulojr-eco/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	// Create account from
	accountNumber := "testNumber"
	ownerName := "Paulo"
	accountFrom, _ := model.NewAccount(bank, accountNumber, ownerName)

	// Create account to
	accountNumber = "testNumberDestination"
	ownerName = "Junior"
	accountTo, _ := model.NewAccount(bank, accountNumber, ownerName)

	// Create pix key to
	kind := "email"
	key := "mail@host.com"
	pixKeyTo, _ := model.NewPixKey(kind, accountTo, key)

	require.NotEqual(t, accountFrom.ID, accountTo.ID)

	amount := 150.60
	statusTransaction := "pending"
	transaction, err := model.NewTransaction(accountFrom, amount, pixKeyTo, "My description")

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "My description")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, _ := model.NewPixKey(kind, accountTo, key)

	_, err = model.NewTransaction(accountTo, amount, pixKeySameAccount, "My description")
	require.NotNil(t, err)

	_, err = model.NewTransaction(accountFrom, 0, pixKeyTo, "My description")
	require.NotNil(t, err)

}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	// Create account from
	accountNumber := "testNumber"
	ownerName := "Paulo"
	accountFrom, _ := model.NewAccount(bank, accountNumber, ownerName)

	// Create account to
	accountNumber = "testNumberDestination"
	ownerName = "Junior"
	accountTo, _ := model.NewAccount(bank, accountNumber, ownerName)

	// Create pix key to
	kind := "email"
	key := "mail@host.com"
	pixKeyTo, _ := model.NewPixKey(kind, accountTo, key)

	amount := 303.10
	transaction, _ := model.NewTransaction(accountFrom, amount, pixKeyTo, "My description")

	transaction.Complete()
	require.Equal(t, transaction.Status, model.TransactionCompleted)

	transaction.Cancel("Error")
	require.Equal(t, transaction.Status, model.TransactionError)
	require.Equal(t, transaction.CancelDescription, "Error")

}
