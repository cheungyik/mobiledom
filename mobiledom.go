package mobiledom

import (
	"encoding/json"
	"regexp"
	"strings"
)

var mobiles = make(map[string]*Mobiledom)

func init() {
	var arr []*Mobiledom
	if err := json.Unmarshal(DefinitionJson, &arr); err != nil {
		panic(err)
	}
	for _, mobiledom := range arr {
		mobiles[mobiledom.SalesModel] = mobiledom
	}
}

func Query(model string) *Mobiledom {
	mobiledom := mobiles[model]
	if mobiledom != nil {
		return mobiledom
	}
	return &Mobiledom{
		Brand: func() string {
			switch regexp.MustCompile(`[0-9]`).ReplaceAllString(strings.ToLower(strings.SplitN(model, " ", 2)[0]), "") {
			case "vivo":
				return "vivo 手机"
			case "mi", "redmi":
				return "小米手机"
			case "meizu":
				return "魅族手机"
			case "oppo":
				return "OPPO 手机"
			default:
				return "UNKNOWN"
			}
		}(),
		ProductName:  model,
		SalesVersion: model,
	}
}
