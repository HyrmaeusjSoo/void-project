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

func (ad *AD) GetAstroDict() (astrodict *model.AstroDictJson, err error) {
	resp, err := http.Get("https://HyrmaeusjSoo.github.io/CHAOS_Project/ding/astrodict_191103/astrodict_191103" + ad.Type + ".json")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	astrodict = &model.AstroDictJson{}
	err = json.NewDecoder(resp.Body).Decode(astrodict)
	return
}
