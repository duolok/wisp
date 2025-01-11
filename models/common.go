package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

const (
	llama3        = "llama3"
	embedEnglish  = "embedEnglish"
	titanG1Lite   = "titanG1Lite"
	Llama3modelId = "eu.meta.llama3-2-1b-instruct-v1:0"
)

func CallStreamingOutputFunction(llm string, output *bedrockruntime.InvokeModelWithResponseStreamOutput, handler StreamingOutputHandler) error {
	switch llm {
	case llama3:
		err := ProcessLlamaStreamingOutput(output, handler)
		if err != nil {
			return err
		}

	case embedEnglish:
		// TODO

	default:
		return fmt.Errorf("unknown llm value: %s", llm)

	}

	return nil
}
