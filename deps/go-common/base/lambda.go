package base

import "github.com/aws/aws-lambda-go/events"

type IntegratedEvent struct {
	events.APIGatewayProxyRequest
	events.EventBridgeEvent
}

const (
	// EventBridgeEvent is the type of the event
	Unknown int = iota
	APIGateway
	EventBridge
)

func (r *IntegratedEvent) Type() int {
	if r.APIGatewayProxyRequest.Path != "" {
		return APIGateway
	} else if r.EventBridgeEvent.Version != "" {
		return EventBridge
	}

	return Unknown
}
