package model

/*
{"id":"cmpl-6yNKfVWeSRoQj7fNujpxe39ZEqK1R","object":"text_completion",
"created":1679847389,"model":"text-davinci-002",
"choices":[{"text":"\n\nHello, how are you?","index":0,"logprobs":null,"finish_reason":"stop"}],
"usage":{"prompt_tokens":5,"completion_tokens":8,"total_tokens":13}}
*/

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
