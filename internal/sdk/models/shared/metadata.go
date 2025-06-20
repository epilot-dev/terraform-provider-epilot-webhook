// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// Metadata - Contains the metadata about the configured event
type Metadata struct {
	// Action that triggered the event
	Action *string `json:"action,omitempty"`
	// When triggered by an automation this is the id of the action
	ActionID *string `json:"action_id,omitempty"`
	// The name of the automation that triggered the event
	AutomationName *string `json:"automation_name,omitempty"`
	// ID used to track the event
	CorrelationID *string `json:"correlation_id,omitempty"`
	// Time of event creation
	CreationTimestamp *string `json:"creation_timestamp,omitempty"`
	// When triggered by an automation this is its execution id
	ExecutionID *string `json:"execution_id,omitempty"`
	// The ID of the given organization
	OrganizationID string `json:"organization_id"`
	// Origin of the event
	Origin *string `json:"origin,omitempty"`
	// The ID of the user for manual triggered events
	UserID *string `json:"user_id,omitempty"`
	// The ID of the webhook configuration
	WebhookID *string `json:"webhook_id,omitempty"`
	// The name of the webhook configuration
	WebhookName *string `json:"webhook_name,omitempty"`
}

func (o *Metadata) GetAction() *string {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *Metadata) GetActionID() *string {
	if o == nil {
		return nil
	}
	return o.ActionID
}

func (o *Metadata) GetAutomationName() *string {
	if o == nil {
		return nil
	}
	return o.AutomationName
}

func (o *Metadata) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *Metadata) GetCreationTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.CreationTimestamp
}

func (o *Metadata) GetExecutionID() *string {
	if o == nil {
		return nil
	}
	return o.ExecutionID
}

func (o *Metadata) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *Metadata) GetOrigin() *string {
	if o == nil {
		return nil
	}
	return o.Origin
}

func (o *Metadata) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *Metadata) GetWebhookID() *string {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

func (o *Metadata) GetWebhookName() *string {
	if o == nil {
		return nil
	}
	return o.WebhookName
}
