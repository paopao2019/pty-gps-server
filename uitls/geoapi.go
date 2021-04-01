package uitls

import (
	"PTY_GPS/config"
	"github.com/asmcos/requests"
	"log"
)

/*
// restapi.amap.com/v3/geocode/geo?key=您的key&address=江浦路1000号&city=上海
获取编码 经纬度
 */

func GetPosition(address string, city string) (err error, location string) {
	var json config.Geo

	p := requests.Params{
		"key": config.GeoKEY,
		"address": address,
		"city":  city,
	}
	URL := config.GeoApiBASE + "/geocode/geo"
	req := requests.Requests()
	resp, err := req.Get(URL, p)
	if err != nil {
		log.Println(err.Error())
		return err, ""
	}
	resp.Json(&json)
	location = json.GeoCodes[0].Location
	return err, location
}



func GetDistance(origins, destination string) (err error, distance string, duration string) {
	var json config.Distance

	p := requests.Params{
		"key": config.GeoKEY,
		"origins": origins,
		"destination":  destination,
		"type": "1", // 1：驾车导航距离
	}
	URL := config.GeoApiBASE + "/distance"
	req := requests.Requests()
	resp, err := req.Get(URL, p)
	if err != nil {
		log.Println(err.Error())
		return err, "",""
	}
	resp.Json(&json)
	distance = json.Results[0].Distance
	duration = json.Results[0].Duration
	return err, distance, duration
}