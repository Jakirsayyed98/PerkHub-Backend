// scheduler/game_scheduler.go
package scheduler

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"PerkHub/constants"
	"PerkHub/model"
	"PerkHub/services"

	"math/rand"

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

		templates := []string{
			"Time to conquer: %s! 🎮",
			"Can you beat %s today?",
			"Level up with %s!",
			"%s is calling your name! 🔥",
			"Adventure awaits in %s!",
			"Don't miss out on %s!",
			"Your next challenge: %s 🕹️",
			"Get ready for %s!",
			"%s is live! Jump in now!",
			"New high scores await in %s!",
			"Have you played %s yet?",
			"Experience the excitement of %s!",
			"Power up with %s!",
			"Epic fun in %s awaits!",
			"Step into the world of %s!",
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Pick a random template
		titleTemplate := templates[r.Intn(len(templates))]
		title := fmt.Sprintf(titleTemplate, game.Name)

		message := game.Description
		data := map[string]interface{}{"game_id": game.Id, "url": game.URL, "icon": game.Assets.Square, "banner": game.Assets.Wall}
		err = s.Notifier.SendNotificationToAllUsers(title, message, data)
		if err != nil {
			log.Println("Failed to send notification:", err)
			return
		}

		err = model.InsertUserNotificationHistory(s.db, &model.UserNotificationHistory{
			Title:       title,
			Message:     message,
			Image:       game.Assets.Square,
			ClickAction: game.URL,
			UserID:      "",
		})
		if err != nil {
			log.Println("Failed to insert notification:", err)
			return
		}

	})
	c.Start()
}
