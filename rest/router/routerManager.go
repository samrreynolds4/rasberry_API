package router

import (
	"log"
	"net/http"
	"rasberryPi_api/utils"

	"github.com/gorilla/mux"
)

type handler func()

type routerManager struct {
	Addr   string
	Router *mux.Router
}

var logger utils.Logger

func NewManager(addr string) routerManager {
	return routerManager{
		Router: mux.NewRouter(),
		Addr:   addr,
	}

}

func (mgr *routerManager) AddFunc(route string, handler func(w http.ResponseWriter, r *http.Request), http string) (err error) {
	log.Printf("Adding an http function to manage %s at %s", http, route)
	mgr.Router.HandleFunc(route, handler)

	return
}

func (mgr *routerManager) Start() (err error) {
	logger.Log("Starting server")
	err = http.ListenAndServe(mgr.Addr, mgr.Router)
	if err != nil {
		logger.LogErrorObj("Server failure", err)
		return err
	}

	return nil
}
