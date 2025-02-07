package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"PerkHub/responses"
	"database/sql"
	"errors"
)

type AdminStore struct {
	db *sql.DB
}

func NewAdminStoreStore(dbs *sql.DB) *AdminStore {
	return &AdminStore{
		db: dbs,
	}
}

func (s *AdminStore) AdminLogin(request *request.AdminLoginRequest) (*model.AdminUser, error) {

	getAdmin, err := model.GetAdmin(s.db, request.Email)
	if err != nil {
		return nil, err
	}

	if getAdmin == nil {
		return nil, errors.New("Invalid username or password 1")
	}

	// password, err := utils.Decrypt(getAdmin.Password)
	// if err != nil {
	// 	return nil, err
	// }
	if request.Password != getAdmin.Password {
		return nil, errors.New("Invalid username or password 2")
	}

	return getAdmin, nil
}

func (s *AdminStore) AdminRegister(request *request.AdminRegister) (interface{}, error) {

	if request.Email == "" && request.Password == "" {
		return nil, errors.New("Invalid username or password")
	}

	if err := model.RegisterAdmin(s.db, request); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AdminStore) GetAdminDashBoardData() (interface{}, error) {
	miniapp, err := model.GetAllMiniApps(s.db)
	if err != nil {
		return nil, err
	}
	gamedata, err := model.GetAllGames(s.db)
	if err != nil {
		return nil, err
	}
	userlist, err := model.AllUsersDetail(s.db)
	if err != nil {
		return nil, err
	}

	response := responses.NewAdminDashBoardData(len(miniapp), len(gamedata), len(userlist))
	return response, nil
}
