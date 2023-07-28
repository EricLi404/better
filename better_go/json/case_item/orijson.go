package case_item

import "encoding/json"

func Marshal_ori(d []*Item) {
	for _, v := range d {
		json.Marshal(v)
	}

}

func Unmarshal_ori(d [][]byte) {
	item := new(Item)
	for _, v := range d {
		json.Unmarshal(v, item)
	}
}
