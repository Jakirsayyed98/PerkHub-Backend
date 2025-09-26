// scheduler/game_scheduler.go
package scheduler

import (
	"database/sql"
	"log"

	"PerkHub/constants"
	"PerkHub/model"
	"PerkHub/services"

	"github.com/robfig/cron/v3"
)

type GameNotificationScheduler struct {
	Notifier *services.NotificationService
	db       *sql.DB
}

func NewGameNotificationScheduler(dbs *sql.DB) *GameNotificationScheduler {
	notificationService := services.NewNotificationService(
		constants.FirebaseProjectID, constants.FireBaseFilePath,
	)
	return &GameNotificationScheduler{db: dbs, Notifier: notificationService}
}

func (s *GameNotificationScheduler) Start() {
	c := cron.New()
	// Run every 6 hours, can adjust to 6–8 with randomness
	c.AddFunc(constants.GameCron, func() {

		game, err := model.GetRandomGame(s.db)
		if err != nil {
			log.Println("No game available or error:", err)
			return
		}

		title := "Play " + game.Name + " 🎮"
		message := game.Description
		data := map[string]interface{}{"game_id": game.Id, "url": game.URL, "icon": game.Assets.Square}
		err = s.Notifier.SendNotificationToAllUsers(title, message, data)
		if err != nil {
			log.Println("Failed to send notification:", err)
			return
		}
	})
	c.Start()
}
