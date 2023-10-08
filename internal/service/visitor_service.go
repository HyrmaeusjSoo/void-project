package service

import (
	"time"
	"void-project/internal/model"
	"void-project/internal/repository/request"
	"void-project/pkg/logger/slog"
)

type VisitorService struct {
	iqr *request.IPQuery
}

func NewVisitorService() *VisitorService {
	return &VisitorService{request.NewIPQuery()}
}

// 查询ip
func (v *VisitorService) IPQuery(ip string) (*model.IPQuery, error) {
	return v.iqr.GetIPInfo(ip)
}

// 读取访问日志
func (v *VisitorService) ReadLog(beginDate, endDate time.Time) ([]slog.LogFile, error) {
	return slog.Read(beginDate, endDate)
}

// 访问统计
func (v *VisitorService) Stat(beginDate, endDate time.Time) ([]map[string]any, error) {
	logs, err := slog.Read(beginDate, endDate)
	if err != nil {
		return nil, err
	}

	maps := map[string]string{}
	counts := map[string]int{}

	for _, f := range logs {
		for _, l := range f.Logs {
			region, ok := maps[l.IP]
			if !ok {
				ip, err := v.iqr.GetIPInfo(l.IP)
				if err != nil {
					return nil, err
				}
				region = ip.Country + ip.RegionName + ip.City
				if region == "" {
					region = "其他"
				}
				maps[l.IP] = region
			}
			counts[region]++
		}
	}

	res := []map[string]any{}
	for k, v := range counts {
		res = append(res, map[string]any{"region": k, "count": v})
	}

	return res, nil
}
