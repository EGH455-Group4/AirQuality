package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AirQualityHandler struct {
	cfg    *config.Config
	router *mux.Router
}

func NewAirQualityHandler(cfg *config.Config) *AirQualityHandler {
	router := mux.NewRouter()

	router.HandleFunc("/air-quality", airQualityHandler).Methods(http.MethodGet)
	router.HandleFunc("/single-read", singleReadHandler).Methods(http.MethodGet)

	router.HandleFunc("/start", startHandler).Methods(http.MethodPost)
	router.HandleFunc("/stop", stopHandler).Methods(http.MethodPost)

	return &AirQualityHandler{
		cfg:    cfg,
		router: router,
	}
}

func (a *AirQualityHandler) Run() error {
	return http.ListenAndServe(a.cfg.Address, a.router)
}

func airQualityHandler(rsp http.ResponseWriter, req *http.Request) {
	airQuality := &models.AirQuality{}

	marshalOkRsp(rsp, airQuality)
}

func singleReadHandler(rsp http.ResponseWriter, req *http.Request) {
	airQuality := &models.AirQuality{}

	marshalOkRsp(rsp, airQuality)
}

func startHandler(rsp http.ResponseWriter, req *http.Request) {
	marshalOkRsp(rsp, struct{}{})
}

func stopHandler(rsp http.ResponseWriter, req *http.Request) {
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
