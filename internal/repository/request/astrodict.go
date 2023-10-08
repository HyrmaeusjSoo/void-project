package request

import (
	"encoding/json"
	"net/http"
	"void-project/internal/model"
	"void-project/pkg"
)

type AD struct {
	url  string
	Type string
}

func NewAD(typ string) *AD {
	return &AD{
		"https://HyrmaeusjSoo.github.io/CHAOS_Project/ding/astrodict_191103/astrodict_191103",
		pkg.IfElse(typ == "", "ce", typ),
	}
}

// 从远程Json获取
func (ad *AD) GetAstroDict() (astrodict *model.AstroDictJson, err error) {
	resp, err := http.Get(ad.url + ad.Type + ".json")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	astrodict = &model.AstroDictJson{}
	err = json.NewDecoder(resp.Body).Decode(astrodict)
	return
}
