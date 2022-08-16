package cmd

import (
	"chat/internal/logic"
	"chat/internal/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type App struct {
	H *logic.Hub
}

func (a *App) Run() {
	go a.H.Run()
	router := gin.New()
	router.LoadHTMLFiles("index.html")
	router.GET("/rooms/history/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		query := c.Request.URL.Query()
		offset := 0
		limit := 100
		history := a.H.Storage().LoadMessageHistory(roomId)
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

	router.GET("/rooms/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/rooms", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, a.H.Storage().ListRooms())
	})

	router.GET("/rooms/search", func(c *gin.Context) {
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
		c.IndentedJSON(http.StatusOK, a.H.Storage().Search(value, room))
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		logic.ServeWs(c.Writer, c.Request, roomId, a.H)
	})

	router.Run("0.0.0.0:8080")
}
