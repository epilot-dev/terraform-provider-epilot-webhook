package provider

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/epilot-dev/terraform-provider-epilot-webhook/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestToSharedWebhookConfigInput_EmptyJsonataExpression(t *testing.T) {
	tests := []struct {
		name           string
		jsonataExpr    types.String
		expectNil      bool
	}{
		{
			name:        "null jsonata expression should be nil",
			jsonataExpr: types.StringNull(),
			expectNil:   true,
		},
		{
			name:        "unknown jsonata expression should be nil",
			jsonataExpr: types.StringUnknown(),
			expectNil:   true,
		},
		{
			name:        "empty string jsonata expression should be nil",
			jsonataExpr: types.StringValue(""),
			expectNil:   true,
		},
		{
			name:        "valid jsonata expression should not be nil",
			jsonataExpr: types.StringValue("$.entity.name"),
			expectNil:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := &WebhookResourceModel{
				Name:              types.StringValue("test-webhook"),
				EventName:         types.StringValue("EntityCreated"),
				JsonataExpression: tt.jsonataExpr,
			}

			result, diags := model.ToSharedWebhookConfigInput(context.Background())
			if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}

			if tt.expectNil && result.JsonataExpression != nil {
				t.Errorf("expected JsonataExpression to be nil, got %q", *result.JsonataExpression)
			}
			if !tt.expectNil && result.JsonataExpression == nil {
				t.Error("expected JsonataExpression to not be nil, got nil")
			}
			if !tt.expectNil && result.JsonataExpression != nil && *result.JsonataExpression != tt.jsonataExpr.ValueString() {
				t.Errorf("expected JsonataExpression to be %q, got %q", tt.jsonataExpr.ValueString(), *result.JsonataExpression)
			}
		})
	}
}

func TestRefreshThenUpdate_EmptyJsonataNotSentInJSON(t *testing.T) {
	// Simulate what happens when API returns empty string for jsonataExpression
	// then we update the webhook - the empty string should NOT appear in the JSON payload

	// Step 1: Simulate API response with empty jsonataExpression
	emptyStr := ""
	apiResponse := &shared.WebhookConfig{
		Name:              "test-webhook",
		EventName:         "EntityCreated",
		JsonataExpression: &emptyStr, // API returns empty string
	}

	// Step 2: Refresh model from API response
	model := &WebhookResourceModel{}
	diags := model.RefreshFromSharedWebhookConfig(context.Background(), apiResponse)
	if diags.HasError() {
		t.Fatalf("refresh failed: %v", diags)
	}

	// Step 3: Convert back to API request (as would happen during update)
	result, diags := model.ToSharedWebhookConfigInput(context.Background())
	if diags.HasError() {
		t.Fatalf("conversion failed: %v", diags)
	}

	// Step 4: Verify jsonataExpression is nil (won't be sent due to omitempty)
	if result.JsonataExpression != nil {
		t.Errorf("expected JsonataExpression to be nil after refresh from empty string, got %q", *result.JsonataExpression)
	}

	// Step 5: Verify JSON serialization doesn't include jsonataExpression
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}

	if strings.Contains(string(jsonBytes), "jsonataExpression") {
		t.Errorf("JSON should not contain jsonataExpression field, got: %s", string(jsonBytes))
	}
}
