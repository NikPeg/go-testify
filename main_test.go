package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
    "github.com/stretchr/testify/require"
)


func TestCorrectRequest(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?city=moscow&count=" + strconv.Itoa(totalCount), nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusOK)

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    require.Equal(t, len(list), totalCount)
}

func TestWrongCityValue(t *testing.T) {
    wrongCity := "karaganda"
    req := httptest.NewRequest("GET", "/cafe?city=" + wrongCity + "&count=1", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

    body := responseRecorder.Body.String()

    require.NotNil(t, body)
    require.Equal(t, body, "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    req := httptest.NewRequest("GET", "/cafe?city=moscow&count=" + strconv.Itoa(totalCount + 1), nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    require.Equal(t, responseRecorder.Code, http.StatusOK)

    body := responseRecorder.Body.String()
    list := strings.Split(body, ",")

    require.Equal(t, len(list), totalCount)
}
