package settings

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusInternalServerError(c *gin.Context, err interface{}, token string) {

	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	if err == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Server Error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
			"error":   err,
		},
	)
}

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
			"message": "Bad Request",
			"error":   err,
		},
	)
}

func StatusBadRequestV3(c *gin.Context, err interface{}, token string) {

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
		http.StatusBadRequest,
		gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
			"error":   err,
		},
	)
}
func StatusUnauthorized(c *gin.Context, err interface{}) {
	c.JSON(
		http.StatusUnauthorized,
		gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized",
			"error":   err,
		},
	)
	c.Abort()
}

func StatusForbidden(c *gin.Context, err interface{}) {

	if err == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusForbidden,
				"message": "Forbidden",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusForbidden,
			"message": "Forbidden",
			"error":   err,
		},
	)
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

func StatusMovedPermanently(c *gin.Context, token string, location string) {

	if token != "" {
		c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	c.Header("Content-Type", "text/html; charset=utf-8")

	c.Redirect(http.StatusMovedPermanently, location)
}
