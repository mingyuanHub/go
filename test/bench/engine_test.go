package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

func BenchmarkEngine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get()
	}
}


//var apiUrl = "https://xiaoming.tradplusad.com/api/v1_2/adconf?v=1_2&sdkv=5.4.6&x=7DF4E3A96B8C3548604371A50F00EEAD&idfa=8818DA59-BBDE-4509-967D-1435D58E89D6&did=8818DA59-BBDE-4509-967D-1435D58E89D6&ddid=0CC00914-2C69-48F8-8C1B-BF847100E802&m=com.oversea.miniworld&iso=VN&z=Asia/Ho_Chi_Minh&l=vi-VN&os=2&appid=2F94F88F6D32AB441554C7E66EF00AC7&app_ver=0.53.11&app_instime=1626603968&device_idfa=0CC00914-2C69-48F8-8C1B-BF847100E802&device_idfv=EAD335F3-EB48-4590-8FB5-791CBDB6E226&device_osv=13.3.1&device_type=1&device_make=Apple&device_model=iPhone&device_contype=2"
//var apiUrl = "http://192.168.1.12:8080/api/v1_2/adconf?v=1.2&sdkv=6.0.0&x=54AC5BA98A4D11FC5D51A8978408ECED&os=1&aid=UID-39b7e296-da0d-4e8a-8172-26fb37bbec07&aaid=28af00fe-08ee-42c9-962b-23c66f24973c&ddid=28af00fe-08ee-42c9-962b-23c66f24973c&device_aaid=28af00fe-08ee-42c9-962b-23c66f24973c&did=UID-39b7e296-da0d-4e8a-8172-26fb37bbec07&br=google&dn=Pixel%203%20XL&m=com.unstall.meetdeleteapp&ct=1&sw=1440&sh=2621&lmt=0&iso=CN&sc=3.5&z=%2B0800&l=zh&o=1&rom=Google&romv=6578210&appid=6640E7E3BDAC951B8F28D4C8C50E50B5&app_ver=2.0&app_instime=1612234395&device_osv=10&device_type=1&device_make=google&device_model=Pixel%203%20XL&device_contype=7"

var apiUrl = "http://172.16.0.84:8080/api/v1_2/adconf?v=1.2&sdkv=6.7.0&x=702208A872E622C1729FC621025D4B1D&os=1&aid=UID-39b7e296-da0d-4e8a-8172-26fb37bbec07&aaid=28af00fe-08ee-42c9-962b-23c66f24973c&ddid=28af00fe-08ee-42c9-962b-23c66f24973c&device_aaid=28af00fe-08ee-42c9-962b-23c66f24973c&did=UID-39b7e296-da0d-4e8a-8172-26fb37bbec07&br=google&dn=Pixel%203%20XL&m=com.unstall.meetdeleteapp&ct=1&sw=1440&sh=2621&lmt=0&iso=CN&sc=3.5&z=%2B0800&l=zh&o=1&rom=Google&romv=6578210&appid=6640E7E3BDAC951B8F28D4C8C50E50B5&app_ver=2.0&app_instime=1612234395&device_osv=10&device_type=1&device_make=google&device_model=Pixel%203%20XL&device_contype=7"

func get() int {

	response, body, err := HttpGetRequest(apiUrl, httpClient1000)
	//fmt.Println("------------ERR START------------")
	//fmt.Println(err)
	//fmt.Println("------------ERR END------------")
	return 0
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if response.StatusCode == 200 {
		fmt.Println(string(body))
		return 1
	}

	return 0
}

var httpClient1000 = createHTTPClient(1000)
func createHTTPClient(requestTimeout int) *http.Client {
	transport := &http.Transport{
		MaxIdleConns: 1000,
		MaxIdleConnsPerHost: 1000,
		IdleConnTimeout:       300 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout: time.Duration(requestTimeout) * time.Millisecond,
	}
	return client
}


func HttpGetRequest(apiUrl string, httpClient *http.Client)  (*http.Response, []byte, error) {
	response, err := httpClient.Get(apiUrl)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, body, nil
}