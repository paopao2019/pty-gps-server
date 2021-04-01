package config

// 高德地图的key
var (
	GeoKEY = "5df120aa91c72b6efd45ee3fa125299f"
	GeoApiBASE = "https://restapi.amap.com/v3"
)


type GeoCode struct {
	FormattedAddress string        `json:"formatted_address"`
	Country          string        `json:"country"`
	Province         string        `json:"province"`
	CityCode         string        `json:"citycode"`
	City             string        `json:"city"`
	District         string        `json:"district"`
	TownShip         string        `json:"township"`
	Neighborhood     *Neighborhood `json:"neighborhood"`
	Building         *Building     `json:"building"`
	AdCode           string        `json:"adcode"`
	Street           string        `json:"street"`
	Number           string        `json:"number"`
	Location         string        `json:"location"`
	Level            string        `json:"level"`
}

type Building struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Neighborhood struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// 编码
type Geo struct {
	Count    string     `json:"count"`
	GeoCodes []*GeoCode `json:"geocodes"`
}

type ReGeoCode struct {
	Roads            []*Road           `json:"roads"`
	AddressComponent *AddressComponent `json:"addressComponent"`
	FormattedAddress string            `json:"formatted_address"`
	POIs             []*POI            `json:"pois"`
}

type AddressComponent struct {
	City         string        `json:"city"`
	Province     string        `json:"province"`
	AdCode       string        `json:"adcode"`
	District     string        `json:"district"`
	TownCode     string        `json:"towncode"`
	StreetNumber *StreetNumber `json:"streetNumber"`
	Country      string        `json:"country"`
	TownShip     string        `json:"township"`
	//BusinessAreas []*BusinessArea `json:"businessAreas"`
	Neighborhood *Neighborhood `json:"neighborhood"`
	Building     *Building     `json:"building"`
	CityCode     string        `json:"citycode"`
}

type StreetNumber struct {
	Number    string `json:"number"`
	Location  string `json:"location"`
	Direction string `json:"direction"`
	Distance  string `json:"distance"`
	Street    string `json:"street"`
}

type BusinessArea struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	BusinessArea string `json:"businessArea"`
}

type Road struct {
	Id           string `json:"id"`
	Location     string `json:"location"`
	Direction    string `json:"direction"`
	Name         string `json:"name"`
	BusinessArea string `json:"businessArea"`
	Distance     string `json:"distance"`
}

type POI struct {
	Id        string `json:"id"`
	Direction string `json:"direction"`
	//BusinessAreas []*BusinessArea `json:"businessAreas"`
	Address   string `json:"address"`
	PoiWeight string `json:"poiweight"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Distance  string `json:"distance"`
	Tel       string `json:"tel"`
	Type      string `json:"type"`
}

type ReGeo struct {
	ReGeoCode *ReGeoCode
}

// 两点之间的距离
type Distance struct {
	Status        string `json:"status"`
	Info        string `json:"info"`
	InfoCode        string `json:"infocode"`
	Count    string     `json:"count"`
	Results []*Result `json:"results"`
}

type Result struct {
	OriginId        string `json:"origin_id"`
	DestId        string `json:"dest_id"`
	Distance        string `json:"distance"`
	Duration        string `json:"duration"`

}


