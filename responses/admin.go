package responses

type AdminDashBoardData struct {
	MiniAppCount int `json:"minapp_count"`
	GamesCount   int `json:"games_count"`
	UserCount    int `json:"user_count"`
}

func NewAdminDashBoardData(miniAppCount, gamesCount, userCount int) *AdminDashBoardData {
	return &AdminDashBoardData{
		MiniAppCount: miniAppCount,
		GamesCount:   gamesCount,
		UserCount:    userCount,
	}
}
