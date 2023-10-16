package model_test

import (
	"testing"

	"github.com/paulojr-eco/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "testNumber"
	ownerName := "Paulo"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil((account.ID)))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.OwnerName, ownerName)
	require.Equal(t, account.Bank.ID, bank.ID)

	_, err = model.NewAccount(bank, "", "")
	require.NotNil(t, err)
}
