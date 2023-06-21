package request

import (
	"encoding/json"
	"net/http"
	"void-project/internal/model"
	"void-project/pkg"
)

type AD struct {
	Type string
}

func NewAD(typ string) *AD {
	return &AD{pkg.IfElse(typ == "", "ce", typ)}
}

func (ad *AD) GetAstroDict() (astrodict *model.AstroDict, err error) {
	resp, err := http.Get("https://gitee.com/HyleaSoo/void-project/raw/master/asset/json/astrodict_191103" + ad.Type + ".json")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	astrodict = &model.AstroDict{}
	err = json.NewDecoder(resp.Body).Decode(astrodict)
	return
}
