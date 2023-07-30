package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AirQualityHandler interface {
	Run() *http.Server
}

type airQualityHandler struct {
	cfg               *config.Config
	router            *mux.Router
	airQualityService service.AirQualityService
}

func NewAirQualityHandler(cfg *config.Config, srv service.AirQualityService) AirQualityHandler {
	airQualityHandler := &airQualityHandler{
		cfg:               cfg,
		airQualityService: srv,
	}

	router := mux.NewRouter()

	router.HandleFunc("/air-quality", airQualityHandler.airQualityHandler).Methods(http.MethodGet)
	router.HandleFunc("/single-read", airQualityHandler.singleReadHandler).Methods(http.MethodGet)

	router.HandleFunc("/start", airQualityHandler.startHandler).Methods(http.MethodPost)
	router.HandleFunc("/stop", airQualityHandler.stopHandler).Methods(http.MethodPost)

	airQualityHandler.router = router

	return airQualityHandler
}

func (a *airQualityHandler) Run() *http.Server {
	srv := &http.Server{
		Addr:    a.cfg.Address,
		Handler: a.router,
	}

	go func() {
		// Will always returns error, but check if it's something that isn't related to graceful shutdown.
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logrus.WithError(err).Error("ListenAndServe failed")
		}
	}()

	return srv
}

func (a *airQualityHandler) airQualityHandler(rsp http.ResponseWriter, req *http.Request) {
	marshalOkRsp(rsp, a.airQualityService.GetAirQuality())
}

func (a *airQualityHandler) singleReadHandler(rsp http.ResponseWriter, req *http.Request) {
	marshalOkRsp(rsp, a.airQualityService.SingleRead())
}

func (a *airQualityHandler) startHandler(rsp http.ResponseWriter, req *http.Request) {
	a.airQualityService.Start()

	marshalOkRsp(rsp, struct{}{})
}

func (a *airQualityHandler) stopHandler(rsp http.ResponseWriter, req *http.Request) {
	a.airQualityService.Stop()

	marshalOkRsp(rsp, struct{}{})
}

func marshalOkRsp(rsp http.ResponseWriter, item interface{}) {
	jsonResponse, err := json.Marshal(item)
	if err != nil {
		logrus.WithError(err).Error("Failed marshaling item")
	}

	rsp.Header().Set("Content-Type", "application/json")
	rsp.WriteHeader(http.StatusOK)

	_, err = rsp.Write(jsonResponse)
	if err != nil {
		logrus.WithError(err).Error("Failed writing item")
	}
}
