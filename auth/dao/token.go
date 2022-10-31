package dao

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"time"

	"github.com/94peter/sterna/api/mid"
	"github.com/94peter/sterna/model/cache"
)

type SerializationToken interface {
	cache.SimpleCacheData
	Decode([]byte) error
	GetParseResult() mid.TokenParserResult
}

func NewSerializationToken(key string, result mid.TokenParserResult) SerializationToken {
	st := &serializationToken{
		key: key,
		data: map[string]interface{}{
			"h": result.Host(),
			"p": result.Perms(),
			"a": result.Account(),
			"n": result.Name(),
			"s": result.Sub(),
			"t": result.Target(),
		},
	}
	return st
}

func NewSerializationTokenByKey(key string) SerializationToken {
	return &serializationToken{
		key:  key,
		data: map[string]any{},
	}
}

type serializationToken struct {
	key  string
	data map[string]interface{}
}

func (t *serializationToken) GetKey() string {
	return "auth:token:" + t.key
}

func (t *serializationToken) GetData() ([]byte, error) {
	if len(t.data) == 0 {
		return nil, errors.New("not data")
	}
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(t.data)
	if err != nil {
		return nil, errors.New(`failed gob Encode: ` + err.Error())
	}
	return []byte(base64.StdEncoding.EncodeToString(b.Bytes())), nil
}

func (t *serializationToken) Expired() time.Duration {
	return time.Minute * 30
}

func (t *serializationToken) Decode(data []byte) error {
	by, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return errors.New(`failed base64 Decode: ` + err.Error())
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&t.data)
	if err != nil {
		return errors.New(`failed gob Decode: ` + err.Error())
	}
	return nil
}

func (t *serializationToken) GetParseResult() mid.TokenParserResult {
	if t.data == nil {
		return nil
	}
	return &tokenResult{
		host:    t.data["h"].(string),
		perm:    t.data["p"].([]string),
		account: t.data["a"].(string),
		sub:     t.data["s"].(string),
		name:    t.data["n"].(string),
		target:  t.data["t"].(string),
	}
}

func NewTokenResult(host, account, sub, name, target string, perms []string) mid.TokenParserResult {
	return &tokenResult{
		host:    host,
		target:  target,
		account: account,
		sub:     sub,
		name:    name,
		perm:    perms,
	}
}

type tokenResult struct {
	host    string
	target  string
	perm    []string
	account string
	sub     string
	name    string
}

func (t *tokenResult) Host() string {
	return t.host
}

func (t *tokenResult) Perms() []string {
	return t.perm
}

func (t *tokenResult) Account() string {
	return t.account
}

func (t *tokenResult) Name() string {
	return t.name
}

func (t *tokenResult) Sub() string {
	return t.sub
}

func (t *tokenResult) Target() string {
	return t.target
}
