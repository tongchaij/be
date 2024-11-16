package main

import (
	"bytes"
	"encoding/json"
	"testing"

	pb "be3/beef"
)

type BeefData struct {
	Data map[string]interface{} `json:"beef"`
}

func TestGetJsonShort(t *testing.T) {
	conn := GetConnection()
	defer conn.Close()
	c := pb.NewBeefClient(conn)

	fileName := "short.txt"
	got := GetJson(c, fileName)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(got), "", "\t")
	if error != nil {
		t.Errorf("JSON parse error: %v", error)
		return
	}
	t.Logf("Json:\n%s\n", prettyJSON.String())

	var gotObj BeefData
	err := json.Unmarshal([]byte(got), &gotObj)
	if err != nil {
		panic(err)
	}
	t.Logf("gotObj:%v", gotObj)

	want := BeefData{
		Data: map[string]interface{}{
			"tip":   4.0,
			"jowl":  1.0,
			"nulla": 3.0,
		},
	}

	areEqual := true

	for k, _ := range want.Data {
		areEqual = areEqual && gotObj.Data[k] == want.Data[k]
	}

	if !areEqual {
		t.Errorf("Compare %v : want %v", gotObj, want)
	}
}

func TestGetJsonBeef(t *testing.T) {
	conn := GetConnection()
	defer conn.Close()
	c := pb.NewBeefClient(conn)

	fileName := "beef.txt"
	got := GetJson(c, fileName)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(got), "", "\t")
	if error != nil {
		t.Errorf("JSON parse error: %v", error)
		return
	}
	t.Logf("Json:\n%s\n", prettyJSON.String())

	var gotObj BeefData
	err := json.Unmarshal([]byte(got), &gotObj)
	if err != nil {
		panic(err)
	}
	t.Logf("gotObj:%v", gotObj)

	want := BeefData{
		Data: map[string]interface{}{
			"pastrami":   53.0,
			"pork":       141.0,
			"beef":       144.0,
			"chicken":    41.0,
			"capicola":   39.0,
			"prosciutto": 35.0,
			"chop":       33.0,
			"drumstick":  41.0,
		},
	}

	areEqual := true

	for k, _ := range want.Data {
		areEqual = areEqual && gotObj.Data[k] == want.Data[k]
	}

	if !areEqual {
		t.Errorf("Compare %v : want %v", gotObj, want)

		for k, _ := range want.Data {
			t.Errorf("%s : %v,%v", k, gotObj.Data[k], want.Data[k])
		}

	}
}
