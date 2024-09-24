// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type EventConfigEntry struct {
	// Either a user friendly label, or the eventName itself.
	EventLabel *string `json:"eventLabel,omitempty"`
	// Name for identifying the event. Unique.
	EventName *string `json:"eventName,omitempty"`
}

func (o *EventConfigEntry) GetEventLabel() *string {
	if o == nil {
		return nil
	}
	return o.EventLabel
}

func (o *EventConfigEntry) GetEventName() *string {
	if o == nil {
		return nil
	}
	return o.EventName
}