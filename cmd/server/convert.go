package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Validate ... Validation function for shortenReqObj
func (c *shortenReqObj) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Url, validation.Required, validation.By(func(value interface{}) error {
			_, err := url.ParseRequestURI(value.(string))
			if err != nil {
				return errors.New("invalid url")
			}
			return nil
		})),
	)
}

var Shorten = func(w http.ResponseWriter, r *http.Request) {
	convertReq := &shortenReqObj{}

	if err := json.
		NewDecoder(r.Body).
		Decode(convertReq); err != nil {
		Logger.Panic(err)
	}

	if err := convertReq.Validate(); err != nil {
		Logger.Panic(err)
	}

	// find last generated hash
	result := mgo.FindOne("url", nil)
	lastDoc := &Url{}
	result.Decode(lastDoc)

	// prepare new document
	url := &Url{}
	url.Link = GenerateHash(lastDoc.Link)
	url.LongUrl = convertReq.Url
	url.ExpireAt = time.Now().Add(time.Hour * 24)
	url.CreatedAt = time.Now()

	// insert new documents
	mgo.Create("url", url)

	// update the link for response
	url.Link = fmt.Sprintf("%s/r/%s", r.Host, url.Link)

	shortenJsonVal, err := json.Marshal(url)
	if err != nil {
		Logger.Panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(shortenJsonVal)
}

var Redirect = func(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.String(), "r/")

	redirectLink := urlParts[len(urlParts)-1]

	if redirectLink == "" {
		Logger.Panic("Invalid short url")
	}

	condition := make(map[string]string)
	condition["link"] = redirectLink

	result := mgo.FindOne("url", condition)

	doc := &Url{}
	result.Decode(doc)

	if doc.Link != redirectLink {
		Logger.Panic("Invalid short url")
	}

	// disable redirect cache
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.Redirect(w, r, doc.LongUrl, 301)
}
