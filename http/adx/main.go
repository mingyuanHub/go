package main

import (
	"compress/gzip"
	"io/ioutil"
	"net"
	"net/http"
	"bytes"
	"io"
	"errors"
	"fmt"
	"time"
)

func main()  {
	fmt.Println("------------START------------")

	for i := 1; i < 100; i ++ {
		res := post()
		if res == 1 {
			break
		}
	}

	fmt.Println("------------END------------")
}

var apiUrl = "http://openbid.pubmatic.com/translator?pubId=160692"
//pubmatic video rewarded
//var bytesData = "{\"id\":\"2a21e3b5-bb2d-43e1-1e98-bcc13230914f\",\"imp\":[{\"id\":\"1\",\"video\":{\"ext\":{\"rewarded\":\"1\"},\"linearity\":1,\"companiontype\":[1,2],\"h\":2131,\"skip\":1,\"skipmin\":1,\"minduration\":3,\"mimes\":[\"video/MP4\",\"video/AVI\"],\"maxduration\":120,\"w\":1080,\"startdelay\":5,\"api\":[1,2,3,4,5,6],\"protocols\":[1,2,3,4,5,6,7,8,9,10]},\"instl\":0,\"tagid\":\"3943799\",\"bidfloor\":0.01,\"bidfloorcur\":\"\",\"clickbrowser\":0,\"secure\":1,\"exp\":0,\"iframebuster\":0,\"metric\":null}],\"app\":{\"id\":\"119102\",\"bundle\":\"pampam.ibf2\",\"domain\":\"tastypill.com\",\"storeurl\":\"https://play.google.com/store/apps/details?id=pampam.ibf2\",\"cat\":[\"IAB1\"]},\"device\":{\"ua\":\"Mozilla/5.0 (Linux; Android 9; vivo 1902 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36\",\"geo\":{\"country\":\"IDN\",\"city\":\"bandung\"},\"ip\":\"203.78.112.246\",\"devicetype\":4,\"make\":\"vivo\",\"model\":\"1902\",\"os\":\"android\",\"osv\":\"9.0.0\",\"connectiontype\":4,\"ifa\":\"07745767-e48a-4b2f-bbc8-2e043760310c\"},\"test\":0,\"at\":1,\"regs\":{}}"

//pubmatic video instl =1
//var bytesData = "{\"id\":\"2a21e3b5-bb2d-43e1-1e98-bcc13230914f\",\"imp\":[{\"id\":\"1\",\"video\":{\"ext\":{\"rewarded\":\"0\"},\"linearity\":1,\"companiontype\":[1,2],\"h\":2131,\"skip\":1,\"skipmin\":1,\"minduration\":3,\"mimes\":[\"video/MP4\",\"video/AVI\"],\"maxduration\":120,\"w\":1080,\"startdelay\":5,\"api\":[1,2,3,4,5,6],\"protocols\":[1,2,3,4,5,6,7,8,9,10]},\"instl\":1,\"tagid\":\"3941854\",\"bidfloor\":0.01,\"bidfloorcur\":\"\",\"clickbrowser\":0,\"secure\":1,\"exp\":0,\"iframebuster\":0,\"metric\":null}],\"app\":{\"id\":\"119102\",\"bundle\":\"pampam.ibf2\",\"domain\":\"tastypill.com\",\"storeurl\":\"https://play.google.com/store/apps/details?id=pampam.ibf2\",\"cat\":[\"IAB1\"]},\"device\":{\"ua\":\"Mozilla/5.0 (Linux; Android 9; vivo 1902 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36\",\"geo\":{\"country\":\"IDN\",\"city\":\"bandung\"},\"ip\":\"203.78.112.246\",\"devicetype\":4,\"make\":\"vivo\",\"model\":\"1902\",\"os\":\"android\",\"osv\":\"9.0.0\",\"connectiontype\":4,\"ifa\":\"07745767-e48a-4b2f-bbc8-2e043760310c\"},\"test\":0,\"at\":1,\"regs\":{}}"

//pubmatic banner instl =1 插屏图片
var bytesData = "{\"id\": \"2a21e3b5-bb2d-43e1-1e98-bcc13230914f\", \"imp\": [{\"id\": \"1\", \"banner\": {\"w\": 1080, \"h\": 2131, \"battr\": [17 ], \"api\": [3, 5, 6 ] }, \"instl\": 1, \"tagid\": \"3945016\", \"bidfloor\": 0, \"bidfloorcur\": \"\", \"clickbrowser\": 0, \"secure\": 1, \"exp\": 0, \"iframebuster\": 0, \"metric\": null } ], \"app\": {\"id\": \"119102\", \"bundle\": \"pampam.ibf2\", \"domain\": \"tastypill.com\", \"storeurl\": \"https://play.google.com/store/apps/details?id=pampam.ibf2\", \"cat\": [\"IAB1\"] }, \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 9; vivo 1902 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36\", \"geo\": {\"country\": \"IDN\", \"city\": \"bandung\"}, \"ip\": \"203.78.112.246\", \"devicetype\": 4, \"make\": \"vivo\", \"model\": \"1902\", \"os\": \"android\", \"osv\": \"9.0.0\", \"connectiontype\": 4, \"ifa\": \"07745767-e48a-4b2f-bbc8-2e043760310c\"}, \"test\": 0, \"at\": 1, \"regs\": {} }"

//pubmatic banner instl =1
//var bytesData = "{\"id\":\"2a21e3b5-bb2d-43e1-1e98-bcc13230914f\",\"imp\":[{\"id\":\"1\",\"banner\":{\"w\":1080,\"h\":2131,\"battr\":[17],\"api\":[3,5,6]},\"instl\":1,\"tagid\":\"3941854\",\"bidfloor\":0,\"bidfloorcur\":\"\",\"clickbrowser\":0,\"secure\":1,\"exp\":0,\"iframebuster\":0,\"metric\":null}],\"app\":{\"id\":\"119102\",\"bundle\":\"pampam.ibf2\",\"domain\":\"tastypill.com\",\"storeurl\":\"https://play.google.com/store/apps/details?id=pampam.ibf2\",\"cat\":[\"IAB1\"]},\"device\":{\"ua\":\"Mozilla/5.0 (Linux; Android 9; vivo 1902 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36\",\"geo\":{\"country\":\"IDN\",\"city\":\"bandung\"},\"ip\":\"203.78.112.246\",\"devicetype\":4,\"make\":\"vivo\",\"model\":\"1902\",\"os\":\"android\",\"osv\":\"9.0.0\",\"connectiontype\":4,\"ifa\":\"07745767-e48a-4b2f-bbc8-2e043760310c\"},\"test\":0,\"at\":1,\"regs\":{}}"

//pubmatic video 贴片
//var bytesData = "{\"id\":\"2a21e3b5-bb2d-43e1-1e98-bcc13230914f\",\"imp\":[{\"id\":\"1\",\"video\":{\"ext\":{\"rewarded\":\"0\"},\"linearity\":1,\"placement\":1,\"companiontype\":[1,2],\"h\":640,\"skip\":1,\"skipmin\":1,\"minduration\":3,\"mimes\":[\"video/MP4\",\"video/AVI\"],\"maxduration\":120,\"w\":320,\"startdelay\":5,\"api\":[1,2,3,4,5,6],\"protocols\":[1,2,3,4,5,6,7,8,9,10]},\"instl\":0,\"tagid\":\"3943799\",\"bidfloor\":0.01,\"bidfloorcur\":\"\",\"clickbrowser\":0,\"secure\":1,\"exp\":0,\"iframebuster\":0,\"metric\":null}],\"app\":{\"id\":\"119102\",\"bundle\":\"pampam.ibf2\",\"domain\":\"tastypill.com\",\"storeurl\":\"https://play.google.com/store/apps/details?id=pampam.ibf2\",\"cat\":[\"IAB1\"]},\"device\":{\"ua\":\"Mozilla/5.0 (Linux; Android 9; vivo 1902 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36\",\"geo\":{\"country\":\"IDN\",\"city\":\"bandung\"},\"ip\":\"203.78.112.246\",\"devicetype\":4,\"make\":\"vivo\",\"model\":\"1902\",\"os\":\"android\",\"osv\":\"9.0.0\",\"connectiontype\":4,\"ifa\":\"07745767-e48a-4b2f-bbc8-2e043760310c\"},\"test\":0,\"at\":1,\"regs\":{}}"



//var apiUrl = "http://api.w.inmobi.com/ortb"
//
//var bytesData = "{\"app\": {\"storeurl\": \"https://play.google.com/store/apps/details?id=com.QuickLoad.MergeCannonDefense\", \"domain\": \"www.inmobi.com\", \"cat\": [\"IAB15\", \"IAB15-10\"], \"bundle\": \"com.QuickLoad.MergeCannonDefense\"}, \"tmax\": 1000, \"regs\": {\"ext\": {}, \"coppa\": \"1\"}, \"id\": \"1af6fcc43c144932bd5bdfc98450b391\", \"source\": {\"pchain\": \"1f6c11bf-56c2-472e-8b31-fe13fe6669fa\"}, \"imp\": [{\"ext\": {\"placementid\": \"1627698383086\"}, \"tagid\": \"1627698383086\", \"bidfloor\": 0.03, \"id\": \"1af6fcc43c144932bd5bdfc98450b391\", \"video\": {\"ext\": {\"rewarded\": \"1\"}, \"battr\": [13, 14 ], \"linearity\": 1, \"companiontype\": [1, 2 ], \"h\": 360, \"skip\": 1, \"skipmin\": 1, \"minduration\": 5, \"mimes\": [\"video/MP4\", \"video/AVI\"], \"maxduration\": 30, \"w\": 640, \"startdelay\": 5, \"api\": [1, 2, 3, 4, 5, 6 ], \"protocols\": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ] }, \"secure\": 0, \"exp\": 0, \"instl\": 1 } ], \"device\": {\"ext\": {\"idv\": \"\", \"gdpr\": \"1\"}, \"os\": \"Android\", \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"hwv\": \"ViVo\", \"ip\": \"23.82.46.51\", \"dnt\": \"\", \"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 Mobile Safari/537.36\", \"osv\": \"7.1.1\", \"model\": \"vivo Y75A\", \"make\": \"vivo\"}, \"user\": {\"gender\": \"F\", \"keywords\": \"intent\", \"id\": \"inmobi-testHackett\", \"Yob\": \"25\"} }"

//var bytesData = "{\"id\": \"d5dd09f5-1a05-2e54-b4b5-387e08209079\", \"imp\": [{\"bidfloor\": 0.01, \"banner\": {\"w\": 320, \"h\": 640, \"id\": \"71ee43cd974c43e9a6ad59b7386279b8\", \"api\": [1, 2, 3, 4, 5, 6 ] }, \"id\": \"71ee43cd974c43e9a6ad59b7386279b8\", \"secure\": 1, \"tagid\": \"1607075241488\", \"exp\": 10800, \"ext\": {\"placementid\": \"1607075241488\"}, \"instl\": 0 } ], \"app\": {\"bundle\": \"com.unstall.meetdeleteapp\", \"storeurl\": \"https://play.google.com/store/apps/details?id=com.mxtech.videoplayer.ad\", \"domain\": \"www.inmobi.com\", \"cat\": [\"IAB15\", \"IAB15-10\"] }, \"device\": {\"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"ip\": \"23.82.46.51\", \"devicetype\": 1, \"make\": \"google\", \"model\": \"Nexus 6\", \"os\": \"Android\", \"osv\": \"7.1.1\", \"hwv\": \"shamu\", \"pxratio\": 3.5, \"connectiontype\": 2, \"ifa\": \"2da994ba-0e4c-4618-a96a-10ca1ad1abe1\"}, \"user\": {\"id\": \"UID-edf3508b-494d-4394-bcc5-9583db5f7b55\", \"ext\": {} }, \"tmax\": 1000, \"regs\": {\"ext\": {} }, \"test\": 0 }"


func post() int {

	var header = make(map[string]string)

	var httpClient1000 = createHTTPClient(2000)

	response, body, err := HttpPostRequest(apiUrl, []byte(bytesData), header, httpClient1000)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Println(response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println(string(body))
		return 1
	}

	return 0
}


func createHTTPClient(requestTimeout int) *http.Client {
	transport := &http.Transport{
		MaxIdleConns: 200,
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

func HttpPostRequest(apiUrl string, bytesData []byte, headers map[string]string, httpClient *http.Client)  (*http.Response, []byte, error){
	var err error

	var isGzip = false
	if len(headers) > 0 {
		for key, item := range headers {
			if key == "Accept-Encoding" && item == "gzip" {
				isGzip = true
				break
			}
		}
	}

	var reader *bytes.Buffer

	if isGzip {
		var zBuf bytes.Buffer
		zw := gzip.NewWriter(&zBuf)
		if _, err = zw.Write(bytesData); err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("gzip error='%s'", err))
		}
		zw.Close()
		reader = &zBuf
	} else {
		reader = bytes.NewBuffer(bytesData)

	}

	request, err := http.NewRequest("POST", apiUrl, reader)

	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("newRequest error='%s'", err))
	}

	if len(headers) > 0 {
		for key, item := range headers {
			request.Header.Set(key, item)
		}
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}

	defer response.Body.Close()

	body := response.Body

	//if response.Header.Get("Content-Encoding") == "gzip" {
	//	body, err = gzip.NewReader(response.Body)
	//	if err != nil {
	//		return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", err))
	//	}
	//}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}
