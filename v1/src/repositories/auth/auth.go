package repositories

import "go/starter-kit/api/src/consts"

func (repo *AuthRepositoryStruct) Register(req consts.RegisterRequest) (r consts.RegisterResponse, err error) {

	query := `INSERT INTO "users" 
			(name,email,password,status,refresh_token,created_at) 
			VALUES
			($1,$2,$3,'active','',NOW()) 
			RETURNING id, name, email, status, created_at
			  `
	err = repo.DB.QueryRowx(query, req.Name, req.Email, req.Password).StructScan(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}
