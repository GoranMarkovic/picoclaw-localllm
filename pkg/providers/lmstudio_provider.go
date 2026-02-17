package providers

// LMStudioProvider is an OpenAI-compatible HTTP provider for LM Studio.
// LM Studio does not require an API key; when apiKey is empty, requests are
// sent without the Authorization header.
type LMStudioProvider struct {
	*HTTPProvider
}

func NewLMStudioProvider(apiKey, apiBase, proxy string) *LMStudioProvider {
	return &LMStudioProvider{HTTPProvider: NewHTTPProvider(apiKey, apiBase, proxy)}
}
