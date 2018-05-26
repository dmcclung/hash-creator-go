package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"io"
	"github.com/dylan-jcloud-assignment/hashstore"
	"strconv"
	"time"
)

func setupHttpRequest(method string, url string, uri string, body io.Reader, hf http.HandlerFunc,
	t *testing.T) (*httptest.ResponseRecorder) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}

	if uri != "" {
		req.RequestURI = uri
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hf)

	handler.ServeHTTP(rr, req)

	return rr
}

func assertStatusCodeEquals(expected int, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Status code fail: got status code %v, expected %v", actual, expected)
	}
}

func assertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Assert equals fail: got %v, expected %v", actual, expected)
	}
}

func TestShutdown(t *testing.T) {
	store = new(hashstore.SimpleKVHashStore)
	store.Start()

	rr := setupHttpRequest(http.MethodGet, "/shutdown", "", nil, shutdown, t)

	assertStatusCodeEquals(http.StatusOK, rr.Code, t)
	assertEquals(true, store.IsShutdown(), t)
}

func TestStats(t *testing.T) {
	store = new(hashstore.SimpleKVHashStore)

	rr := setupHttpRequest(http.MethodGet, "/stats","", nil, stats, t)

	assertStatusCodeEquals(http.StatusOK, rr.Code, t)
	assertEquals("{\"average\":\"NaN\",\"total\":\"0\"}", rr.Body.String(), t)
}

func TestHashStoreWrongMethod(t *testing.T) {
	rr := setupHttpRequest(http.MethodGet, "/hash", "",
		strings.NewReader("password=angryMonkey"), hashPost, t)
	assertStatusCodeEquals(http.StatusNotImplemented, rr.Code, t)
}

func TestHashStore(t *testing.T) {
	store = new(hashstore.SimpleKVHashStore)
	store.Start()

	rr := setupHttpRequest(http.MethodPost, "/hash", "",
		strings.NewReader("password=angryMonkey"), hashPost, t)
	assertStatusCodeEquals(http.StatusOK, rr.Code, t)
	assertEquals("1", rr.Body.String(), t)
}

func TestHashGet(t *testing.T) {
	store = new(hashstore.SimpleKVHashStore)
	store.Start()
	id := store.StoreHash("hashedpassword1")
	time.Sleep(100)

	rr := setupHttpRequest(http.MethodGet, "/hash/" + strconv.FormatUint(id, 10),
		"/hash/" + strconv.FormatUint(id, 10), nil, hashGet, t)

	assertStatusCodeEquals(http.StatusOK, rr.Code, t)
	assertEquals("Y5Z14m/HOZwKHWHuWe6/1dq3P60FWZn4MQV5B1hxOvAupssa+8G+n288oupIMnAmIYODcTyODhhTC1LJ3BR6Gw==",
		rr.Body.String(), t)
}

func BenchmarkHashStore(b *testing.B) {
	store = new(hashstore.SimpleKVHashStore)
	store.Start()
	for n := 0; n < b.N; n++ {
		req, err := http.NewRequest(http.MethodPost, "/hash", strings.NewReader("password=angryMonkey"))
		if err != nil {
			b.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(hashPost)

		handler.ServeHTTP(rr, req)
	}
}
