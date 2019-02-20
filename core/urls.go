package core

import "net/http"

type API struct {
	URL    string
	Method string
}

var APIS map[string]API

func InitAPIS() {
	APIS = map[string]API{
		"PROFILE": API{"https://story.kakao.com/a/profiles/", http.MethodGet},
		"FEEDS":   API{"https://story.kakao.com/a/feeds/", http.MethodGet},
		// "PROFILE": API{"https://story.kakao.com/a/profiles/", http.MethodGet},
	}
}
