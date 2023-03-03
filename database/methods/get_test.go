package methods

import (
	"context"
	"testing"

	"github.com/balance/api/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestGetBalance(t *testing.T){
	var db Postgres
	db.conn = database.ConnectDB()

	var id int64 = 3
	balance,currency,err := db.GetBalance(context.Background(), id)

	require.NoError(t, err)
	require.Empty(t,currency)
	
	assert.NotZero(t, id)
	assert.Negative(t,balance)
	assert.NotZero(t,balance)
}