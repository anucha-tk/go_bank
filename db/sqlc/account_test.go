package db

import (
	"context"
	"testing"

	"github.com/anucha-tk/go_bank/util"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func createFakerAccount(t *testing.T) Account {
	user := createFakerUser(t)

	var params util.CreateFakerAccount
	err := faker.FakeData(&params)
	require.NoError(t, err)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  params.Balance,
		Currency: params.Currency,
	}

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

	params := util.UpdateFakerAccount{
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
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createFakerAccount(t)
	}
	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
