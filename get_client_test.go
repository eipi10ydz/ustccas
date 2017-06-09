package ustccas

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"testing"

	//"github.com/axgle/mahonia"
)

func Test_Env(test *testing.T) {
	var err error
	var res string
	os.Setenv("USERNAME", "nicaicai")
	if res, err = GetEnvWrapper("uSeRName"); err != nil {
		test.Fatal(err)
	}
	if res != "nicaicai" {
		test.Fatal(err)
	}
}

func Test_Client(test *testing.T) {
	// os.Setenv("PASSWORD", "your password")

	var client *http.Client
	var resp *http.Response
	var err error

	username := "your username"
	var password string
	if password, err = GetEnvWrapper("password"); err != nil {
		test.Fatal(err)
	}

	if client, err = GetClientAllPara(username, password); err != nil {
		test.Fatal(err)
	}
	resp, err = client.Get("https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do")
	resp, err = client.Get("http://mis.teach.ustc.edu.cn/casLoginNext.do")
	defer resp.Body.Close()
	str, _ := ioutil.ReadAll(resp.Body)
	re, _ := regexp.Compile(username)
	if string(re.Find(str)) != username {
		test.Fatal("login failed")
	}
}

func Test_Client_AllPara(test *testing.T) {
	// os.Setenv("PASSWORD", "your password")

	var client *http.Client
	var resp *http.Response
	var err error

	username := "your username"
	var password string
	if password, err = GetEnvWrapper("password"); err != nil {
		test.Fatal(err)
	}

	if client, err = GetClientAllPara(username, password); err != nil {
		test.Fatal(err)
	}
	// Get Ticket to access mis.teach.ustc.edu.cn
	resp, err = client.Get("https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do")
	// Check success by student No
	resp, err = client.Get("http://mis.teach.ustc.edu.cn/casLoginNext.do")
	defer resp.Body.Close()
	str, _ := ioutil.ReadAll(resp.Body)
	re, _ := regexp.Compile(username)
	if string(re.Find(str)) != username {
		test.Fatal("login failed")
	}
}
