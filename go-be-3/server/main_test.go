package main

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type BeefData struct {
	Data map[string]interface{} `json:"beef"`
}

func TestGetJson(t *testing.T) {
	content := "abc def ghi"
	got := GetJson(content)
	t.Logf("got:%v", got)

	var gotObj BeefData
	err := json.Unmarshal([]byte(got), &gotObj)
	if err != nil {
		panic(err)
	}
	t.Logf("gotObj:%v", gotObj)

	want := BeefData{
		Data: map[string]interface{}{
			"abc": 1.0,
			"def": 1.0,
			"ghi": 1.0,
		},
	}

	areEqual := cmp.Equal(want, gotObj)
	//	num := gotObj.Data["abc"].(float64)
	//	t.Logf("abc:%f", num)
	if !areEqual {
		t.Errorf("GetJson(\"%s\") = %s", content, got)
		t.Errorf("Compare %v : want %v", gotObj, want)
	}
}

func TestGetJson2(t *testing.T) {
	content := "abc,. def. , ghi, abc. ghi... def"
	got := GetJson(content)
	t.Logf("got:%v", got)

	var gotObj BeefData
	err := json.Unmarshal([]byte(got), &gotObj)
	if err != nil {
		panic(err)
	}
	t.Logf("gotObj:%v", gotObj)

	want := BeefData{
		Data: map[string]interface{}{
			"abc": 2.0,
			"def": 2.0,
			"ghi": 2.0,
		},
	}

	areEqual := cmp.Equal(want, gotObj)
	if !areEqual {
		t.Errorf("GetJson(\"%s\") = %s", content, got)
		t.Errorf("Compare %v : want %v", gotObj, want)
	}
}
