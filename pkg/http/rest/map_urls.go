package rest

import (
	"github.com/zilohumberto/notification-sns-go/pkg/ping"
	"github.com/zilohumberto/notification-sns-go/pkg/sending"
)

func mapUrls(s sending.Service) {
	router.GET("/ping", ping.Ping)
	router.POST("/push", sending.SendPushController(s))
	router.POST("/topic", sending.SendTopicController(s))
}
