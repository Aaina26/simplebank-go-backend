package db

import (
	"context"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(-10, 10),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	arg := CreateRandomEntry(t)

	entry, err := testQueries.GetEntry(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.ID, entry.ID)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.WithinDuration(t, arg.CreatedAt, entry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T){
	account:=CreateRandomAccount(t)
	for i:=0;i<10;i++{
		arg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomInt(-10, 10),
		}
		entry, err:= testQueries.CreateEntry(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, entry)

		require.Equal(t, arg.AccountID, entry.AccountID)
		require.Equal(t, arg.Amount, entry.Amount)
		require.NotZero(t, entry.CreatedAt)
		require.NotZero(t, entry.ID)

	}

	arg1:=ListEntriesParams{
		AccountID: account.ID,
		Offset: 5,
		Limit: 5,
	}

	entries, err:=testQueries.ListEntries(context.Background(),arg1)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry:=range entries{
		require.NotEmpty(t, entry)
	}
}
