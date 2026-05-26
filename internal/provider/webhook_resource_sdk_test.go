package provider

import (
	"context"
	"testing"

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
