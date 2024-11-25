package models

type Llama3Request struct {
	Prompt       string  `json:"prompt"`
	MaxGenLength string  `json:"max_gen_len,omitempty`
	Temperature  float64 `json:"max_gen_len,omitempty`
}

type Llama3Response struct {
    Generation string `json:"generation"`
}

func (r Llama3Response) SetContent(content string) {
    r.Generation = content
}

func (r Llama3Response) GetContent() string {
    return r.Generation
}
