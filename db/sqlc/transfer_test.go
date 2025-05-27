package db

import (
	"context"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account_from := CreateRandomAccount(t)
	account_to := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account_from.ID,
		ToAccountID:   account_to.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotEmpty(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	arg := CreateRandomTransfer(t)
	transfer, err := testQueries.GetTransfer(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.ID, transfer.ID)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.WithinDuration(t, arg.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account_from := CreateRandomAccount(t)
	account_to := CreateRandomAccount(t)

	var err error
	var transfer Transfer
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: account_from.ID,
			ToAccountID:   account_to.ID,
			Amount:        util.RandomMoney(),
		}
		transfer, err = testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)

		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
		require.Equal(t, arg.Amount, transfer.Amount)

		require.NotEmpty(t, transfer.CreatedAt)
	}

	arg_list := ListTransfersParams{
		FromAccountID: account_from.ID,
		ToAccountID:   account_to.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err1 := testQueries.ListTransfers(context.Background(), arg_list)
	require.NoError(t, err1)
	require.Len(t, transfers, 5)

	for _, transfer_obj := range transfers {
		require.NotEmpty(t, transfer_obj)
	}

}
