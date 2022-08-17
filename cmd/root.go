package cmd

import (
	"chat/internal/logic"
	"chat/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type App struct {
	H *logic.Hub
}

func (a *App) historyHandler(c *gin.Context) {
	roomId := c.Param("roomId")
	if !a.H.Storage().CheckRoom(roomId) {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	query := c.Request.URL.Query()
	offset := 0
	limit := 100
	history := a.H.Storage().LoadMessageHistory(roomId)
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

func (a *App) searchHandler(c *gin.Context) {
	query := c.Request.URL.Query()
	value := query.Get("value")
	room := query.Get("room")
	if value == "" {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	if room == "" {
		c.IndentedJSON(http.StatusOK, a.H.Storage().GlobalSearch(value))
		return
	}
	if !a.H.Storage().CheckRoom(room) {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, a.H.Storage().Search(value, room))
}

func (a *App) Run() {
	go a.H.Run()
	router := gin.New()
	router.LoadHTMLFiles("index.html")
	router.GET("/rooms/:roomId/history", a.historyHandler)

	router.GET("/rooms/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/rooms", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, a.H.Storage().ListRooms())
	})

	// Participants
	router.GET("/rooms/:roomId/participants", func(c *gin.Context) {
		room := c.Param("roomId")
		if !a.H.Storage().CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		c.IndentedJSON(http.StatusOK, a.H.Storage().LoadParticipants(room))
	})

	router.POST("/rooms/:roomId/participants", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		room := c.Param("roomId")
		if !a.H.Storage().CheckRoom(room) {
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
		a.H.Storage().StoreParticipant(*pers, room)
		c.IndentedJSON(http.StatusCreated, nil)
	})

	router.PATCH("/rooms/:roomId/participants", func(c *gin.Context) {
		room := c.Param("roomId")
		if !a.H.Storage().CheckRoom(room) {
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
		a.H.Storage().EditParticipant(*pers, room)
		c.IndentedJSON(http.StatusOK, nil)
	})

	router.DELETE("/rooms/:roomId/participants", func(c *gin.Context) {
		room := c.Param("roomId")
		if !a.H.Storage().CheckRoom(room) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		query := c.Request.URL.Query()
		id := query.Get("id")
		a.H.Storage().DeleteParticipant(id, room)
		c.IndentedJSON(http.StatusOK, nil)
	})

	router.GET("/rooms/search", a.searchHandler)

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		// if !a.H.Storage().CheckRoom(roomId) {
		// 	c.IndentedJSON(http.StatusBadRequest, nil)
		// 	return
		// }
		// check for client in body + check if client can connect to room
		a.H.Storage().AddRoom(roomId)
		logic.ServeWs(c.Writer, c.Request, roomId, a.H)
	})

	router.Run("0.0.0.0:8080")
}
