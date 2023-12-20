package suga

import (
	"testing"
)

func TestGetUserAgent(t *testing.T) {
	my_agent := "my super User Agent"
	user_agents = []string{
		my_agent,
	}
	ua, err := GetUserAgent()
	if err != nil {
		t.Error(err)
	}

	if ua != my_agent {
		t.Error("This is not my_agent: ", ua)
	}
}

func TestRefresh(t *testing.T) {
	user_agents = []string{
		"yay it's me",
		"nope it's me",
	}
	Refresh()
	if len(user_agents) != 0 {
		t.Error("This is not refreshed: ", user_agents)
	}
}
