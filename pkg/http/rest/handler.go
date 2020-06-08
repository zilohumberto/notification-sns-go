package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/zilohumberto/notification-sns-go/pkg/sending"
	"log"
	"net/http"
)

func Handler(a sending.Service) http.Handler{
	router := httprouter.New()

	router.POST("/push", sendPush(a))
	return router
}

func sendPush(s sending.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)
		var newNotification sending.Notification
		err := decoder.Decode(&newNotification)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.SendPush(newNotification)
		if err != nil{
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode("Push sent")
		if err != nil{
			log.Fatal(err)
		}
	}
}