package sipuni_api_wrapper

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strings"
)

var order = []string{
	"anonymous", "dtmfUserAnswer", "firstTime", "from", "fromNumber", "names",
	"numbersInvolved", "numbersRinged", "outgoingLine", "showTreeId", "state",
	"to", "toAnswer", "toNumber", "tree", "type", "user",
}

// Arguments are used for passing URL parameters to the client for making API calls.
type Arguments map[string]string

// ToURLValuesAndHashMd5 returns the argument's URL value representation with hash.
func (args Arguments) ToURLValuesAndHashMd5() url.Values {
	v := url.Values{}
	key := args["key"]
	delete(args, "key")
	for key, value := range args {
		v.Set(key, value)
	}

	// Объединяем аргументы в нужном порядке для хэша
	union := joinForHash(args, key)

	hashBytes := md5.Sum([]byte(union))
	hashString := hex.EncodeToString(hashBytes[:])

	v.Add("hash", hashString)
	return v
}

func joinForHash(args Arguments, key string) string {

	var sb strings.Builder
	for _, field := range order {
		sb.WriteString(args[field])
		sb.WriteString("+")
	}
	sb.WriteString(key)

	return sb.String()
}
