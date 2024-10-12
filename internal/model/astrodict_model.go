package model

type (
	// 天文学词典Json
	AstroDictJson struct {
		AstroDict []struct {
			C string `json:"C"`
			E string `json:"E"`
		} `json:"AstroDict"`
	}
	// 天文学词典
	Astrodict struct {
		C string `` // 中文
		E string `` // 英文
	}
	// 天文学词典，英-中
	AstrodictEC Astrodict
)

func (*Astrodict) TableName() string {
	return "Astrodict_CE"
}

func (*AstrodictEC) TableName() string {
	return "Astrodict_EC"
}

func (*AstroDictJson) Name() string {
	return "astrodict_ce"
}
