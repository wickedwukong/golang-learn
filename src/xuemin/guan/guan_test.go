package guan

import "testing"

func TestGreeting(t *testing.T) {
	greeting := Greeting()
	if (greeting != "Hello") {
		t.Error("expected Hello, got ", greeting)
	}
}