package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func updateRemoteIp(param CMDParam) {
	var defaultTransport http.RoundTripper = &http.Transport{
		Proxy: nil,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          30,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{Transport: defaultTransport}

	// [1] get local ip adress
	resp, err := client.Get("http://ip4only.me/api/")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	_, _ = reader.ReadString(',')

	ip, _ := reader.ReadString(',')
	ipv4 := ip[0 : len(ip)-1]

	// log.Printf("IP address:%s\n", ipv4)

	// [2] set remote dns ip
	HttpsUpdateUrl := fmt.Sprintf("https://dynamicdns.park-your-domain.com/update?host=%s&domain=%s&password=%s&ip=%s", param.Host, param.Domain, param.Password, ipv4)
	// log.Println(HttpsUpdateUrl)

	_, err = http.Get(HttpsUpdateUrl)
	if err != nil {
		log.Printf("http get remote url %s failed\n", HttpsUpdateUrl)
	}
	log.Println("updating remote dns ip address:", ipv4)
}
