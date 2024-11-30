package stores

import (
	"PerkHub/connection"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Stores struct {
	LoginStore              *LoginStore
	CategoryStore           *CategoryStore
	MiniAppStore            *MiniAppStore
	BannerStore             *BannerStore
	HomePageStore           *HomePageStore
	AffiliatesStore         *AffiliatesStore
	MiniAppTransactionStore *MiniAppTransactionStore
	GamesStore              *GamesStore
	Withdrawal              *WithdrawalStore
}

func NewStores(db *sql.DB) *Stores {
	return &Stores{
		LoginStore:              NewLoginStore(db),
		CategoryStore:           NewCategoryStore(db),
		MiniAppStore:            NewMiniAppStore(db),
		BannerStore:             NewBannerStore(db),
		HomePageStore:           NewHomePageStore(db),
		AffiliatesStore:         NewAffiliatesStore(db),
		MiniAppTransactionStore: NewMiniAppTransactionStore(db),
		GamesStore:              NewGameStore(db),
		Withdrawal:              NewWithdrawalStore(db),
	}
}

func (s *Stores) BindStore(awsIstance *connection.Aws) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("login_store", s.LoginStore)
		ctx.Set("category_store", s.CategoryStore)
		ctx.Set("miniapp_store", s.MiniAppStore)
		ctx.Set("banner_store", s.BannerStore)
		ctx.Set("homepage_store", s.HomePageStore)
		ctx.Set("affiliates_store", s.AffiliatesStore)
		ctx.Set("miniapptransaction_store", s.MiniAppTransactionStore)
		ctx.Set("games_store", s.GamesStore)
		ctx.Set("withdrawal_store", s.Withdrawal)
		ctx.Set("aws_instance", awsIstance)
		ctx.Next()
	}
}

func GetStores(c *gin.Context) (*Stores, error) {

	loginStore, lok := c.MustGet("login_store").(*LoginStore)

	if !lok {
		return nil, errors.New("login store error")
	}
	categoryStore, cok := c.MustGet("category_store").(*CategoryStore)

	if !cok {
		return nil, errors.New("categoryStore Store Error")
	}

	miniappStore, miniOk := c.MustGet("miniapp_store").(*MiniAppStore)

	if !miniOk {
		return nil, errors.New("MiniApp Store Error")
	}

	bannerStore, bannerOk := c.MustGet("banner_store").(*BannerStore)

	if !bannerOk {
		return nil, errors.New("Banner Store Error")
	}

	homepageStore, homepageOk := c.MustGet("homepage_store").(*HomePageStore)

	if !homepageOk {
		return nil, errors.New("HomePage Store Error")
	}

	affiliates_store, ok := c.MustGet("affiliates_store").(*AffiliatesStore)
	if !ok {
		return nil, errors.New("Affiliate Store Error")
	}

	miniapptransaction_store, miniTok := c.MustGet("miniapptransaction_store").(*MiniAppTransactionStore)
	if !miniTok {
		return nil, errors.New("MiniApp Transaction Store Error")
	}

	gamestore, gameok := c.MustGet("games_store").(*GamesStore)
	if !gameok {
		return nil, errors.New("Games Store Error")
	}

	withdrawalStore, err := c.MustGet("withdrawal_store").(*WithdrawalStore)
	if !err {
		return nil, errors.New("Withdrawal Store Error")
	}

	return &Stores{
		LoginStore:              loginStore,
		CategoryStore:           categoryStore,
		MiniAppStore:            miniappStore,
		BannerStore:             bannerStore,
		HomePageStore:           homepageStore,
		AffiliatesStore:         affiliates_store,
		MiniAppTransactionStore: miniapptransaction_store,
		GamesStore:              gamestore,
		Withdrawal:              withdrawalStore,
	}, nil
}

func GetAwsInstance(ctx *gin.Context) (*connection.Aws, error) {
	awsInstance, aOk := ctx.MustGet("aws_instance").(*connection.Aws)

	if !aOk {
		return nil, fmt.Errorf("aws instance not found")
	}

	return awsInstance, nil
}
