package querystring

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/martian"
	"github.com/google/martian/parse"
)

func init() {
	parse.Register("querystring.QyWechatApiModifier", qyWechatApiModifierFromJSON)
}

// MarvelModifier contains the private and public Marvel API key
type QyWechatApiModifier struct {
	authServerUrl,realm,clientid,secret string
}

// MarvelModifierJSON to Unmarshal the JSON configuration
type QyWechatApiModifierJSON struct {
	AuthServerUrl  string               `json:"auth-server-url"`
	Realm string               `json:"realm"`
	ClientId string `json:"client-id"`
	Secret string `json:"secret"`
	Scope   []parse.ModifierType `json:"scope"`
}

// ModifyRequest modifies the query string of the request with the given key and value.
func (m *QyWechatApiModifier) ModifyRequest(req *http.Request) error {
	query := req.URL.Query()
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	//hash := GetMD5Hash(ts + m.private + m.public)
	//query.Set("apikey", m.public)
	query.Set("corpid", "wldfc8d2cef1")
	//query.Set("hash", hash)
	//TODO 
	req.URL.RawQuery = query.Encode()

	return nil
}

// GetMD5Hash returns the md5 hash from a string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// MarvelNewModifier returns a request modifier that will set the query string
// at key with the given value. If the query string key already exists all
// values will be overwritten.
func QyWechatApiNewModifier(authServerUrl,realm,clientid,secret string) martian.RequestModifier {
	return &QyWechatApiModifier{
		authServerUrl:  authServerUrl,
		realm: realm,
		clientid: clientid,
		secret: secret,
	}
}

// marvelModifierFromJSON takes a JSON message as a byte slice and returns
// a querystring.modifier and an error.
//
// Example JSON:
// {
//  "public": "apikey",
//  "private": "apikey",
//  "scope": ["request", "response"]
// }
func qyWechatApiModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &QyWechatApiModifierJSON{}

	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return parse.NewResult(QyWechatApiNewModifier(msg.AuthServerUrl, msg.Realm, msg.ClientId,msg.Secret), msg.Scope)
}