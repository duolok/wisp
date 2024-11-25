package models

import (
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type ModelWrapper struct {
	BedrockRuntimeClient *bedrockruntime.Client
}

type LLM interface {
	Invoke() (string, error)
	Stream() (*bedrockruntime.InvokeModelWithResponseStreamOutput, error)
}

type LLMPrompt struct {
	bedrock ModelWrapper
	prompt  string
}

type Llama struct {
    LLMPrompt
}

type EmbedEnglish struct {
    LLMPrompt
}
