package model

type AstroDictJson struct {
	AstroDict []struct {
		C string `json:"C"`
		E string `json:"E"`
	} `json:"AstroDict"`
}

type Astrodict struct {
	C string ``
	E string ``
}

func (*Astrodict) TableName() string {
	return "Astrodict_CE"
}

type AstrodictEC Astrodict

func (*AstrodictEC) TableName() string {
	return "Astrodict_EC"
}
