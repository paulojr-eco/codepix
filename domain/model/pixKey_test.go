package model_test

import (
	"testing"

	"github.com/paulojr-eco/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixelKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "testNumber"
	ownerName := "Paulo"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "mail@host.com"
	pixKey, _ := model.NewPixKey(kind, account, key)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	_, err := model.NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
