package controllers

import (
	"chat/internal/storage"
	"chat/internal/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddSearch(store storage.Storage, router *gin.Engine) {
	router.GET("/rooms/search", func(ctx *gin.Context) {
		historyHandler(store, ctx)
	})

}

func historyHandler(store storage.Storage, c *gin.Context) {
	roomId := c.Param("roomId")
	if !store.CheckRoom(roomId) {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	query := c.Request.URL.Query()
	offset := 0
	limit := 100
	history := store.LoadMessageHistory(roomId)
	if history == nil {
		c.IndentedJSON(http.StatusBadRequest,
			nil)
		return
	}
	msgLen := history.Total
	if limit > msgLen {
		limit = msgLen
	}
	offst := query.Get("offset")

	if val, err := strconv.Atoi(offst); offst != "" && err == nil {
		if val < msgLen {
			offset = val
		}
	}
	lmt := query.Get("limit")
	if val, err := strconv.Atoi(lmt); lmt != "" && err == nil {
		if val < msgLen {
			limit = val
		}
	}
	c.IndentedJSON(http.StatusOK,
		types.MessageHistory{
			Data:  history.Data[offset : offset+limit],
			Total: len(history.Data),
		})
}
