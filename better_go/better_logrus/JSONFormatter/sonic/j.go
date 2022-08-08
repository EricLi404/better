package s

import (
	"bytes"
	"encoding/json"
	"github.com/bytedance/sonic"
	encoder2 "github.com/bytedance/sonic/encoder"
)

type T struct {
	A string   `json:"a"`
	B int64    `json:"b"`
	C string   `json:"c"`
	D []string `json:"d"`
	E float64  `json:"e"`
}

var testData = &T{
	A: "hello hello hello",
	B: 2333,
	C: "",
	D: []string{"zxcv", "asdf", "qwer"},
	E: 6666.66,
}

var successString = "{\"a\":\"hello hello hello\",\"b\":2333,\"c\":\"\",\"d\":[\"zxcv\",\"asdf\",\"qwer\"],\"e\":6666.66}\n"

func nativeJsonEncoder(data interface{}) ([]byte, error) {
	b := &bytes.Buffer{}
	encoder := json.NewEncoder(b)
	err := encoder.Encode(data)
	return b.Bytes(), err
}

func nativeJson(data interface{}) ([]byte, error) {
	marshal, err := json.Marshal(data)
	return append(marshal, '\n'), err
}

func sonicJsonEncoder(data interface{}) ([]byte, error) {
	b := &bytes.Buffer{}
	encoder := encoder2.NewStreamEncoder(b)
	err := encoder.Encode(data)
	return b.Bytes(), err
}

func sonicJson(data interface{}) ([]byte, error) {
	marshal, err := sonic.Marshal(data)
	return append(marshal, '\n'), err
}
