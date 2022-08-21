package cmd

import (
	"chat/internal/controllers"
	"chat/internal/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	H *logic.Hub
}

func (a *App) Run() {
	go a.H.Run()
	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/v1/rooms/test/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Participants
	controllers.AddParticipantCRUDL(a.H.Storage(), router)
	controllers.AddRoomParticipantCDL(a.H.Storage(), router)

	// Search
	controllers.AddSearch(a.H.Storage(), router)
	// Rooms
	controllers.AddRoomCRUDL(a.H.Storage(), router)
	// messages
	controllers.AddMessagesRUL(a.H.Storage(), router)
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

	router.Run("0.0.0.0:9080")
}
