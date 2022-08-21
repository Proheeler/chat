package controllers

import (
	"chat/internal/storage"
	"chat/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func listMessages(store storage.Storage, router *gin.Engine) {
	router.GET("/v1/rooms/:roomId/messages", func(c *gin.Context) {
		roomId := c.Param("roomId")
		if !store.CheckRoom(roomId) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		query := c.Request.URL.Query()
		offset := 0
		limit := 100
		history := store.ListMessages(roomId)
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
	})
}

func readMessage(store storage.Storage, router *gin.Engine) {
	router.GET("/v1/rooms/:roomId/messages/:id", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		messageID := c.Param("id")
		msg := store.GetMessage(messageID, room)
		if msg.ID == "" {
			c.IndentedJSON(http.StatusNotFound, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, msg)
	})
}
func updateMessage(store storage.Storage, router *gin.Engine) {
	router.PATCH("/v1/rooms/:roomId/messages", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		pers := &types.Client{}
		err = json.Unmarshal(jsonData, pers)
		if err != nil {
			fmt.Println(err.Error())
		}
		store.EditParticipant(*pers, room)
		c.IndentedJSON(http.StatusOK, nil)
	})
}

func deleteMessage(store storage.Storage, router *gin.Engine) {
	router.DELETE("/v1/rooms/:roomId/messages", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		query := c.Request.URL.Query()
		id := query.Get("id")
		store.DeleteParticipant(id, room)
		c.IndentedJSON(http.StatusOK, nil)
	})
}

func AddMessagesRUDL(store storage.Storage, router *gin.Engine) {
	readMessage(store, router)
	updateMessage(store, router)
	deleteMessage(store, router)
	listMessages(store, router)
}
