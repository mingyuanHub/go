package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
	"errors"
	"fmt"
)

var wg sync.WaitGroup

func BenchmarkBidding(b *testing.B) {

	wg.Add(250)
	for i := 0; i < 250; i++ {
		go bid()
	}
	wg.Wait()

}

//var bidApiUrl = "http://127.0.0.1:8088/api/v1/headbidding"
var bidApiUrl = "https://bidder.tradplusad.com/api/v1/headbidding"
var bytesData = "{\"app\": {\"storeurl\": \"https://play.google.com/store/apps/details?id=com.mxtech.videoplayer.ad\", \"domain\": \"www.inmobi.com\", \"cat\": [\"IAB15\", \"IAB15-10\"], \"bundle\": \"com.mxtech.videoplayer.ad\"}, \"tmax\": 500, \"regs\": {\"ext\": {\"gdpr\": 1 }, \"coppa\": 1, \"ccpa\": 1 }, \"id\": \"71ee43cd974c43e9a6ad59b7386279b8\", \"test\": 1, \"source\": {\"pchain\": \"af67e927-ec85-4abd-a093-cbe8a2c458bb\"}, \"imp\": [{\"ext\": {\"placementid\": \"1604670939163\"}, \"tagid\": \"1604670939163\", \"bidfloor\": 0.03, \"banner\": {\"battr\": [13, 14 ], \"pos\": 2, \"topframe\": 0, \"w\": 300, \"h\": 250, \"format\": [{\"w\": 320, \"h\": 50 } ], \"id\": \"71ee43cd974c43e9a6ad59b7386279b8\", \"api\": [1, 2, 3, 4, 5, 6 ], \"mimes\": [\"img/JPG\", \"img/PNG\"] }, \"id\": \"71ee43cd974c43e9a6ad59b7386279b8\", \"secure\": 0, \"exp\": 0 } ], \"device\": {\"ext\": {\"idfv\": \"12313\", \"gdpr\": \"1\"}, \"os\": 2, \"idfv\": \"dsfasf\", \"ifa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"idfa\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"gaid\": \"9d8fe0a9-c0dd-4482-b16b-5709b00c608d\", \"hwv\": \"ViVo\", \"ip\": \"59.144.134.146\", \"dnt\": 1, \"ua\": \"Mozilla/5.0 (Linux; Android 7.1.1; vivo Y75A Build/N6F26Q; wv)AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/55.0.2883.91 MobileSafari/537.36\", \"osv\": \"7.1.1\", \"model\": \"vivo Y75A\", \"connectiontype\": 1, \"make\": \"vivo\", \"lmt\": 1 }, \"user\": {\"gender\": \"F\", \"keywords\": \"intent\", \"id\": \"inmobi-testHackett\", \"yob\": 2500 }, \"adsourceplacements\": [{\"networkid\": 18, \"id\": 20350, \"networkver\": \"6.55.6\", \"networksdkname\": \"Mintegral\", \"buyeruid\": \"123222222222222222222222222222222\"}, {\"networkid\": 33, \"id\": 20831, \"networkver\": \"6.6.6\", \"networkname\": \"Mytarget\", \"buyeruid\": \"xUc7NnzjfTNwiafSGoTTxVcaWVcaN3DwiUtSiAQSfjjTx3z76jRFiUleNnQcWnfoxnvwfal9foT9xaRBWnJjGnlMNnS7fajAigeIHUzUDnftHrNwfa32fZT2fFfQWkjeiniwHnx0DnhFfAftGkjb6deIL5SEYFPQGoMB6aQInVQ6f09FWUHIinRTi09MiavMiaS9iURMGo9MiavMiaSInkK1LkesDZI2WUvlp7QNL7K/HnslN2S5R7QNL7K/HZSyVBvefcIMR7euLFVlnkcURjKnRcluRjcMh7eQ5F50ZFQTWADMfZ9eWUj2RotWZcxfnoMlY7Q8HZSdHkf8YB3lnkK0LkeQWAj2xnjTGdeIk22IinDeiAlefahbfnV2igMBWUcIin39GZ9BfUVFiU5IGal/iUj9fA5IkAjFfAhbiUR9++MPfbMe6aDAGnlAWURMiAjBf+MAGaj9fZ92fANBin3=\"} ], \"biddingwaterfall\": [] }"

func bid() {
	defer wg.Done()
	var header = make(map[string]string)

	response, body, err := HttpPostRequest(bidApiUrl, []byte(bytesData), header, httpClient10002)
	//return

	fmt.Println(string(body))

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println(string(body))
		return
	}

	return
}

var httpClient10002 = createHTTPClient2(3000)
func createHTTPClient2(requestTimeout int) *http.Client {
	transport := &http.Transport{
		MaxIdleConns: 250,
		MaxIdleConnsPerHost: 250,
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

	if response.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", err))
		}
	}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}






type AdRequest struct {
	Id                 string
	Imp    []*Imp
	App                *App
	Device             *Device
	User               *User
	Tmax               int64
	Test               int
	Cur                string
	Source             *Source
	Regs               *Regs

	TP                 *TP
	AdSourcePlacements []*AdSourcePlacement
	BiddingWaterfall   []*Segment
}

type Imp struct {
	Id                string
	Banner            *Banner
	NativeAd          *Native   // sdk native 为关键字，传nativeAd
	DisplayManager    string
	DisplayManagerVer string
	Instl             int
	TagId             string
	BidFloor          float64
	BidFloorCur       string
	ClickBrowser      int
	Secure            int
	Exp               int
	Ext               *DemoExt
}

type Banner struct {
	Format   []*Format
	W        int
	H        int
	Mimes    []string
	TopFrame int
	Id       string
	Exp      *DemoExt
}

type Format struct {
	W      int
	H      int
	WRatio int
	HRatio int
	WMin   int
	Ext    *DemoExt
}

type Native struct {
	Request string
	Ver     string
	Api     []int
	BAttr   []int
	Ext    *DemoExt
}

type TP struct {
	AppId  string
	UnitId string
}

type AdSourcePlacement struct {
	Id            int
	Name          string
	NetworkId     int
	NetworkName   string
	NetworkSdkVer string
	NetworkVer    string
	Buyeruid      string
}

type Segment struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
	Price string `json:"price"`
}

type App struct {
	Id            string
	Name          string
	Bundle        string
	Domain        string
	StoreUrl      string
	SectionCat    []string
	PageCat       []string
	Ver           string
	PrivacyPolicy int
	Keywords      string
	Ext           *AppExt
}

type AppExt struct {
	Orientation int
}

type Device struct {
	Geo             *Geo
	Lmt             int
	DeviceType 		int
	Make			string
	Model       	string
	Os    			int
	Osv   			string
	Hwv   			string
	H     			int
	W 	  			int
	Ppi             int
	PxRatio         float64
	Js              int
	GeoFetch        int
	FlashVer        string
	Language 		string
	Carrier         string
	ConnectionType 	int
	Ifa         	string
	Mcc      		string
	Mnc      		string
	Idfa           	string
	Idfv           	string
	Gaid 			string
}

//todo: zip
type Geo struct {
	Lat     float64
	Lon     float64
	Type    int
	Country string
	Region  string
}

type User struct {
	Id       string
	Yob      int
	Gender   string
	Keywords string
}

type Source struct {
	Fd     int
	PChain string
}

type Regs struct {
	Coppa int
	Ccpa int
	Ext   *RegsExt
}

type RegsExt struct {
	Gdpr int
}

type DemoExt struct {

}


