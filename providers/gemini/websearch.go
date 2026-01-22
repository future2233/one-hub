package gemini

import "strings"

const webSearchModelSuffix = "-websearch"

func parseWebSearchModel(model string) (effectiveModel string, enableWebSearch bool) {
	if !strings.HasSuffix(model, webSearchModelSuffix) {
		return model, false
	}

	effectiveModel = strings.TrimSuffix(model, webSearchModelSuffix)
	if effectiveModel == "" {
		return model, false
	}

	return effectiveModel, true
}

func ensureGoogleSearchTool(request *GeminiChatRequest) {
	for _, tool := range request.Tools {
		if tool.GoogleSearch != nil {
			return
		}
	}

	request.Tools = append(request.Tools, GeminiChatTools{
		GoogleSearch: &GeminiCodeExecution{},
	})
}
