package util

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	var params FakerPassword
	err := faker.FakeData(&params)
	require.NoError(t, err)

	hashedPassword, err := HashPassword(params.Password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(params.Password, hashedPassword)
	require.NoError(t, err)

	var wrongParams FakerPassword
	err = faker.FakeData(&wrongParams)
	require.NoError(t, err)

	err = CheckPassword(wrongParams.Password, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
