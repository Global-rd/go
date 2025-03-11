package provider

import (
	"csv-writer/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	OLLAMA_MODEL   = "llama3.2"
	USE_STREAM     = false
	CHECK_REQ      = "who wrote the Hitchhiker's Guide to the Galaxy?"
	CHECK_EXPECTED = "Douglas Adams"
	DATA_REQUEST   = "generate a json list of 21 elements, each containing a field firstName, lastNAme, userName, phone and email. return only the generated json nothing else"
)

type OllamaProvider struct {
	Url string
}

type OllamaRequest struct {
	Model  string `json:"model"`  // Field for model
	Stream bool   `json:"stream"` // Field for stream
	Prompt string `json:"prompt"`
}

func (o *OllamaProvider) CheckSource() error {
	resp, err := o.sendRequest(CHECK_REQ)
	if err != nil {
		return err
	}
	if !strings.Contains(resp, CHECK_EXPECTED) {
		return fmt.Errorf("wrong answer from ollama, data source not reliable")
	}
	return nil
}

func convertToSafeJsonString(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	return strings.ReplaceAll(str, "\"", `"`)
}

func (o *OllamaProvider) GetData() ([]model.JsonData, error) {
	respStr, err := o.sendRequest(DATA_REQUEST)
	if err != nil {
		return nil, err
	}
	respStr = convertToSafeJsonString(respStr)
	var data []model.JsonData
	data, err = model.ParseJsonData(respStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (o *OllamaProvider) sendRequest(question string) (string, error) {
	req := OllamaRequest{
		Model:  OLLAMA_MODEL,
		Stream: USE_STREAM,
		Prompt: question,
	}
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	reader := strings.NewReader(string(jsonReq))
	var resp *http.Response
	resp, err = http.Post(o.Url, "application/json", reader)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to get response from ollama, status code: %d", resp.StatusCode)
	}
	var respBody []byte
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Join(fmt.Errorf("failed to read response from ollama"), err)
	}
	var respJsonBody map[string]interface{}
	err = json.Unmarshal(respBody, &respJsonBody)
	if err != nil {
		return "", err
	}
	ollamaRespStr, ok := respJsonBody["response"].(string)
	if !ok {
		return "", errors.New("failed to parse answer from ollama")
	}
	return ollamaRespStr, nil
}

func NewOllamaProvider(url string) *OllamaProvider {
	return &OllamaProvider{Url: url}
}
