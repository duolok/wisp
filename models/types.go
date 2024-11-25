package models

import (
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type ModelWrapper struct {
    BedrockRuntimeClient *bedrockruntime.Client
}
