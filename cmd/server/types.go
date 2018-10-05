package main

// shortenReqObj ... Shorten Request Data Object
type shortenReqObj struct {
	Url string `json:"url"`
}

// shortenResObj ... Shorten Response Data Object
type shortenResObj struct {
	Link      string `json:"link"`
	LongUrl   string `json:"long_url"`
	ExpireAt  string `json:"expire_at"`
	CreatedAt string `json:"created_at"`
}
