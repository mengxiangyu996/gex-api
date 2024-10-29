package ip

import (
	"encoding/json"
	"net"
	"ruoyi-go/constant"
	"ruoyi-go/utils"
	"ruoyi-go/utils/curl"
)

var (
	unknown = "未知地址"
)

type IpAddress struct {
	Ip         string `json:"ip"`
	Pro        string `json:"pro"`
	ProCode    string `json:"proCode"`
	City       string `json:"city"`
	CityCode   string `json:"cityCode"`
	Region     string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr       string `json:"addr"`
}

// 获取地址
func GetAddress(ip string) *IpAddress {

	if netIp := net.ParseIP(ip); netIp == nil || netIp.IsLoopback() {
		return &IpAddress{
			Ip:   ip,
			Addr: unknown,
		}
	}

	if utils.CheckRegex(constant.InternalIp, ip) {
		return &IpAddress{
			Ip:   ip,
			Addr: "内网地址",
		}
	}

	request := curl.DefaultClient()

	body, err := request.Send(&curl.RequestParam{
		Url: constant.IPUrl,
		Query: map[string]interface{}{
			"ip":   ip,
			"json": true,
		},
	})

	var ipAddress IpAddress

	if err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = unknown
		return &ipAddress
	}

	if err := json.Unmarshal([]byte(body), &ipAddress); err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = unknown
		return &ipAddress
	}

	return &ipAddress
}
