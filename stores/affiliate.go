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

		commission, err := strconv.Atoi(req.Commission)
		if err != nil {
			fmt.Println("Error converting commission:", err)
			return nil, err
		}

		userCommissionPercentageInt, err := strconv.Atoi(UserCommisionPercentage)
		if err != nil {
			fmt.Println("Error converting commission percentage:", err)
			return nil, err
		}

		usercommision := (int(commission) / 100) * userCommissionPercentageInt

		request := model.NewMiniAppTransaction()
		if err := request.Bind(req, fmt.Sprintf("%d", usercommision)); err != nil {
			return nil, err
		}
		req.CommissionPercentage = UserCommisionPercentage
		err = model.InsertMiniAppTransaction(s.db, request)
		if err != nil {
			fmt.Println("UserComision :- ", err.Error())
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
