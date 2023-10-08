package request

import (
	"encoding/json"
	"net/http"
	"void-project/internal/model"
)

type IPQuery struct {
	url   string
	param string
}

func NewIPQuery() *IPQuery {
	return &IPQuery{"http://ip-api.com/json/", "?lang=zh-CN"}
}

// 从远程Json获取
//
//	也可使用ip2region本地库方式, 地址:github.com/lionsoul2014/ip2region/binding/golang
func (iq *IPQuery) GetIPInfo(ip string) (info *model.IPQuery, err error) {
	resp, err := http.Get(iq.url + ip + iq.param)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	info = &model.IPQuery{}
	err = json.NewDecoder(resp.Body).Decode(info)
	return
}
