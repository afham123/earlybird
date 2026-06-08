package scan

import (
	"encoding/json"
	"net/http"
	"strings"

	cfgReader "github.com/americanexpress/earlybird/v4/pkg/config"
)

func isGeminiModel(model string) bool {
	return strings.Contains(strings.ToLower(model), "gemini")
}

func setHeaderGemini(req *http.Request, value string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-goog-api-key", value)
}

func buildGeminiRequestBody(cfg *cfgReader.EarlybirdConfig, scanJob LLMJob, chunk llmFileChunk) ([]byte, error) {

	requestBody := GeminiRequestBody{
		Contents: []GeminiContent{
			{
				Role:  "user",
				Parts: []Parts{{Text: llmSystemPrompt(cfg)}},
			},
			{
				Role:  "user",
				Parts: []Parts{{Text: buildLLMUserPrompt(scanJob, chunk)}},
			},
		},
	}
	return json.Marshal(requestBody)
}

func parseGeminiResponse(body []byte) ([]LLMFinding, error) {
	var response GeminiResponseBody
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return response.Candidates.parts[0], nil
}
