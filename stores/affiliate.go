package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type AffiliatesStore struct {
	db *sql.DB
}

func NewAffiliatesStore(dbs *sql.DB) *AffiliatesStore {
	return &AffiliatesStore{
		db: dbs,
	}
}

func (s *AffiliatesStore) CueLinkCallBack(req *request.CueLinkCallBackRequest) (interface{}, error) {

	_, err := model.FindMiniAppTransactionBySubID(s.db, req)
	if err != nil {
		if err.Error() != "transaction not found" {
			return nil, err
		}
		UserCommisionPercentage := "70"

		commission, err := strconv.ParseFloat(req.Commission, 64)
		if err != nil {
			return nil, err
		}

		userCommissionPercentageInt, err := strconv.Atoi(UserCommisionPercentage)
		if err != nil {
			return nil, err
		}

		usercommision := (commission / 100) * float64(userCommissionPercentageInt)

		request := model.NewMiniAppTransaction()
		if err := request.Bind(req, fmt.Sprintf("%.2f", usercommision)); err != nil {
			return nil, err
		}
		request.CommissionPercentage = UserCommisionPercentage
		err = model.InsertMiniAppTransaction(s.db, request)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
	err = model.UpdateMiniAppTransaction(s.db, req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) CreateAffiliate(req *request.CreateAffiliateRequest) (interface{}, error) {
	// Create a new affiliate in the database
	err := model.CreateAffiliate(s.db, req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) UpdateAffiliate(req *request.CreateAffiliateRequest) (interface{}, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid affiliate ID: %w", err)
	}

	// Update the affiliate in the database
	err = model.UpdateAffiliate(s.db, req, id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) DeleteAffiliate(id string) (interface{}, error) {
	ids, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid affiliate ID: %w", err)
	}
	// Delete the affiliate from the database
	err = model.DeleteAffiliate(s.db, ids)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) UpdateAffiliateFlag(id string, status bool) (interface{}, error) {
	ids, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid affiliate ID: %w", err)
	}
	// Update the affiliate flag in the database
	err = model.UpdateAffiliateFlag(s.db, ids, status)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) ListAffiliates() (interface{}, error) {
	affiliates, err := model.ListAffiliates(s.db)
	if err != nil {
		return nil, err
	}
	return affiliates, nil
}

func (s *AffiliatesStore) GetAffiliateByID(id string) (interface{}, error) {

	affiliate, err := model.GetAffiliateByID(s.db, id)
	if err != nil {
		return nil, err
	}

	return affiliate, nil
}
