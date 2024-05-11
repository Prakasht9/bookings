package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gs", "/general_suite", "GET", []postData{}, http.StatusOK},
	{"ls", "/luxury_suite", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make_reservation", "/make_reservation", "GET", []postData{}, http.StatusOK},
	{"reservation_summary", "/reservation_summary", "GET", []postData{}, http.StatusOK},

	{"post-search_availability", "/search_availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-02-01"},
	}, http.StatusOK},

	{"post-search_availability-json", "/search_availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-02-01"},
	}, http.StatusOK},

	{"make_reservation", "/make_reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Michal"},
		{key: "email", value: "John@john.com"},
		{key: "phone", value: "7679359482"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}

		} else {
			values := url.Values{}

			for _, val := range test.params {
				values.Add(val.key, val.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+test.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}

		}
	}

}
