package controllers

import (
	"encoding/json"
	"testing"
)

func TestRouteMappingIncludesIONV2Endpoints(t *testing.T) {
	tests := map[string]string{
		"rate":      "on_rate",
		"reconcile": "on_reconcile",
	}

	for forwardRoute, wantCallbackRoute := range tests {
		t.Run(forwardRoute, func(t *testing.T) {
			if got := RouteMapping[forwardRoute]; got != wantCallbackRoute {
				t.Fatalf("RouteMapping[%q] = %q, want %q", forwardRoute, got, wantCallbackRoute)
			}
		})
	}
}

func TestRequestContextCorrelationIDs(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "ION v2 camelCase",
			body: `{"context":{"transactionId":"txn-1","messageId":"msg-1"}}`,
		},
		{
			name: "legacy snake_case",
			body: `{"context":{"transaction_id":"txn-1","message_id":"msg-1"}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var requestContext RequestContext
			if err := json.Unmarshal([]byte(test.body), &requestContext); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}

			transactionID, messageID := requestContext.CorrelationIDs()
			if transactionID != "txn-1" || messageID != "msg-1" {
				t.Fatalf("CorrelationIDs() = (%q, %q), want (%q, %q)", transactionID, messageID, "txn-1", "msg-1")
			}
		})
	}
}
