package main

import (
	"os"
	"testing"
)

func TestEnvVars(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("OPENAI_AUTH_TOKEN", "my-auth-token")
	os.Setenv("OPENAI_BASE_URL", "https://api.example.com/v2")
	os.Setenv("OPENAI_ORG_ID", "my-org-id")
	os.Setenv("OPENAI_MTY_MSG_LIM", "300")

	// Test environment variables has been set correctly
	loadConfigFromEnv()
	if cfg.AuthToken != "my-auth-token" {
		t.Errorf("AuthToken is not set correctly")
	}
	if cfg.BaseURL != "https://api.example.com/v2" {
		t.Errorf("BaseURL is not set correctly")
	}
	if cfg.OrgID != "my-org-id" {
		t.Errorf("OrgID is not set correctly")
	}
	if cfg.EmptyMessagesLimit != 300 {
		t.Errorf("EmptyMessagesLimit is not set correctly")
	}
}
