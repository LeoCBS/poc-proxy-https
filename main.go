package main

import (
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	proxy    string
	user     string
	password string
	dest     string
)

func main() {

	flag.StringVar(&proxy, "proxy", "", "provide proxy URL: scheme://ip:port")
	flag.StringVar(&user, "user", "", "provide proxy user")
	flag.StringVar(&password, "password", "", "provide proxy password")
	flag.StringVar(&dest, "dest", "", "provide URL to access")
	flag.Parse()

	parsedProxy := strings.Split(proxy, "://")
	proxyScheme := parsedProxy[0]
	proxyHost := parsedProxy[1]

	fmt.Printf("scheme: %s\n", proxyScheme)
	fmt.Printf("host: %s\n", proxyHost)

	req, _ := http.NewRequest("GET", dest, nil)
	proxyURL := url.URL{
		Scheme: proxyScheme,
		Host:   proxyHost}

	transport := &http.Transport{
		Proxy:           http.ProxyURL(&proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	if user != "" && password != "" {
		fmt.Println("Setting basic auth")
		auth := fmt.Sprintf("%s:%s", user, password)
		basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Add("Proxy-Authorization", basic)
	}

	fmt.Println("making request")
	resp, err := client.Do(req)
	fmt.Println("done")

	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("code: %s\n", resp.StatusCode)
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s reading response\n", err)
		return
	}

	fmt.Println(os.Stdout, string(htmlData))
}
