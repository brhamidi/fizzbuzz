package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/brhamidi/fizzbuzz/pkg/mock"
)

func TestHandler_Fizzbuzz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)
	mockLogger := mock.NewMockLogger(ctrl)

	t.Run("should return BadRequest due to invalid input", func(t *testing.T) {
		router := NewServer(gin.TestMode, nil, mockLogger)

		// accesslog
		mockLogger.EXPECT().Info(gomock.Any())

		w := httptest.NewRecorder()

		const queryString = `?int1="42"`
		req, _ := http.NewRequest(http.MethodGet, FizzbuzzRoute+queryString, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("should return BadRequest due to bad input, some parameters are missing", func(t *testing.T) {
		router := NewServer(gin.TestMode, nil, mockLogger)

		// accesslog
		mockLogger.EXPECT().Info(gomock.Any())

		w := httptest.NewRecorder()

		const queryString = `?int1=42&int2=21`
		req, _ := http.NewRequest(http.MethodGet, FizzbuzzRoute+queryString, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("should return 200 but log an storage error", func(t *testing.T) {
		errMock := errors.New("mock")
		mockStorage.EXPECT().Increment("3,5,10,hello,world").Return(errMock)
		mockLogger.EXPECT().Error(gomock.Any())

		// accesslog
		mockLogger.EXPECT().Info(gomock.Any())

		router := NewServer(gin.TestMode, mockStorage, mockLogger)
		w := httptest.NewRecorder()

		const queryString = `?int1=3&int2=5&limit=10&str1=hello&str2=world`
		req, _ := http.NewRequest(http.MethodGet, FizzbuzzRoute+queryString, nil)
		router.ServeHTTP(w, req)

		expectedBody := `{"data":["1","2","hello","4","world","hello","7","8","hello","world"]}`

		assert.Equal(t, expectedBody, w.Body.String())
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("should be ok", func(t *testing.T) {
		mockStorage.EXPECT().Increment("3,5,10,hello,world")

		// accesslog
		mockLogger.EXPECT().Info(gomock.Any())

		router := NewServer(gin.TestMode, mockStorage, mockLogger)
		w := httptest.NewRecorder()

		const queryString = `?int1=3&int2=5&limit=10&str1=hello&str2=world`
		req, _ := http.NewRequest(http.MethodGet, FizzbuzzRoute+queryString, nil)
		router.ServeHTTP(w, req)

		expectedBody := `{"data":["1","2","hello","4","world","hello","7","8","hello","world"]}`

		assert.Equal(t, expectedBody, w.Body.String())
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}
