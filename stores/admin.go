package stores

import (
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/request"
	"PerkHub/responses"
	"database/sql"
	"errors"
	"time"
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
	startTime := time.Now()
	getAdmin, err := model.GetAdmin(s.db, request.Email)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	if getAdmin == nil {
		log := logger.LogData{
			Message:   "Invalid username or password",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("Invalid username or password")
	}

	// password, err := utils.Decrypt(getAdmin.Password)
	// if err != nil {
	// 	return nil, err
	// }
	if request.Password != getAdmin.Password {
		log := logger.LogData{
			Message:   "Invalid username or password",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("Invalid username or password")
	}

	return getAdmin, nil
}

func (s *AdminStore) AdminRegister(request *request.AdminRegister) (interface{}, error) {
	startTime := time.Now()
	if request.Email == "" && request.Password == "" {
		log := logger.LogData{
			Message:   "Invalid username or password",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("Invalid username or password")
	}

	if err := model.RegisterAdmin(s.db, request); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return nil, nil
}

func (s *AdminStore) GetAdminDashBoardData() (interface{}, error) {
	startTime := time.Now()
	miniapp, err := model.GetAllMiniApps(s.db)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	gamedata, err := model.GetAllGames(s.db)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	userlist, err := model.AllUsersDetail(s.db)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	response := responses.NewAdminDashBoardData(len(miniapp), len(gamedata), len(userlist))
	return response, nil
}

func (S *AdminStore) AffiliateTransactions(request *request.AdminAffiliateTransactionsRequest, status string) (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetAllAffiliateTransactions(S.db, request.PageNo, request.Limit, status)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return data, nil
}
