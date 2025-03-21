package settings

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusBadRequest(c *gin.Context, err interface{}, token string) {

	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	if err == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
			"error":   err,
		},
	)
}

func StatusUnauthorized(c *gin.Context, err interface{}) {
	c.JSON(
		http.StatusUnauthorized,
		gin.H{
			"status":  http.StatusUnauthorized,
			"message": fmt.Sprintf("%v", err),
		},
	)
	c.Abort()
}

func StatusNotFound(c *gin.Context, err interface{}, token string) {

	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	if err == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "Not Found",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusNotFound,
			"message": "Not Found",
			"error":   err,
		},
	)
}

func StatusOk(c *gin.Context, data interface{}, message string, token string) {

	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	if data == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusOK,
				"message": message,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"data":    data,
			"message": message,
		},
	)
}

func StatusCreated(c *gin.Context, data interface{}, message string, token string) {

	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	if data == nil {
		c.JSON(
			http.StatusCreated,
			gin.H{
				"status":  http.StatusCreated,
				"message": message,
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"data":    data,
			"message": message,
		},
	)
}
