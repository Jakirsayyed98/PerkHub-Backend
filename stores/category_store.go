package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"database/sql"
)

type CategoryStore struct {
	db *sql.DB
}

func NewCategoryStore(dbs *sql.DB) *CategoryStore {
	return &CategoryStore{
		db: dbs,
	}
}

func (s *CategoryStore) SaveCategory(req *request.Category) (interface{}, error) {

	if err := model.InsertCategory(s.db, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *CategoryStore) UpdateCategory(req *request.Category) (interface{}, error) {

	if err := model.UpdateCategory(s.db, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *CategoryStore) DeleteCategory(id string) (interface{}, error) {

	if err := model.DeleteCategoryByID(s.db, id); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *CategoryStore) GetAllCategory() (interface{}, error) {

	result, err := model.GetAllCategory(s.db)
	if err != nil {
		return nil, err
	}

	// res := response.NewCategoryRes()

	// result, err := res.BindMultipleUsers(data)
	// if err != nil {
	// 	return nil, err
	// }

	return result, nil
}

func (s *CategoryStore) GetCategoryByID(id string) (interface{}, error) {
	result, err := model.GetCategoryByID(s.db, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
