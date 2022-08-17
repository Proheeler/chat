package cmd

import (
	"chat/internal/controllers"
	"chat/internal/logic"
	"chat/internal/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (a *App) Run() {
	go a.H.Run()
	router := gin.New()
	router.LoadHTMLFiles("index.html")
	router.GET("/rooms/:roomId/history", a.historyHandler)

	router.GET("/rooms/test/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Participants
	controllers.AddParticipantCRUD(a.H.Storage(), router)
	controllers.AddSearch(a.H.Storage(), router)
	controllers.AddRoomCRUD(a.H.Storage(), router)
	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		if !a.H.Storage().CheckRoom(roomId) {
			c.IndentedJSON(http.StatusBadRequest, nil)
			return
		}
		// check for client in body + check if client can connect to room
		// a.H.Storage().AddRoom(roomId)
		logic.ServeWs(c.Writer, c.Request, roomId, a.H)
	})

	router.Run("0.0.0.0:8080")
}
