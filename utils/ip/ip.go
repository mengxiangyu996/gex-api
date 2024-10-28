package ip

import (
	"encoding/json"
	httputils "ruoyi-go/utils/http-utils"
)

var IpUrl = "http://whois.pconline.com.cn/ipJson.jsp"

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

	request := httputils.DefaultClient()

	body, err := request.Send(&httputils.RequestParam{
		Url: IpUrl,
		Query: map[string]interface{}{
			"ip":   ip,
			"json": true,
		},
	})

	var ipAddress IpAddress

	if err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	if err := json.Unmarshal([]byte(body), &ipAddress); err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return &ipAddress
	}

	return &ipAddress
}
