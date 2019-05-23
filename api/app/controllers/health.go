package controllers

import (
	"api/app/controllers/core"

	define "api/app/util/define"
	log "api/app/util/log"
	redis "api/app/util/redis"

	"net/http"

	"github.com/revel/revel"
)

type HealthResponse struct {
	Code int `json:"code"`
}

type HealthController struct {
	core.BaseController
}

type RedisData struct {
	value int
	param string
}

func (c HealthController) Index() revel.Result {
	log.Println("Method:", c.Request.Method)

	log.Println("RemoteAddr:", c.Request.RemoteAddr)

	if c.Request.Method == "OPTIONS" {
		c.Response.Status = http.StatusOK
		return c.RenderText("")
	}

	res := HealthResponse{}

	res.Code = define.SUCCESS
	c.Response.Status = http.StatusOK

	// if redis.Get(37, "health") > 5 {
	// 	res.Code = define.SUCCESS
	// 	c.Response.Status = http.StatusTooManyRequests
	// }

	// redis.Limit(37, "health")

	return c.RenderJSON(res)
}

func (c HealthController) Index2() revel.Result {
	log.Println("Method:", c.Request.Method)

	log.Println("RemoteAddr:", c.Request.RemoteAddr)

	if c.Request.Method == "OPTIONS" {
		c.Response.Status = http.StatusOK
		return c.RenderText("")
	}

	res := HealthResponse{}

	res.Code = define.SUCCESS
	c.Response.Status = http.StatusOK

	if redis.GetIp(c.Request, "health") > 5 {
		res.Code = define.SUCCESS
		c.Response.Status = http.StatusTooManyRequests
		return c.RenderJSON(res)
	}

	redis.LimitIp(c.Request, "health")

	return c.RenderJSON(res)
}
