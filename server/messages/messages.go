package messages

import (
	"encoding/json"
	"io/ioutil"
)

// Messages stores message strings
var Messages = map[string]map[string]string{}

func init() {
	input, _ := ioutil.ReadFile("messages/messages.json")
	json.Unmarshal(input, &Messages)
}
