package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()
	router := gin.New()
	router.LoadHTMLFiles("index.html")
	router.GET("/rooms/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		query := c.Request.URL.Query()
		offset := 0
		limit := 100
		msgLen := h.history[roomId].Total
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
			MessageHistory{
				Data:  h.history[roomId].Data[offset : offset+limit],
				Total: len(h.history[roomId].Data),
			})
	})

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	router.Run("0.0.0.0:8080")
}
