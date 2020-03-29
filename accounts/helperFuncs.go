package accounts

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func getAccountsTestServer(request string) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if request == "single" {
			fmt.Fprintln(w, "eh")
		}
		if request == "multi" {
			fmt.Fprintln(w, "eh")
		}
	}))

	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
}
