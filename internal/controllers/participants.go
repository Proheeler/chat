package controllers

import (
	"chat/internal/storage"
	"chat/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createParticipant(store storage.Storage, router *gin.Engine) {
	router.POST("/rooms/:roomId/participants", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		pers := &types.Client{}
		err = json.Unmarshal(jsonData, pers)
		if err != nil {
			fmt.Println(err.Error())
		}
		id := uuid.New()
		pers.ID = id.String()
		pers.CreateAt = time.Now()
		pers.UpdateAt = time.Now()
		store.StoreParticipant(*pers, room)
		c.IndentedJSON(http.StatusCreated, nil)
	})
}

func readParticipant(store storage.Storage, router *gin.Engine) {
	router.GET("/rooms/:roomId/participants", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, store.LoadParticipants(room))
	})
}
func updateParticipant(store storage.Storage, router *gin.Engine) {
	router.PATCH("/rooms/:roomId/participants", func(c *gin.Context) {
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

func deleteParticipant(store storage.Storage, router *gin.Engine) {
	router.DELETE("/rooms/:roomId/participants", func(c *gin.Context) {
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

func AddParticipantCRUD(store storage.Storage, router *gin.Engine) {
	createParticipant(store, router)
	readParticipant(store, router)
	updateParticipant(store, router)
	deleteParticipant(store, router)
}
