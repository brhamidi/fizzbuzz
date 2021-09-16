package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/brhamidi/fizzbuzz/pkg/mock"
)

func TestHandler_GetStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)
	mockLogger := mock.NewMockLogger(ctrl)

	t.Run("should return 500 and log an error if storage operation fail", func(t *testing.T) {
		mockStorage.EXPECT().Max().Return("", 0, errors.New("mock"))
		mockLogger.EXPECT().Error(gomock.Any())

		router := NewServer(gin.TestMode, mockStorage, mockLogger)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, StatsRoute, nil)
		router.ServeHTTP(w, req)

		expectedBody, _ := json.Marshal(newResponseError(errUserInternal))

		assert.Equal(t, string(expectedBody), w.Body.String())
		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("should return http status code 204 no content if fuzzbuzz has never been called", func(t *testing.T) {
		mockStorage.EXPECT().Max().Return("", 0, nil)

		router := NewServer(gin.TestMode, mockStorage, mockLogger)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, StatsRoute, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Result().StatusCode)
	})

	t.Run("should be ok", func(t *testing.T) {
		mockStorage.EXPECT().Max().Return("key", 42, nil)

		router := NewServer(gin.TestMode, mockStorage, mockLogger)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, StatsRoute, nil)
		router.ServeHTTP(w, req)

		resp := fmt.Sprintf(templateStatsResponse, 42, "key")
		expectedBody, _ := json.Marshal(&ResponseSuccess{resp})

		assert.Equal(t, string(expectedBody), w.Body.String())
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestHandler_DeleteStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)
	mockLogger := mock.NewMockLogger(ctrl)

	t.Run("should return 500 and log an error if storage operation fail", func(t *testing.T) {
		mockStorage.EXPECT().Reset().Return(errors.New("mock"))
		mockLogger.EXPECT().Error(gomock.Any())

		router := NewServer(gin.TestMode, mockStorage, mockLogger)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodDelete, StatsRoute, nil)
		router.ServeHTTP(w, req)

		expectedBody, _ := json.Marshal(newResponseError(errUserInternal))

		assert.Equal(t, string(expectedBody), w.Body.String())
		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("should be ok", func(t *testing.T) {
		mockStorage.EXPECT().Reset()

		router := NewServer(gin.TestMode, mockStorage, mockLogger)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodDelete, StatsRoute, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}
