package models

import "errors"

func (wrapper ModelWrapper) LoadModel(modelName string, prompt string) (string, error) {
	switch modelName {
	case llama3:
		llama := Llama{LLMPrompt{wrapper, prompt}}
		response, err := llama.Invoke()
		if err != nil {
			return "", err
		}
		return response, nil
    case embedEnglish:
        // TODO
    case titanG1Lite:
        // TODO

	default:
		return "", errors.New("No such model: " + modelName)

	}

	return "", errors.New("No such model" + modelName)
}
