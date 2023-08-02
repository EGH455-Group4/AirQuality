package handler_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/handler"
	"github.com/ImTheTom/air-quality/models"
	"github.com/ImTheTom/air-quality/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type HandlerTestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	cfg *config.Config

	airQualityService *mocks.MockAirQualityService

	airQualityHandler handler.AirQualityHandler
}

func (s *HandlerTestSuite) SetupTest() {
	config.SetupConfig("")

	s.cfg = config.GetConfig()

	s.ctrl = gomock.NewController(s.T())

	s.airQualityService = mocks.NewMockAirQualityService(s.ctrl)

	s.airQualityHandler = handler.NewAirQualityHandler(
		s.cfg,
		s.airQualityService,
	)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (s *HandlerTestSuite) TestHandlerRun() {
	srv := s.airQualityHandler.Run()

	err := srv.Shutdown(context.TODO())

	assert.NoError(s.T(), err)
}

func (s *HandlerTestSuite) TestHandlerGetAirQuality() {
	s.airQualityService.EXPECT().GetAirQuality().Return(
		&models.AirQuality{},
	)

	req := httptest.NewRequest(http.MethodGet, "/air-quality", nil)
	w := httptest.NewRecorder()

	s.airQualityHandler.AirQualityHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "{\"sensors\":null,\"current_time\":\"0001-01-01T00:00:00Z\"}", string(data))
}

func (s *HandlerTestSuite) TestHandlerSingleRead() {
	s.airQualityService.EXPECT().SingleRead().Return(
		&models.AirQuality{},
	)

	req := httptest.NewRequest(http.MethodGet, "/single-read", nil)
	w := httptest.NewRecorder()

	s.airQualityHandler.SingleReadHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "{\"sensors\":null,\"current_time\":\"0001-01-01T00:00:00Z\"}", string(data))
}

func (s *HandlerTestSuite) TestHandlerStart() {
	s.airQualityService.EXPECT().Start()

	req := httptest.NewRequest(http.MethodPost, "/start", nil)
	w := httptest.NewRecorder()

	s.airQualityHandler.StartHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "{}", string(data))
}

func (s *HandlerTestSuite) TestHandlerStop() {
	s.airQualityService.EXPECT().Stop()

	req := httptest.NewRequest(http.MethodPost, "/stop", nil)
	w := httptest.NewRecorder()

	s.airQualityHandler.StopHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "{}", string(data))
}
