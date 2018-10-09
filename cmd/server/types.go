package main

import "time"

// shortenReqObj ... Shorten Request Data Object
type shortenReqObj struct {
	Url string `json:"url"`
}

// Model: url
type Url struct {
	Link      string    `bson:"link" json:"link"`
	LongUrl   string    `bson:"long_url" json:"long_url"`
	ExpireAt  time.Time `bson:"expire_at" json:"expire_at"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
