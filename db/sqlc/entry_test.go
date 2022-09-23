package db

import (
	"context"
	"github.com/anhuet/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}
func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	CreateRandomEntry(t, account)
}
func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := CreateRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry2.ID, entry.ID)
	require.Equal(t, entry2.Amount, entry.Amount)
	require.Equal(t, entry2.AccountID, entry.AccountID)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}
func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		CreateRandomEntry(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    0,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
