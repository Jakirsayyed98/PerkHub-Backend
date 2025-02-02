package admin

import (
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func GetAdminDashBoard(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := store.AdminStore.GetAdminDashBoardData()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data":    result,
		"message": "Admin Dashboard",
	})
}

func GetUserList(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := store.LoginStore.GetAllUserDetail()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data":    result,
		"message": "Admin Dashboard",
	})
}
