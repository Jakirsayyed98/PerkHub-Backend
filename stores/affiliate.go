package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"database/sql"
	"fmt"
	"strconv"
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
