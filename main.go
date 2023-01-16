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
	// need to use a different type here else will crash
	aux := &struct {
		Identity map[string]string `json:"Identity"`
		Birthday time.Time         `json:"Birthday"`
	}{}

	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}

	// custome setting of identity here
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
	// anon struct as json object
	fakeIdent := struct {
		Identity map[string]string
		Birthday time.Time
	}{
		Identity: map[string]string{"firstname": "john", "lastname": "smith"},
		Birthday: time.Now(),
	}

	// marhsal to emulate inbound request
	mident, _ := json.Marshal(fakeIdent)

	// we want to combine firstname + lastname onto 'identity' field of 'identity' struct via custom unmarshal
	var uident identity

	json.Unmarshal(mident, &uident)

	fmt.Printf("%+v", uident)
}
