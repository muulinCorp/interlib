package health

import (
	"net/http"
	"time"

	apitool "github.com/94peter/api-toolkit"
	"github.com/94peter/api-toolkit/errors"
	"github.com/94peter/morm/conn"
	"github.com/gin-gonic/gin"
)

func NewAPI() apitool.GinAPI {
	return &healthAPI{}
}

type healthAPI struct {
	errors.CommonApiErrorHandler
}

func (a *healthAPI) GetAPIs() []*apitool.GinApiHandler {
	return []*apitool.GinApiHandler{
		// health check
		{Method: "GET", Path: "__health", Handler: a.healthHandler, Auth: false},
	}
}

func (a *healthAPI) healthHandler(c *gin.Context) {
	healthResp := healthResponse{
		Status: "ok",
	}
	dbclt := conn.GetMgoDbConnFromGin(c)
	err := dbclt.Ping()
	if err != nil {
		healthResp.Connection.Mongo.Status = "red"
		healthResp.Connection.Mongo.Msg = err.Error()
	} else {
		healthResp.Connection.Mongo.Status = "green"
		healthResp.Connection.Mongo.Msg = "ok"
	}
	healthResp.Now = time.Now()
	c.JSON(http.StatusOK, healthResp)
}

type healthResponse struct {
	Now        time.Time       `json:"now"`
	Status     string          `json:"status"`
	Connection connectionState `json:"connection"`
}

type connectionState struct {
	Mongo struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	} `json:"mongo"`
}
