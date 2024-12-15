package db

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func createFakerAccount(t *testing.T) Account {
	type CreateFakerAccount struct {
		Owner    string `faker:"first_name"`
		Balance  int64  `faker:"boundary_start=31, boundary_end=88"`
		Currency string `faker:"oneof:USD,EUR,CAD"`
	}

	var params CreateFakerAccount
	err := faker.FakeData(&params)
	require.NoError(t, err)

	arg := CreateAccountParams(params)

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createFakerAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createFakerAccount(t)
	result, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, account.ID, result.ID)
	require.Equal(t, account.Balance, result.Balance)
	require.Equal(t, account.Currency, result.Currency)
}

func TestUpdateAccount(t *testing.T) {
	account := createFakerAccount(t)

	type UpdateFakerAccount struct {
		ID      int64 `faker:"-"`
		Balance int64 `faker:"boundary_start=500, boundary_end=999"`
	}

	params := UpdateFakerAccount{
		ID: account.ID,
	}

	err := faker.FakeData(&params)
	require.NoError(t, err)

	arg := UpdateAccountParams(params)
	result, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.ID, account.ID)
	require.Equal(t, result.Owner, account.Owner)
	require.Equal(t, result.Balance, params.Balance)
	require.Equal(t, result.Currency, account.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := createFakerAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	// make sure account is delete
	result, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, result)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createFakerAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, v := range accounts {
		require.NotEmpty(t, v)
	}
}
