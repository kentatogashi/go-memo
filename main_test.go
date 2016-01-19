package main

import (
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

func TestRoot(t *testing.T) {
	m := web.New()
	RootRouter(m)
	ts := httptest.NewServer(m)
	defer ts.Close()
	req, _ := http.NewRequest("GET", ts.URL, nil)
	req.SetBasicAuth("pass", "pass")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error("unexpected")
	}

	_, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("invalid status code")
	}
}

func TestContent(t *testing.T) {
	m := web.New()
	ContentRouter(m)
	ts := httptest.NewServer(m)
	defer ts.Close()
	req, _ := http.NewRequest("GET", ts.URL+"/content/index", nil)
	req.SetBasicAuth("pass", "pass")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error("unexpected")
	}
	_, s := ParseResponse(res)
	if s != http.StatusOK {
		t.Error("invalid status code")
	}
}
