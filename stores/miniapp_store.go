package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"PerkHub/response"
	"PerkHub/utils"
	"database/sql"
	"fmt"
)

type MiniAppStore struct {
	db *sql.DB
}

func NewMiniAppStore(dbs *sql.DB) *MiniAppStore {
	return &MiniAppStore{
		db: dbs,
	}
}

func (s *MiniAppStore) CreateMiniApp(req *request.MiniAppRequest) (interface{}, error) {
	if err := model.InsertMiniAppData(s.db, req); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *MiniAppStore) ActivateSomekey(req *request.ActiveDeactiveMiniAppReq) (interface{}, error) {
	if err := model.ActivateSomekey(s.db, req.Key, req.ID, req.Value); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *MiniAppStore) GetAllMiniApps() (interface{}, error) {

	data, err := model.GetAllMiniApps(s.db)

	if err != nil {
		return nil, err
	}
	res := response.NewMiniAppRes()

	result, err := res.BindMultipleUsers(data)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *MiniAppStore) GetMiniAppsBycategoryID(id string) (interface{}, error) {

	data, err := model.GetMiniAppsByCategoryID(s.db, id)

	if err != nil {
		return nil, err
	}
	res := response.NewMiniAppRes()

	result, err := res.BindMultipleUsers(data)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *MiniAppStore) SearchMiniApps(req *request.MiniAppSearchReq) (interface{}, error) {

	data, err := model.SearchMiniApps(s.db, req.Name)

	if err != nil {
		return nil, err
	}
	res := response.NewMiniAppRes()
	if len(data) == 0 {
		return data, nil
	}

	result, err := res.BindMultipleUsers(data)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *MiniAppStore) DeletMniApp(id string) (interface{}, error) {
	if err := model.DeleteMiniAppByID(s.db, id); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *MiniAppStore) GenrateSubid(miniapp_id, userID string) (interface{}, error) {
	data, err := model.SearchMiniApps(s.db, miniapp_id)

	if err != nil {
		return nil, err
	}
	subid1, err := utils.GenerateRandomUUID(20)
	subid2 := userID
	subid3 := miniapp_id
	url := fmt.Sprintf("%s&subid=%s&subid2=%s&subid3=%s", data[0].Url, subid1, subid2, subid3)

	err = model.InsertGenratedSubId(s.db, miniapp_id, userID, subid1, subid2)
	if err != nil {
		return nil, err
	}
	return url, nil
}
