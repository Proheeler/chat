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
		pers := &types.Client{}
		err = json.Unmarshal(jsonData, pers)
		if err != nil {
			fmt.Println(err.Error())
		}
		id := uuid.New()
		pers.ID = id.String()
		pers.CreatedAt = time.Now()
		pers.UpdatedAt = time.Now()
		store.StoreParticipant(*pers, room)
		c.IndentedJSON(http.StatusCreated, nil)
	})
}

func listParticipants(store storage.Storage, router *gin.Engine) {
	router.GET("/v1/rooms/:roomId/participants", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, store.ListParticipants(room))
	})
}

func readParticipant(store storage.Storage, router *gin.Engine) {
	router.GET("/v1/rooms/:roomId/participants/:id", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		clientID := c.Param("id")
		part := store.GetParticipant(room, clientID)
		if part.ID == "" {
			c.IndentedJSON(http.StatusNotFound, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, part)
	})
}
func updateParticipant(store storage.Storage, router *gin.Engine) {
	router.PATCH("/v1/rooms/:roomId/participants", func(c *gin.Context) {
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
	router.DELETE("/v1/rooms/:roomId/participants", func(c *gin.Context) {
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

func AddParticipantCRUDL(store storage.Storage, router *gin.Engine) {
	createParticipant(store, router)
	readParticipant(store, router)
	updateParticipant(store, router)
	deleteParticipant(store, router)
	listParticipants(store, router)
}
