package providers

import (
	"testing"

	"github.com/sipeed/picoclaw/pkg/config"
)

func TestCreateProvider_LMStudio_NoAPIKey(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.Agents.Defaults.Provider = "lmstudio"
	cfg.Agents.Defaults.Model = "deepseek/deepseek-chat"
	cfg.Providers.LMStudio.APIBase = "http://host.docker.internal:1234/v1"

	provider, err := CreateProvider(cfg)
	if err != nil {
		t.Fatalf("CreateProvider(lmstudio) error = %v", err)
	}

	lmstudio, ok := provider.(*LMStudioProvider)
	if !ok {
		t.Fatalf("CreateProvider(lmstudio) returned %T, want *LMStudioProvider", provider)
	}

	if lmstudio.apiKey != "" {
		t.Fatalf("lmstudio apiKey = %q, want empty", lmstudio.apiKey)
	}
	if lmstudio.apiBase != "http://host.docker.internal:1234/v1" {
		t.Fatalf("lmstudio apiBase = %q, want host.docker.internal endpoint", lmstudio.apiBase)
	}
}

func TestCreateProvider_LMStudio_DefaultBase(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.Agents.Defaults.Provider = "lmstudio"
	cfg.Agents.Defaults.Model = "gpt-oss-20b"

	provider, err := CreateProvider(cfg)
	if err != nil {
		t.Fatalf("CreateProvider(lmstudio default base) error = %v", err)
	}

	lmstudio, ok := provider.(*LMStudioProvider)
	if !ok {
		t.Fatalf("CreateProvider(lmstudio) returned %T, want *LMStudioProvider", provider)
	}

	if lmstudio.apiBase != "http://127.0.0.1:1234/v1" {
		t.Fatalf("lmstudio apiBase = %q, want default local LM Studio base", lmstudio.apiBase)
	}
}
