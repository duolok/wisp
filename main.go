package main

import (
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/duolok/wisp/models"
	"github.com/gorilla/websocket"
)

type MLWrapper struct {
	wrapper models.ModelWrapper
}

func main() {
	cfg := loadConfig()
	brc := bedrockruntime.NewFromConfig(cfg)
	modelWrapper := models.ModelWrapper{BedrockRuntimeClient: brc}
	wrapper := MLWrapper{modelWrapper}
}
