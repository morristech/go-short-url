
## Currency Exchange Server - Golang 

[![Build Status](https://travis-ci.org/me-io/go-short-url.svg?branch=master)](https://travis-ci.org/me-io/go-short-url)
[![Go Report Card](https://goreportcard.com/badge/github.com/me-io/go-short-url)](https://goreportcard.com/report/github.com/me-io/go-short-url)
[![Coverage Status](https://coveralls.io/repos/github/me-io/go-short-url/badge.svg?branch=master)](https://coveralls.io/github/me-io/go-short-url?branch=master)
[![GoDoc](https://godoc.org/github.com/me-io/go-short-url?status.svg)](https://godoc.org/github.com/me-io/go-short-url)
[![GitHub release](https://img.shields.io/github/release/me-io/go-short-url.svg)](https://github.com/me-io/go-short-url/releases)


[![Blog URL](https://img.shields.io/badge/Author-blog-green.svg?style=flat-square)](https://meabed.com)
[![COMMIT](https://images.microbadger.com/badges/commit/meio/go-short-url-server.svg)](https://microbadger.com/images/meio/go-short-url-server)
[![SIZE-LAYERS](https://images.microbadger.com/badges/image/meio/go-short-url-server.svg)](https://microbadger.com/images/meio/go-short-url-server)
[![Pulls](https://shields.beevelop.com/docker/pulls/meio/go-short-url-server.svg?style=flat-square)](https://hub.docker.com/r/meio/go-short-url-server)

## TODO LIST
- [x] Shorten method
- [x] Redirect / method
- [ ] Option to request attributes "ip range, user agent"
- [ ] Throttling and rate limiting " on creation and redirect " 
- [ ] Expiry time
- [ ] add custom params to the redirect url " placeholder "
- [ ] Web hooks
- [ ] Admin 
- [ ] error structure for empty json or regex not matched
- [ ] convert panic to api json error
- [ ] increase tests
- [ ] verbose logging
- [ ] godoc 
- [ ] static bundle public folder `./cmd/server/public`
- [ ] v 1.0.0 release ( docker / binary github / homebrew mac )
- [ ] support historical rates if possible
- [ ] benchmark & performance optimization ` memory leak`
- [ ] contributors list 

## Contributing

Anyone is welcome to [contribute](CONTRIBUTING.md), however, if you decide to get involved, please take a moment to review the guidelines:

* [Only one feature or change per pull request](CONTRIBUTING.md#only-one-feature-or-change-per-pull-request)
* [Write meaningful commit messages](CONTRIBUTING.md#write-meaningful-commit-messages)
* [Follow the existing coding standards](CONTRIBUTING.md#follow-the-existing-coding-standards)


## License

The code is available under the [MIT license](LICENSE.md).
