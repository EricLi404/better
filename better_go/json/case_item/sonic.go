package case_item

import (
	"github.com/bytedance/sonic"
)

func Marshal_sonic(d []*Item) {
	for _, v := range d {
		sonic.Marshal(v)
	}

}

func Unmarshal_sonic(d [][]byte) {
	item := new(Item)
	for _, v := range d {
		sonic.Unmarshal(v, item)
	}
}
