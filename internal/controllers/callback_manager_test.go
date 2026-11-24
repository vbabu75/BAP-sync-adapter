package controllers

import "testing"

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
