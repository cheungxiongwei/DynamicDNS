package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// 获取本机公网 ipv4 地址
func GetLocalHostAddress() (string, error) {
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
	resp, err := client.Get("http://ip4only.me/api/")
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	_, _ = reader.ReadString(',')

	ip, _ := reader.ReadString(',')
	ipv4 := ip[0 : len(ip)-1]

	return ipv4, nil
}

func updateRemoteIp(param CMDParam) {

	ipv4, err := GetLocalHostAddress()
	if err != nil {
		log.Printf("http get remote ipv4 failed\n")
		return
	}

	// [2] set remote dns ip
	HttpsUpdateUrl := fmt.Sprintf("https://dynamicdns.park-your-domain.com/update?host=%s&domain=%s&password=%s&ip=%s", param.Host, param.Domain, param.Password, ipv4)
	// log.Println(HttpsUpdateUrl)

	_, err = http.Get(HttpsUpdateUrl)
	if err != nil {
		log.Printf("http get remote url %s failed\n", HttpsUpdateUrl)
	}
	log.Println("updating remote dns ip address:", ipv4)
}
