package mobiledom

import (
	"encoding/json"
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
	return mobiles[model]
}
