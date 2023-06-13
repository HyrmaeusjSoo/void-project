package request

import (
	"chat/internal/model"
	"chat/pkg"
	"encoding/json"
	"net/http"
)

type AD struct {
	Type string
}

func NewAD(typ string) *AD {
	typ = pkg.IfElse(typ == "", "ce", typ)
	return &AD{typ}
}

func (ad *AD) GetAstroDict() (astrodict *model.AstroDict, err error) {
	resp, err := http.Get("https://hyleasoo.github.io/CHAOS_Project/ding/astrodict_191103/astrodict_191103" + ad.Type + ".json")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	astrodict = &model.AstroDict{}
	err = json.NewDecoder(resp.Body).Decode(astrodict)
	return
}
