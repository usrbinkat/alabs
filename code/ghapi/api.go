package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var urlApiUser = "https://api.github.com/users/"
var login = "usrbinkat"

type Reply struct {
	Login       string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func main() {
	name, repoCount, _ := githubInfo(login)
	fmt.Printf("name: %s\n", name)
	fmt.Printf("repo_count: %d\n", repoCount)
}

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	var r Reply
	uri := urlApiUser + url.PathEscape(login)

	// Request user info
	resp, err := http.Get(uri)
	if err != nil {
		//log.Fatalf(err.Error())
		return "", 0, err
	}

	// Check for errors
	if resp.StatusCode != http.StatusOK {
		//log.Fatalf("error: %s", resp.Status)
		return "", 0, fmt.Errorf("%#v - %s", urlApiUser, resp.Status)
	}

	defer resp.Body.Close()

	// Decode user info
	//  body, err := ioutil.ReadAll(resp.Body)
	//  if err != nil {
	//      log.Fatalf(err.Error())
	//  }

	// Print user info
	//  fmt.Println(string(body))

	// Decode response body JSON
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		// log.Fatalf("error decoding: %s", err.Error())
		return "", 0, err
	}

	// Print Reply JSON
	//fmt.Println(r)
	//fmt.Printf("%#v\n", r)

	return r.Login, r.PublicRepos, nil
}
