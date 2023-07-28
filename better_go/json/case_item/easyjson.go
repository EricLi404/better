package case_item

import "github.com/mailru/easyjson/jwriter"

var w = &jwriter.Writer{}

func Marshal_easyjson(d []*Item) {
	for _, v := range d {
		v.MarshalJSON()
	}
	return
}

// func Marshal_easyjson(d []*Item) (r [][]byte) {
// 	for _, v := range d {
// 		bytes, err := v.MarshalJSON()
// 		if err != nil {
// 			r = append(r, bytes)
// 		}
// 	}
// 	return
// }

// func Marshal_easyjson2(d []*Item) (r [][]byte) {
// 	for _, v := range d {
// 		v.MarshalEasyJSON(w)
// 		r = append(r, w.Buffer.BuildBytes())
// 	}
// 	return
// }

func Unmarshal_easyjson(d [][]byte) {
	item := new(Item)
	for _, v := range d {
		item.UnmarshalJSON(v)
	}
}
