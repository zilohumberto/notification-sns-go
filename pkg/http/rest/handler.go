package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/zilohumberto/notification-sns-go/pkg/sending"
)

var (
	router = gin.Default()
)

// Handler generate all endpoints rest for the service
func Handler(){
	sender := sending.NewService()
	mapUrls(sender)
	_ = router.Run(":8980")
}
