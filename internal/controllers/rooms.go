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

func createRoom(store storage.Storage, router *gin.Engine) {
	router.POST("/rooms", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		room := &types.Room{}
		err = json.Unmarshal(jsonData, room)
		if err != nil {
			fmt.Println(err.Error())
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		id := uuid.New()
		room.ID = id.String()
		room.CreatedAt = time.Now()
		room.UpdatedAt = time.Now()
		room.Participants = []types.Client{}
		room.PinnedMessages = []string{}
		room.History = &types.MessageHistory{}
		store.AddRoom(room)
		c.IndentedJSON(http.StatusCreated, nil)
	})
}

func readRoom(store storage.Storage, router *gin.Engine) {
	router.GET("/rooms/:roomId", func(c *gin.Context) {
		room := c.Param("roomId")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, store.GetRoom(room))
	})
}

func updateRoom(store storage.Storage, router *gin.Engine) {
	router.PATCH("/rooms/:roomId", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
			c.IndentedJSON(http.StatusBadRequest, nil)
		}
		room := &types.Room{}
		err = json.Unmarshal(jsonData, room)
		if err != nil {
			fmt.Println(err.Error())
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		store.EditRoom(room)
		c.IndentedJSON(http.StatusOK, nil)
	})
}

func deleteRoom(store storage.Storage, router *gin.Engine) {
	router.DELETE("/rooms", func(c *gin.Context) {
		query := c.Request.URL.Query()
		room := query.Get("room")
		if !store.CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		store.DeleteRoom(room)
		c.IndentedJSON(http.StatusOK, nil)
	})
}

func listRooms(store storage.Storage, router *gin.Engine) {
	router.GET("/rooms", func(c *gin.Context) {
		rooms := store.ListRooms()
		c.IndentedJSON(http.StatusOK, types.ShortRoomInfoList{
			Total: len(rooms),
			Data:  rooms,
		})
	})
}

func AddRoomCRUD(store storage.Storage, router *gin.Engine) {
	createRoom(store, router)
	readRoom(store, router)
	updateRoom(store, router)
	deleteRoom(store, router)
	listRooms(store, router)
}
