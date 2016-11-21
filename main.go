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
)

var (
	proxy    string
	user     string
	password string
	dest     string
)

func main() {

	flag.StringVar(&proxy, "proxy", "", "provide proxy URL: IP:PORT")
	flag.StringVar(&user, "user", "", "provide proxy user")
	flag.StringVar(&password, "password", "", "provide proxy password")
	flag.StringVar(&dest, "dest", "", "provide URL to access")
	flag.Parse()

	req, _ := http.NewRequest("GET", dest, nil)
	req.Header.Set("Host", "www.google.com.br")

	proxyURL := url.URL{
		Scheme: "http",
		Host:   proxy}

	auth := fmt.Sprintf("%s:%s", user, password)
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Proxy-Authorization", basic)

	transport := &http.Transport{
		Proxy:           http.ProxyURL(&proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//transport.ProxyConnectHeader = req.Header
	client := &http.Client{Transport: transport}
	req.RequestURI = ""

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("erro: %s", err)
		return
	}
	fmt.Printf("code: %s", resp.StatusCode)
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(os.Stdout, string(htmlData))
}
