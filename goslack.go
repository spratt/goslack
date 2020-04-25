package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	appName = "goslack"
	version = "v0.0.0"

	TOKEN_ENVVAR = "TOKEN"
	CHANNEL_ENVVAR = "CHANNEL"
)

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

var text = flag.String("text", "Hello from goslack", "The text you want to send")

func main() {
	flag.Parse()
	fmt.Printf("%s %s\n", appName, version)

	token := os.Getenv(TOKEN_ENVVAR)
	channel := os.Getenv(CHANNEL_ENVVAR)

	// https://golang.org/pkg/net/url/#Values
	vals := url.Values{}
	vals.Add("token", token)
	vals.Add("channel", channel)
	vals.Add("text", *text)
	vals.Add("as_user", "true")
	encodedVals := vals.Encode()
	fmt.Println(encodedVals)

	// https://golang.org/pkg/net/http/#NewRequest
	// https://api.slack.com/methods/chat.postMessage
	req, err := http.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage?" + encodedVals,
		nil,
	)
	maybePanic(err)

	// https://golang.org/pkg/net/http/#Request
	// https://golang.org/pkg/net/http/#Header
	// https://golang.org/pkg/os/#Getenv
	req.Header.Add("Authentication", token)
	req.Header.Add("Content-Type", "text/plain; charset=utf-8")
	
	fmt.Printf("%+v\n\n", req)

	res, err := http.DefaultClient.Do(req)
	maybePanic(err)
	fmt.Printf("%+v\n\n", res)

	resBytes, err := ioutil.ReadAll(res.Body)
	maybePanic(err)
	res.Body.Close()
	fmt.Println(string(resBytes))
}
