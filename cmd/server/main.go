package main

import (
	"flag"
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	host *string
	port *int
	// Logger ... Logger Driver
	Logger = logging.MustGetLogger("go-short-url-server")

	format = logging.MustStringFormatter(
		`%{color}%{time:2006-01-02T15:04:05.999999} %{shortfunc} ▶ %{level:.8s} %{id:03x}%{color:reset} %{message}`,
	)

	routes = map[string]func(w http.ResponseWriter, r *http.Request){
		`/shorten`:  Shorten,
		`/r/`: Redirect,
	}
	_, filename, _, _ = runtime.Caller(0)
	defaultStaticPath = filepath.Dir(filename) + `/public`
	staticPath        = &defaultStaticPath
)

// init ... init function of the server
func init() {
	// Logging
	backendStderr := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatted := logging.NewBackendFormatter(backendStderr, format)
	// Only DEBUG and more severe messages should be sent to backend1
	backendLevelFormatted := logging.AddModuleLevel(backendFormatted)
	backendLevelFormatted.SetLevel(logging.DEBUG, "")
	// Set the backend to be used.
	logging.SetBackend(backendLevelFormatted)

	// Caching
	host = flag.String(`H`, `0.0.0.0`, `Host binding address`)
	port = flag.Int(`P`, 5000, `Host binding port`)
	staticPath = flag.String(`STATIC_PATH`, defaultStaticPath, `Webserver static path`)

	flag.Parse()
}

// main ... main function start the server
func main() {

	Logger.Infof("host %s", *host)
	Logger.Infof("port %d", *port)
	Logger.Infof("Static dir %s", *staticPath)

	// handle routers
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	go serveHTTP(*host, *port)
	select {}
}

// serveHTTP ... initiate the HTTP Server
func serveHTTP(host string, port int) {

	mux := http.NewServeMux()
	for k, v := range routes {
		mux.HandleFunc(k, v)
	}

	handleStatic(mux)

	addr := fmt.Sprintf("%v:%d", host, port)
	server := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	Logger.Infof("Server Started @ %v:%d", host, port)

	err := server.ListenAndServe()
	Logger.Error(err.Error())
}

func handleStatic(mux *http.ServeMux) {
	mux.Handle(`/static`, http.FileServer(http.Dir(*staticPath)))
}
