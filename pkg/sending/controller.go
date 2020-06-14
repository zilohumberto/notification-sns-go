package sending

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func middleRequest(c *gin.Context, endPoint func(n Notification) error){
	var newNotification Notification
	if err := c.ShouldBindJSON(&newNotification); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := endPoint(newNotification); err != nil{
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Notification sent")
}

// sendPush function called when the request POST /push was triggered
// create and send a notification push
func SendPushController(s Service) func(c *gin.Context){
	return func(c *gin.Context) {
		middleRequest(c, s.SendPush)
	}
}

func SendTopicController(s Service) func(c *gin.Context){
	return func(c *gin.Context) {
		middleRequest(c, s.SendTopic)
	}
}