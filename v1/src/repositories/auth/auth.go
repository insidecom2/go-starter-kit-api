package repositories

import (
	"go/starter-kit/api/src/consts"
	"go/starter-kit/api/src/utils"
)

func (repo *authRepositoryStruct) Register(req consts.RegisterRequest) (r consts.UserResponse, err error) {
	query := `INSERT INTO "users"
	(name,email,password,status,refresh_token,created_at,updated_at)
	VALUES
	($1,$2,$3,'active','',NOW(),NOW())
	RETURNING id, name, email, status, created_at
	`
	err = repo.DB.QueryRowx(query, req.Name, req.Email, req.Password).StructScan(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (repo *authRepositoryStruct) Login(req consts.LoginRequest) (r consts.UserResponse, err error) {

	var userData consts.UserEntity
	query := `SELECT * FROM
				"users"
				WHERE email = $1
				AND status= 'active'`

	err = repo.DB.QueryRowx(query, req.Email).StructScan(&userData)
	if err != nil {
		return r, err
	}

	isPasswordMatch := utils.ComparePassword(userData.Password, req.Password)
	if isPasswordMatch {
		r = consts.UserResponse{
			Id:        userData.Id,
			Name:      userData.Name,
			Email:     userData.Email,
			Status:    userData.Status,
			CreatedAt: userData.CreatedAt,
		}
		return r, nil
	}

	return r, err

}

func (repo *authRepositoryStruct) UpdateRefreshToken(req consts.UserEntity) (r consts.UserEntity, err error) {

	query := `UPDATE "users" SET
				refresh_token=$1
				WHERE id=$2
				RETURNING *`

	err = repo.DB.QueryRowx(query, req.ReFreshToken, req.Id).StructScan(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}
