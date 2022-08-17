package controllers

import (
	"chat/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSearch(store storage.Storage, router *gin.Engine) {
	router.GET("/rooms/search", func(ctx *gin.Context) {
		searchHandler(store, ctx)
	})
}

func searchHandler(store storage.Storage, c *gin.Context) {
	query := c.Request.URL.Query()
	value := query.Get("value")
	room := query.Get("room")
	if value == "" {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	if room == "" {
		c.IndentedJSON(http.StatusOK, store.GlobalSearch(value))
		return
	}
	if !store.CheckRoom(room) {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, store.Search(value, room))
}
