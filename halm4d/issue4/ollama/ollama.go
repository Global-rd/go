package ollama

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Ollama interface {
	PullModel() error
	Generate(prompt string) (string, error)
}

type OllamaRequest struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Options map[string]interface{} `json:"options"`
}

type OllamaGenerateResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	Context            []int     `json:"context"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int       `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}

type OllamaPullResponse struct {
	Status string `json:"status"`
}

func NewOllama(baseUrl, model string, verbose bool) Ollama {
	return &OllamaClient{
		url:     fmt.Sprintf("%s/api", baseUrl),
		model:   model,
		verbose: verbose,
	}
}

type OllamaClient struct {
	url     string
	model   string
	verbose bool
}

func (o *OllamaClient) PullModel() error {
	ollamaRequest := OllamaRequest{
		Model: o.model,
	}

	jsonData, err := json.Marshal(ollamaRequest)
	if err != nil {
		return errors.Join(err, fmt.Errorf("error marshalling JSON"))
	}

	resp, err := http.DefaultClient.Post(
		fmt.Sprintf("%s/pull", o.url),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		errors.Join(err, fmt.Errorf("error making request"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Join(err, fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	decoder := json.NewDecoder(resp.Body)
	for {
		var pullResponse OllamaPullResponse
		if err := decoder.Decode(&pullResponse); err == io.EOF {
			break
		} else if err != nil {
			return errors.Join(err, fmt.Errorf("error decoding response body"))
		}
		if o.verbose {
			fmt.Println("Response chunk:", pullResponse.Status)
		}
	}
	return nil
}

func (o *OllamaClient) Generate(prompt string) (string, error) {
	// ollama pull llama3.2
	jsonData, err := json.Marshal(OllamaRequest{
		Model:  o.model,
		Prompt: prompt,
	})
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("error marshalling JSON"))
	}

	resp, err := http.DefaultClient.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", errors.Join(err, fmt.Errorf("error making request"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Join(err, fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	decoder := json.NewDecoder(resp.Body)
	var generateResponse OllamaGenerateResponse
	response := ""
	for {
		if err := decoder.Decode(&generateResponse); err == io.EOF {
			break
		} else if err != nil {
			return "", errors.Join(err, fmt.Errorf("error decoding resp body"))
		}
		if o.verbose {
			fmt.Printf("Response chunk done: %v, Partial response: %s\n", generateResponse.Done, generateResponse.Response)
		}
		response += generateResponse.Response
	}

	if generateResponse.Done {
		return response, nil
	} else {
		return "", errors.Join(err, fmt.Errorf("response not done"))
	}
}
