package maxMind

import (
	"encoding/csv"
	"errors"
	"net"
	"os"
	"strconv"
)

var geolite2_country_mmdb = "./assets/GeoLite2-Country.mmdb"
var geolite2_city_mmdb = "./assets/GeoLite2-City.mmdb"
var geolite2_country_csv = "./assets/GeoLite2-Country-Locations-zh-CN.csv"

var geoip_city_US_mmdb = "./assets/GeoIP-City-Redacted-US.mmdb"

var (
	countryDatabase *Reader
	cityDatabase    *Reader
	cityUSDatabase    *Reader
	geoCountryDict  = make(map[string]GeoCountry) // isoCode => GeoCountry
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

func InitGeoIp()  error{
	var err error
	var path string

	countryDatabase, err = Open(geolite2_country_mmdb)
	if err != nil {
		return errors.New("fail to open geoip country db: path=" + path + ", err=" + err.Error())
	}

	cityDatabase, err = Open(geolite2_city_mmdb)
	if err != nil {
		return errors.New("fail to open geoip city db: path=" + path + ", err=" + err.Error())
	}

	cityUSDatabase, err = Open(geoip_city_US_mmdb)
	if err != nil {
		return errors.New("fail to open geoip city db: path=" + path + ", err=" + err.Error())
	}

	var csvFile *os.File
	csvFile, err = os.Open(geolite2_country_csv)
	if err != nil {
		return errors.New("fail to open geoip country csv: path=" + path + ", err=" + err.Error())
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	var csvData [][]string
	csvData, err = csvReader.ReadAll()

	for idx, item := range csvData {
		if idx == 0 {
			continue
		}
		if len(item) < 7 {
			return errors.New("fail to load geoip country csv: path=" + path + ", lineno=" + string(idx+1) + ", err='no enough fields'")
		}
		var geoNameId64 uint64
		geoNameId64, err = strconv.ParseUint(item[0], 10, 32)
		if err != nil {
			return errors.New("fail to load geoip country csv: path=" + path + ", lineno=" + string(idx+1) + ", err='parse geoNameId failed'")
		}

		geoCountry := GeoCountry{
			GeoNameId:         uint(geoNameId64),
			IsoCode:           item[4],
			Name:              item[5],
			Latitude:          0.0,
			Longitude:         0.0,
		}

		if geoCountry.IsoCode == "" || geoCountry.Name == "" {
			continue
		}
		geoCountryDict[geoCountry.IsoCode] = geoCountry
	}

	return nil
}

func GetGeoCountryByIp(ipStr string) (*GeoCountry, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, errors.New("fail to getGeoCountryByIp: invalid ip string, ip=" + ipStr)
	}
	geoCountry := new(GeoCountry)

	if record, err := countryDatabase.Country(ip); err != nil {
		return nil, errors.New("fail to getGeoCountryByIp: cannot find geoipUtil record, ip=" + ipStr)
	} else {
		geoCountry.GeoNameId = record.Continent.GeoNameID
		geoCountry.IsoCode = record.Country.IsoCode
		geoCountry.Name = record.Country.Names["zh-CN"]
	}

	if record, err := cityDatabase.City(ip); err != nil {
		geoCountry.Latitude = 0.0
		geoCountry.Longitude = 0.0
	} else {
		//fmt.Println("GeoIP:", ipStr, common.String(record))
		geoCountry.Latitude = record.Location.Latitude
		geoCountry.Longitude = record.Location.Longitude
		geoCountry.City = record.City.Names["en"]
		geoCountry.CityCode = int(record.City.GeoNameID)
	}

	return geoCountry, nil
}

func GetGeoCountryByIp2(ipStr string) (*GeoCountry, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, errors.New("fail to getGeoCountryByIp: invalid ip string, ip=" + ipStr)
	}
	geoCountry := new(GeoCountry)

	if record, err := countryDatabase.Country(ip); err != nil {
		return nil, errors.New("fail to getGeoCountryByIp: cannot find geoipUtil record, ip=" + ipStr)
	} else {
		geoCountry.GeoNameId = record.Continent.GeoNameID
		geoCountry.IsoCode = record.Country.IsoCode
		geoCountry.Name = record.Country.Names["zh-CN"]
	}

	if record, err := cityUSDatabase.City(ip); err != nil {
		geoCountry.Latitude = 0.0
		geoCountry.Longitude = 0.0
	} else {
		//fmt.Println("GeoIP:", ipStr, common.String(record))
		geoCountry.Latitude = record.Location.Latitude
		geoCountry.Longitude = record.Location.Longitude
		geoCountry.City = record.City.Names["en"]
		geoCountry.CityCode = int(record.City.GeoNameID)
	}

	return geoCountry, nil
}

//获取City对应code
func GetCityGeoIp2Code(ip string) int {
	var geoIp2Code int
	if len(ip) > 0 {
		geoCountry , err := GetGeoCountryByIp(ip)
		if err == nil {
			geoIp2Code = geoCountry.CityCode
		}
	}
	return geoIp2Code
}