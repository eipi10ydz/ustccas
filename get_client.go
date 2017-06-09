package ustccas

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Find token in the login html, which will be used in post data.
// For token is a hidden input, and the first one,
// so it's easy to get the token by finding first input tag in login html and get its' value.
// found is used to mark whether the token is found.
func FindToken(node *html.Node, found *bool) string {
	var res string
	if node.Type == html.ElementNode && node.Data == "input" {
		for _, a := range node.Attr {
			if a.Key == "value" && !*found {
				// fmt.Println(a.Val)
				res = a.Val
				*found = true
				return res
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if !*found {
			res = FindToken(c, found)
		}
	}
	return res
}

// Get environment variable with upper case only, if not exist, return error.
// Othewise return result of the environment variable.
func GetEnvWrapper(env string) (res string, err error) {
	err = nil
	res = os.Getenv(strings.ToUpper(env))
	if len(strings.TrimSpace(res)) == 0 {
		err = errors.New("cannot get environment variable: " + env)
	}
	return res, err
}

func GetClient(username string) (client *http.Client, err error) {
	var password string
	if password, err = GetEnvWrapper("password"); err != nil {
		return nil, err
	}
	client, err = GetClientAllPara(username, password)
	return client, err
}

func GetClientAllPara(username string, password string) (client *http.Client, err error) {
	client = &http.Client{}

	defer func() {
		if rec := recover(); rec != nil {
			log.Fatal(err)
		}
	}()

	if len(username) == 0 {
		if username, err = GetEnvWrapper("username"); err != nil {
			panic(err.Error())
		}
	}
	client.Jar, err = cookiejar.New(nil)
	URL := "http://i.ustc.edu.cn/login"

	var resp *http.Response
	var token string
	found := false

	if resp, err = client.Get(URL); err != nil {
		panic(err.Error())
	} else {
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			panic(err.Error())
		} else {
			var root *html.Node
			if root, err = html.Parse(strings.NewReader(string(body))); err != nil {
				panic(err.Error())
			}
			token = FindToken(root, &found)
		}
	}
	resp.Body.Close()

	form := url.Values{}
	form.Set("_token", token)
	form.Set("login", username)
	form.Set("password", password)

	passportURL := "https://passport.ustc.edu.cn/login?service="
	passportURL += url.QueryEscape(URL + "?jump=/")

	if resp, err = client.PostForm(passportURL, form); err != nil {
		panic(err.Error())
	}

	return client, err
}
