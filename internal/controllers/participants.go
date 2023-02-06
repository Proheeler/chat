package controllers

import (
	"chat/internal/storage"
	"chat/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func createParticipant(store storage.Storage, router *gin.Engine) {
	router.POST("/v1/participants", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		pers := &types.Client{}
		err = json.Unmarshal(jsonData, pers)
		if err != nil {
			fmt.Println(err.Error())
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		if pers.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}

		pers.CreatedAt = time.Now()
		pers.UpdatedAt = time.Now()
		store.StoreParticipant(*pers)
		c.IndentedJSON(http.StatusCreated, nil)
	})
}

func listParticipants(store storage.Storage, router *gin.Engine) {
	router.GET("/v1/participants", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, store.ListParticipants())
	})
}

func readParticipant(store storage.Storage, router *gin.Engine) {
	router.GET("/v1/participants/:external_id", func(c *gin.Context) {
		clientID := c.Param("external_id")
		ID, err := strconv.Atoi(clientID)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}
		part := store.GetParticipant(uint(ID))
		if part.ID == 0 {
			c.IndentedJSON(http.StatusNotFound, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, part)
	})
}
func updateParticipant(store storage.Storage, router *gin.Engine) {
	router.PATCH("/v1/participants", func(c *gin.Context) {
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
		store.EditParticipant(*pers)
		c.IndentedJSON(http.StatusOK, nil)
	})
}

func deleteParticipant(store storage.Storage, router *gin.Engine) {
	router.DELETE("/v1/participants", func(c *gin.Context) {
		query := c.Request.URL.Query()
		id := query.Get("id")
		ID, err := strconv.Atoi(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}
		store.DeleteParticipant(uint(ID))
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
