package ip2location

import (
	"errors"
	"fmt"
	"encoding/json"

	"github.com/ip2location/ip2location-go/v9"
)

var ip2locationIPLite = "./assets/ip2location/IP2LOCATION-LITE-DB11.BIN"

var ip2locationIP = "./assets/ip2location/IP-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE.BIN"

var (
	ip2locationIPDB   	*ip2location.DB
	ip2locationIPDBLite *ip2location.DB
)

type GeoCountry struct {
	GeoNameId uint
	IsoCode   string
	Name      string
	City      string
	CityCode  int
	Latitude  float64
	Longitude float64
}

func InitIp2Location()  error{
	var err error

	ip2locationIPDB, err = ip2location.OpenDB(ip2locationIP)
	if err != nil {
		return errors.New("fail to open ip2locationIP db: path=" + ip2locationIP + ", err=" + err.Error())
	}

	ip2locationIPDBLite, err = ip2location.OpenDB(ip2locationIPLite)
	if err != nil {
		return errors.New("fail to open ip2locationIP db: path=" + ip2locationIP + ", err=" + err.Error())
	}

	return nil
}

func GetGeoCountryByIpLite(ipStr string) (ip2location.IP2Locationrecord, error) {
	results, err := ip2locationIPDBLite.Get_all(ipStr)

	if err != nil {
		fmt.Print(err)
		return ip2location.IP2Locationrecord{}, err
	}

	return results, nil
}

func GetGeoCountryByIp(ipStr string) (ip2location.IP2Locationrecord, error) {
	results, err := ip2locationIPDB.Get_all(ipStr)

	if err != nil {
		fmt.Print(err)
		return ip2location.IP2Locationrecord{}, err
	}

	//fmt.Println(String(results))

	return results, nil
}

func String(ad interface{}) string {
	b, err := json.Marshal(ad)
	if err != nil {
		return fmt.Sprintf("%v", ad)
	}
	return string(b)
}