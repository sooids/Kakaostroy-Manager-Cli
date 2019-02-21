package core

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/rs/xid"
	"github.com/sooid/Kakaostory-Manager-Cli/core/models"
	"golang.org/x/net/publicsuffix"
)

type KakaoRequest struct {
	Jar     *cookiejar.Jar
	Session []*http.Cookie
}

func NewKakaoRequest() *KakaoRequest {
	return new(KakaoRequest)
}

func (kr *KakaoRequest) SetSession(cookie []*http.Cookie) {
	url, _ := url.Parse("https://story.kakao.com/")
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	cookieJar.SetCookies(url, cookie)

	kr.Session = cookie
	kr.Jar = cookieJar
}

func (kr *KakaoRequest) ClearSession() {
	kr.Session = nil
	kr.Jar = nil
}

func (kr *KakaoRequest) Request(URL string, Method string, obj interface{}) error {
	req, err := http.NewRequest(Method, URL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-Kakao-ApiLevel", "46")
	req.Header.Set("X-Kakao-VC", xid.New().String())
	req.Header.Set("X-Kakao-DeviceInfo", "web:d;-;-")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	req.Header.Set("DNT", "1")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "ko")
	req.Header.Set("Cache-Control", "no-cache")

	req.Header.Set("Referer", "https://story.kakao.com/")
	req.Header.Set("Host", "story.kakao.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36")

	client := &http.Client{Jar: kr.Jar}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := json.NewDecoder(reader).Decode(obj); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// FIX THIS!
func (kr *KakaoRequest) GetProfile(id string) *models.GetProfile {
	profile := new(models.GetProfile)
	if err := kr.Request(APIS["PROFILE"].URL+id, APIS["PROFILE"].Method, profile); err != nil {
		log.Fatal(err)
		return nil
	}

	return profile
}

func (kr *KakaoRequest) GetFeeds() *models.GetFeed {
	feeds := new(models.GetFeed)
	if err := kr.Request(APIS["FEEDS"].URL, APIS["FEEDS"].Method, feeds); err != nil {
		log.Fatal(err)
		return nil
	}

	return feeds
}
