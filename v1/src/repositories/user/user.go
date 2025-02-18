package repositories

import (
	"go/starter-kit/api/src/consts"
)

func (repo *UserRepositoryStruct) GetUserById(id float64) (r consts.UserResponse, err error) {
	query := `SELECT id,email,name,status,created_at 
				FROM "users"
				WHERE id = $1`

	err = repo.DB.QueryRowx(query, id).StructScan(&r)

	if err != nil {
		return r, err
	}

	return r, nil
}
