package main

import (
	"fmt"
	"time"
	"errors"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"sync"
)

type AdRequest struct {
	TP                  *TP  `json:"tp"`
	AdSourcePlacements 	[]*AdSourcePlacement `json:"adsourceplacements"`
	App 				*App `json:"app"`
	Device 				*Device `json:"device"`
	User                *User `json:"user"`
}

type TP struct {
	AppId  string `json:"appid"`
	UnitId string `json:"unitid"`
}

type AdSourcePlacement struct {
	Id 			int `json:"id"`
	Name 		string `json:"name"`
	NetworkId 	int `json:"networkid"`
	NetworkName string `json:"networkname"`
	NetworkVer 	string `json:"networkver"`
	Buyeruid    string `json:"buyeruid"`
}

type App struct {
	Ver string `json:"ver"`
	Ext *Ext `json:"ext"`
}

type Ext struct {
	Orientation int `json:"orientation"`
}

type Device struct {
	DeviceType 		int `json:"devicetype"`
	Make			string `json:"make"`
	Model       	string `json:"model"`
	Os    			int		`json:"os"`
	Osv   			string `json:"osv"`
	H     			int		`json:"h"`
	W 	  			int		`json:"w"`
	Language 		string  `json:"language"`
	Mcc      		string  `json:"mcc"`
	Mnc      		string `json:"mnc"`
	ConnectionType 	int 	`json:"connectiontype"`
	Idfa           	string	`json:"idfa"`
	Idfv           	string	`json:"idfv"`
	Gaid 			string	`json:"gaid"`
}

type User struct {
	Id string
}

var adRequest *AdRequest

func init() {
	adRequest = new(AdRequest)
	adRequest.TP = new(TP)
	adRequest.AdSourcePlacements = make([]*AdSourcePlacement, 0, 0)
	adRequest.App = new(App)
	adRequest.App.Ext = new(Ext)
	adRequest.Device = new(Device)
	adRequest.User = new(User)

	adRequest.TP.AppId = "22222"
	adRequest.TP.UnitId = "22222"
	adRequest.App.Ver = "3.3"
	adRequest.App.Ext.Orientation = 1
	adRequest.Device.DeviceType = 1
	adRequest.Device.W = 414
	adRequest.Device.H = 736
	adRequest.Device.Idfv = "F1FA2D7D-03A8-4EAC-ACBF-28A39A4A4FBF"
	adRequest.Device.Mcc = "460"
	adRequest.Device.Osv = "14.0"
	adRequest.Device.ConnectionType = 2
	adRequest.Device.Os = 2
	adRequest.Device.Mnc = "01"
	adRequest.Device.Language = "zh-Hans-CN"
	adRequest.Device.Make = "Apple"
	adRequest.Device.Idfv = "B6281A9E-3BE0-4884-8D22-7D980A0F4132"
	adRequest.Device.Model = "iPhone"

	var adsourceplacement = new(AdSourcePlacement)
	adsourceplacement.NetworkId = 18
	adsourceplacement.Id = 20672
	adsourceplacement.NetworkVer = "6.6.6"
	adsourceplacement.NetworkName = "Mintegral"
	adsourceplacement.Buyeruid = "xUc7NnzjfTNwiafSGoTTxVcaWVcaN3DwiUtSiAQSfjjTx3z76jRFiUleNnQcWnfoxnvwfal9foT9xaRBWnJjGnlMNnS7fajAigeIHUzUDnftHrNwfa32fZT2fFfQWkjeiniwHnx0DnhFfAftGkjb6deIL5SEYFPQGoMB6aQInVQ6f09FWUHIinRTi09MiavMiaS9iURMGo9MiavMiaSInkK1LkesDZI2WUvlp7QNL7K/HnslN2S5R7QNL7K/HZSyVBvefcIMR7euLFVlnkcURjKnRcluRjcMh7eQ5F50ZFQTWADMfZ9eWUj2RotWZcxfnoMlY7Q8HZSdHkf8YB3lnkK0LkeQWAj2xnjTGdeIk22IinDeiAlefahbfnV2igMBWUcIin39GZ9BfUVFiU5IGal/iUj9fA5IkAjFfAhbiUR9++MPfbMe6aDAGnlAWURMiAjBf+MAGaj9fZ92fANBin3"

	var adsourceplacement2 = new(AdSourcePlacement)
	adsourceplacement2.NetworkId = 33
	adsourceplacement2.Id = 20660
	adsourceplacement2.NetworkVer = "6.6.6"
	adsourceplacement2.NetworkName = "Mytarget"
	adsourceplacement2.Buyeruid = "xUc7NnzjfTNwiafSGoTTxVcaWVcaN3DwiUtSiAQSfjjTx3z76jRFiUleNnQcWnfoxnvwfal9foT9xaRBWnJjGnlMNnS7fajAigeIHUzUDnftHrNwfa32fZT2fFfQWkjeiniwHnx0DnhFfAftGkjb6deIL5SEYFPQGoMB6aQInVQ6f09FWUHIinRTi09MiavMiaS9iURMGo9MiavMiaSInkK1LkesDZI2WUvlp7QNL7K/HnslN2S5R7QNL7K/HZSyVBvefcIMR7euLFVlnkcURjKnRcluRjcMh7eQ5F50ZFQTWADMfZ9eWUj2RotWZcxfnoMlY7Q8HZSdHkf8YB3lnkK0LkeQWAj2xnjTGdeIk22IinDeiAlefahbfnV2igMBWUcIin39GZ9BfUVFiU5IGal/iUj9fA5IkAjFfAhbiUR9++MPfbMe6aDAGnlAWURMiAjBf+MAGaj9fZ92fANBin3"

	adRequest.AdSourcePlacements = append(adRequest.AdSourcePlacements, adsourceplacement)
	adRequest.AdSourcePlacements = append(adRequest.AdSourcePlacements, adsourceplacement2)

	adRequest.User.Id = "2222"

}

func main() {
	num := 2
	for i :=0 ; i < num; i ++ {
		run()
		time.Sleep(time.Millisecond * time.Duration(500))
	}
}

func run()  {
	start := time.Now().UnixNano() / 1e6
	fmt.Println("start time:", start)

	num := 10
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i :=0 ; i < num; i ++ {
		go func() {
			bidResponse ,err:= HttpBidRequest(adRequest, 3000)
			if err != nil {
				fmt.Println("err:", err)
			} else {
				fmt.Println("bidResponse:", bidResponse)
			}

			wg.Done()
		}()
	}
	wg.Wait()

	end := time.Now().UnixNano() / 1e6
	fmt.Println("end time:", end)

	fmt.Println("cost time:", end - start)
}

func HttpBidRequest(bidRequest *AdRequest, bidTimeOut int) (string, error) {

	var apiUrl = "http://test-bidder.tradplus.com/api/v1/headbidding"

	var err error

	bytesData, err := json.Marshal(bidRequest)
	if err != nil {
		return "nil", errors.New(fmt.Sprintf("fail to httpBidRequest, json Marshal. error='%v'", err))
	}

	var jsonStr = []byte(bytesData)
	reader := bytes.NewBuffer(jsonStr)

	request, err := http.NewRequest("POST", apiUrl, reader)

	if err != nil {
		return "nil", errors.New(fmt.Sprintf("fail to httpBidRequest, newRequest. error='%v'", err))
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	defer request.Body.Close()

	client := &http.Client{
		Timeout: time.Millisecond * time.Duration(bidTimeOut),
	}
	response, err := client.Do(request)
	if err != nil {
		return "nil", errors.New(fmt.Sprintf("fail to httpBidRequest, clientDo. error='%v'", err))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("fail to httpBidRequest, ioutilReadAll. error='%v'", err))
	}

	return string(body), nil

	return string(bidTimeOut), nil
}