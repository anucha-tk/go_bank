package db

import (
	"context"
	"testing"
	"time"

	"github.com/anucha-tk/go_bank/util"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
)

func createFakerUser(t *testing.T) User {
	var p util.FakerPassword
	err := faker.FakeData(&p)
	require.NoError(t, err)

	HashedPassword, err := util.HashPassword(p.Password)
	require.NoError(t, err)

	var params util.CreateFakerUser
	err = faker.FakeData(&params)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       params.Username,
		HashedPassword: HashedPassword,
		FullName:       params.FullName,
		Email:          params.Email,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createFakerUser(t)
}

func TestGetUser(t *testing.T) {
	user := createFakerUser(t)
	result, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, user.Username, result.Username)
	require.Equal(t, user.HashedPassword, result.HashedPassword)
	require.Equal(t, user.FullName, result.FullName)
	require.Equal(t, user.Email, result.Email)
	require.WithinDuration(t, user.PasswordChangedAt, result.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, result.CreatedAt, time.Second)
}
