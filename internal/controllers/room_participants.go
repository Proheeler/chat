package controllers

import (
	"chat/internal/storage"
	"chat/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createParticipantInRoom(store storage.Storage, router *gin.Engine) {
	router.POST("/v1/rooms/:roomId/participants", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		req := &types.Req{}
		err = json.Unmarshal(jsonData, req)
		if err != nil {
			fmt.Println(err.Error())
		}

		store.AddParticipantInRoom(req.ID, room)
		c.IndentedJSON(http.StatusCreated, nil)
	})
}

func listParticipantsInRoom(store storage.Storage, router *gin.Engine) {
	router.GET("/v1/rooms/:roomId/participants", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, store.ListParticipantsInRoom(room))
	})
}

func deleteParticipantInRoom(store storage.Storage, router *gin.Engine) {
	router.DELETE("/v1/rooms/:roomId/participants", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		query := c.Request.URL.Query()
		id := query.Get("id")
		store.DeleteParticipantInRoom(id, room)
		c.IndentedJSON(http.StatusOK, nil)
	})
}

func AddRoomParticipantCDL(store storage.Storage, router *gin.Engine) {
	createParticipantInRoom(store, router)
	deleteParticipantInRoom(store, router)
	listParticipantsInRoom(store, router)
}
