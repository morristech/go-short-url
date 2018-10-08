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

	// generate hash from the given url
	hash := Base64Encode(convertReq.Url)

	shortenRes := &shortenResObj{}
	shortenRes.Link = r.Host + "/r/" + string(hash)
	shortenRes.LongUrl = convertReq.Url
	shortenRes.ExpireAt = time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	shortenRes.CreatedAt = time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")

	shortenJsonVal, err := json.Marshal(shortenRes)
	if err != nil {
		Logger.Panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(shortenJsonVal)
}

var Redirect = func(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.String(), "/")

	redirectUrlHash := urlParts[len(urlParts)-1]

	redirectUrl := Base64Decode(redirectUrlHash)
	fmt.Println(redirectUrl)
	http.Redirect(w, r, redirectUrl, 301)
}
