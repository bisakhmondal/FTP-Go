package credentials

import (
	"os"
	"testing"
)

func TestCredentials_FromJSON(t *testing.T) {
	f,_ := os.Open("a.json")
	//con,_ := ioutil.ReadAll(f)
	//t.Log(string(con))
	var creds CredArr
	err := creds.FromJSON(f)
	//err := json.Unmarshal(con, &creds)
	if err !=nil{
		t.Fatal(err)
	}

	for _,obj := range creds{
		t.Log(obj.Password, " ", obj.Username)
	}
}
