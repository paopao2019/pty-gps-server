package uitls

import (
	"PTY_GPS/config"
	"github.com/asmcos/requests"
	"github.com/bitly/go-simplejson"
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

// 获取乘坐交通的时间
func GetTranTime(origin, destination string) (err error,  tranTime string) {
	p := requests.Params{
		"key": config.GeoKEY,
		"origin": origin,
		"city": "上海市",
		"destination":  destination,
	}
	URL := config.GeoApiBASE + "/direction/transit/integrated"
	req := requests.Requests()
	resp, err := req.Get(URL, p)
	if err != nil {
		log.Println(err.Error())
		return err, ""
	}
	tranStr := resp.Text()
	jsonObject, err := simplejson.NewJson([]byte(tranStr))
	tranTime = jsonObject.Get("route").Get("transits").GetIndex(0).Get("duration").MustString()
	return err, tranTime
}

