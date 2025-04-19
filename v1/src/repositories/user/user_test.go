package repositories

import (
	"go/starter-kit/api/src/consts"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := UserRepositoryStruct{DB: sqlxDB}

	mock.ExpectQuery(`SELECT id,email,name,status,created_at
				FROM \"users\"
				WHERE id = \$1`).WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "email", "name", "status", "created_at"}).AddRow(1, "a@a.com", "tom", "active", "2024-01-01"),
	)

	user, err := r.GetUserById(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, consts.UserResponse{
		Id:        1,
		Email:     "a@a.com",
		Name:      "tom",
		Status:    "active",
		CreatedAt: "2024-01-01",
	}, user)
}
