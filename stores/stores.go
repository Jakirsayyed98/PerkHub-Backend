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
	AdminStore              *AdminStore
	TicketStore             *TicketStore
	OffersStore             *OffersStore
	NotificationStore       *NotificationStore
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
		AdminStore:              NewAdminStoreStore(db),
		TicketStore:             NewTicketStore(db),
		OffersStore:             NewOffersStore(db),
		NotificationStore:       NewNotificationStore(db),
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
		ctx.Set("admin_store", s.AdminStore)
		ctx.Set("ticket_store", s.TicketStore)
		ctx.Set("offers_store", s.OffersStore)
		ctx.Set("notification_store", s.NotificationStore)
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

	adminStore, err := c.MustGet("admin_store").(*AdminStore)
	if !err {
		return nil, errors.New("Admin Store Error")
	}

	ticketStore, err := c.MustGet("ticket_store").(*TicketStore)
	if !err {
		return nil, errors.New("Ticket Store Error")
	}
	offerStore, err := c.MustGet("offers_store").(*OffersStore)
	if !err {
		return nil, errors.New("Offer Store Error")
	}
	notificationStore, err := c.MustGet("notification_store").(*NotificationStore)
	if !err {
		return nil, errors.New("Notification Store Error")
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
		AdminStore:              adminStore,
		TicketStore:             ticketStore,
		OffersStore:             offerStore,
		NotificationStore:       notificationStore,
	}, nil
}

func GetAwsInstance(ctx *gin.Context) (*connection.Aws, error) {
	awsInstance, aOk := ctx.MustGet("aws_instance").(*connection.Aws)

	if !aOk {
		return nil, fmt.Errorf("aws instance not found")
	}

	return awsInstance, nil
}
