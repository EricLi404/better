package case_item

import "github.com/json-iterator/go"

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal_jsoniterator(d []*Item) {
	for _, v := range d {
		json2.Marshal(v)
	}

}

func Unmarshal_jsoniterator(d [][]byte) {
	item := new(Item)
	for _, v := range d {
		json2.Unmarshal(v, item)
	}
}
