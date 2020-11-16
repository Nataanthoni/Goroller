package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rollbar/rollbar-go"
)

func main() {
	rollbar.SetToken("c9fd87f675544261a2f17d507db095b6")
	rollbar.SetEnvironment("production")                     // defaults to "development"
	rollbar.SetCodeVersion("v2")                             // optional Git hash/branch/tag (required for GitHub integration)
	rollbar.SetServerHost("web.1")                           // optional override; defaults to hostname
	rollbar.SetServerRoot("github.com/nataanthoni/goroller") // path of project (required for GitHub integration and non-project stacktrace collapsing)

	route := mux.NewRouter()
	err := route.HandleFunc("/get", DoSomething).Methods("GET")
	if err != nil {
		rollbar.Critical(err)
	}
	rollbar.Info("Message about my respone and more")

	rollbar.Wait()
	http.ListenAndServe(":8090", nil)
	println("Connecting on port 8090")
}

//DoSomething ...
func DoSomething(w http.ResponseWriter, r *http.Request) {
	url := "https://api.github.com/users/tensorflow"
	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Print(string(body))

}
