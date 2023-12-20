package suga

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	random_ua_api = "https://user-agents.net/random"
	user_agents   = []string{}
)

func parseWebPage() error {
	response, err := http.Get(random_ua_api)
	if err != nil {
		return errors.New("error on parse")
	}
	defer response.Body.Close()

	query, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return errors.New("error on parse")
	}

	query.Find("ol li a").Each(func(_ int, user_agent *goquery.Selection) {
		user_agents = append(user_agents, user_agent.Text())
	})
	return nil
}

func Refresh() {
	// Emptying the user_agents to get new ones
	user_agents = user_agents[:0]
}

func GetUserAgent() (string, error) {
	if len(user_agents) == 0 {
		err := parseWebPage()
		if err != nil {
			return "", err
		}
	}
	return user_agents[rand.Intn(len(user_agents))], nil
}
