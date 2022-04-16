package controllers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type ResponseSimulator struct {
}

func (r ResponseSimulator) AllSuccessSimulator(ctx *gin.Context) {
	ctx.Status(200)
}

func (r ResponseSimulator) InternalServerErrorSimulator(ctx *gin.Context) {
	ctx.Status(500)
}

func (r ResponseSimulator) SlowResponseSimulator(ctx *gin.Context) {
	randomNumber := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(randomNumber))
	ctx.Status(200)
}
