package sipuni_wrapper

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main() {
	args := Arguments{
		"anonymous":       AnonymousTrue,
		"firstTime":       FirstTimeFalse,
		"from":            time.Now().AddDate(0, 0, -1).Format("02.01.2006"),
		"fromNumber":      EmptyString,
		"numbersRinged":   NumbersRingedFalse,
		"outgoingLine":    OutgoingLineTrue,
		"showTreeId":      ShowTreeIdTrue,
		"state":           StateMissed,
		"to":              time.Now().Format("02.01.2006"),
		"toAnswer":        EmptyString,
		"toNumber":        EmptyString,
		"tree":            EmptyString,
		"type":            TypeIncoming,
		"user":            "099172", // also needs for md5
		"dtmfUserAnswer":  DtmfUserAnswerTrue,
		"names":           NamesTrue,
		"numbersInvolved": NumbersInvolvedFalse,
		"key":             "0.mkcp3z3odb",
	}
	ptClient := NewClient("0.mkcp3z3odb", "099172")
	records := ptClient.Post("", args)
	var builder = strings.Builder{}
	encoder := json.NewEncoder(&builder)
	for _, record := range records {
		encoder.Encode(record)
	}
	fmt.Println(builder.String())
}
