package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/a-h/templ"
	"github.com/byitkc/template-golang-templ/components"
	"github.com/byitkc/template-golang-templ/pages"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

const (
	defaultPort = 8080
	defaultHost = "localhost"
)

type ServerConfig struct {
	Host string
	Port uint16
}

func main() {
	parseFlags()
	config, err := readEnv()
	if err != nil {
		panic(err)
	}

	http.Handle("/test", templ.Handler(pages.Home("Test")))
	http.Handle("/", templ.Handler(components.Greeting("Brandon", "Home")))
	serveFile("static/styles/dist/styles.css", "/styles/styles.css")

	host := fmt.Sprintf("%s:%d", config.Host, config.Port)
	glog.Infof("Starting server on %s", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		panic(err)
	}
}

func parseFlags() {
	flag.Parse()
}

func readEnv() (ServerConfig, error) {
	// Loading environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		glog.Infof("Unable to load .env file.")
	}

	var config ServerConfig
	var port int64

	hostString := os.Getenv("SERVER_HOST")
	portString := os.Getenv("SERVER_PORT")

	if hostString == "" {
		return ServerConfig{}, fmt.Errorf("server host not specified, please ensure that a SERVER_HOST is specified in your environment variables")
	}

	if portString == "" {
		return ServerConfig{}, fmt.Errorf("server port not specified, please ensure that a SERVER_PORT is specified in your environment variables")
	}

	port, err = strconv.ParseInt(portString, 10, 64)
	if err != nil {
		glog.Errorf("Error parsing server port number, using the default of %d", os.Getenv("SERVER_PORT"), defaultPort)
		config.Port = 8080
	}

	if port < 1 || port > 65535 {
		panic(fmt.Sprintf("Invalid port specified, must be betwen 1 and 65535 inclusive, you specified: %d.", port))
	}

	config.Host = hostString
	config.Port = uint16(port)

	return config, nil
}

func serveFile(localPath, serverPath string) {
	http.HandleFunc(serverPath, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, localPath)
	})

}
