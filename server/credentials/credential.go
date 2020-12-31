package credentials

import (
	"encoding/json"
	"io"
)

type Credentials struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredArr []Credentials
func (p * CredArr) FromJSON(r io.Reader)error{
	en := json.NewDecoder(r)
	return en.Decode(p)
}



