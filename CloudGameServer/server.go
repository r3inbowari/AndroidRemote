package CloudGameServer

import (
	bilicoin "CloudGameServer/utils"
	"github.com/gorilla/mux"
	"github.com/wuwenbao/gcors"
	"log"
	"net/http"
)

func BCApplication() {
	bilicoin.Info("Cloud Game Server running")
	bilicoin.Info("[BCS] Listened on " + bilicoin.GetConfig().APIAddr)
	r := mux.NewRouter()

	r.HandleFunc("/hello", Handle)

	// allow CORS
	cors := gcors.New(
		r,
		gcors.WithOrigin("*"),
		gcors.WithMethods("POST, GET, PUT, DELETE, OPTIONS"),
		gcors.WithHeaders("Authorization"),
	)
	log.Fatal(http.ListenAndServe(bilicoin.GetConfig().APIAddr, cors))

}

func Handle(w http.ResponseWriter, r *http.Request) {
	bilicoin.ResponseCommon(w, "try succeed", "ok", 1, http.StatusOK, 0)
}
