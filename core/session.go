package core

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func ImportSession(f *os.File) []*http.Cookie {
	decoder := json.NewDecoder(f)

	var cookies []*http.Cookie
	if err := decoder.Decode(&cookies); err != nil {
		log.Fatal(err)
		return nil
	}
	return cookies
}

func ExportSession(f *os.File, cookies []*http.Cookie) error {
	data, err := json.Marshal(cookies)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(string(data)); err != nil {
		return err
	}
	return nil
}
