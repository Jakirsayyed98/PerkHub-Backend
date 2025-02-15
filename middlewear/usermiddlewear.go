package middlewear

import (
	"PerkHub/settings"
	"PerkHub/utils"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			settings.StatusUnauthorized(ctx, errors.New("authorization header is not provided"))
			return
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			settings.StatusUnauthorized(ctx, errors.New("invalid token format, token should be in Bearer <token> format"))
			return
		}

		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		tokenData, err := utils.VerifyJwt(token)
		if err != nil {
			settings.StatusUnauthorized(ctx, err)
			return
		}

		tokenDataParts := strings.Split(tokenData, "|")

		if len(tokenDataParts) < 2 {
			settings.StatusUnauthorized(ctx, errors.New("invalid token format, token data is incorrect"))
			return
		}

		ctx.Set("user_id", tokenDataParts[0])
		ctx.Next()
	}
}
