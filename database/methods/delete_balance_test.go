package methods

import (
	"context"
	"testing"

	"github.com/balance/api/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteBalance(t *testing.T) {
	var db Postgres
	db.conn = database.ConnectDB()

	var id int64 = 3
	err := db.DeleteBalance(context.Background(), id)

	require.NoError(t, err)
	assert.NotZero(t, id)
}
