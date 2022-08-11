package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Nuwan-Walisundara/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	createAccount1(t)
}

func createAccount1(t *testing.T) Account {

	args := CreateAccountParams{Owner: util.RandomOwner(),
		Balance:  util.RandomMony(),
		Currency: util.RandomCurrency(),
	}

	account, error := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, error)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.NotZero(t, account.ID)
	return account
}

func TestGetAccount(t *testing.T) {
	account1 := createAccount1(t)
	account2, error := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, error)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.Equal(t, account1.Currency, account2.Currency)
}
func TestUpdateAccount(t *testing.T) {
	account1 := createAccount1(t)
	args := UpdateAccountParams{account1.ID, util.RandomMony()}
	account2, error := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, error)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, args.Balance, account2.Balance)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createAccount1(t)
	error := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, error)

	account2, error := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, error)
	require.EqualError(t, error, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}
