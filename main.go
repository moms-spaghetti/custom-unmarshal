package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type identity struct {
	Identity string
	Birthday time.Time
}

func (i *identity) UnmarshalJSON(b []byte) error {
	aux := &struct {
		Identity map[string]string `json:"Identity"`
		Birthday time.Time         `json:"Birthday"`
	}{}

	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}

	index := 0
	for _, v := range aux.Identity {
		i.Identity += v
		index++

		if index != len(aux.Identity) {
			i.Identity += " "
		}
	}

	i.Birthday = aux.Birthday

	return nil
}

func main() {
	fakeIdent := struct {
		Identity map[string]string
		Birthday time.Time
	}{
		Identity: map[string]string{"firstname": "john", "lastname": "smith"},
		Birthday: time.Now(),
	}

	mident, _ := json.Marshal(fakeIdent)

	var uident identity

	json.Unmarshal(mident, &uident)

	fmt.Printf("%+v", uident)
}
