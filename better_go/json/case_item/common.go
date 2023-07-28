package case_item

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Item struct {
	Flag        int8     `json:"-"`
	Time        int      `json:"impTime"`
	CTime       int      `json:"cTime"`
	IndexInPool uint32   `json:"-"`
	IidSign     uint64   `json:"-"`
	Iid         string   `json:"iid"`
	Fps         []string `json:"fps"`
	FpTitleHash string   `json:"fp_title_hash,omitempty"`
}

func getDataMarshal() []*Item {
	var r []*Item
	f, err := os.Open("test_data.txt")
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			if err == io.EOF && len(line) == 0 {
				break
			}
		}
		i := new(Item)
		err = json.Unmarshal(line, i)
		if err != nil {
			panic(err)
		}
		r = append(r, i)
	}
	if len(r) != 100 {
		fmt.Println("len not enough")
		panic(len(r))
	}
	return r
}

func getDataUnmarshal() [][]byte {
	var r [][]byte
	f, err := os.Open("test_data.txt")
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			if err == io.EOF && len(line) == 0 {
				break
			}
		}
		r = append(r, line)
	}
	if len(r) != 100 {
		fmt.Println("len not enough")
		panic(len(r))
	}
	return r
}
