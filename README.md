## USTC-CAS Spider Helper
[中文文档](./README_zh.md)
### Usage
- Preparation
    - Install
        ```bash
        go get github.com/eipi10ydz/ustccas
        ```
    - Get website needed from [http://i.ustc.edu.cn/](http://i.ustc.edu.cn/)
    - Use `get client` function to get client authenticated by `i.ustc.edu.cn`
    - Use the client to make a `Get` request to the website gotten in step 2
- Example
    - If the website is `https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do`
    - Use `get client` function to get client authenticated by `i.ustc.edu.cn`
    - Use the client to make a `Get` request to `https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do`
        ```go
        client.Get("https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do")
        ```
    - Now you can inquire about your grade using the code below.
        ```go
        client.Post("http://mis.teach.ustc.edu.cn/querycjxx.do", "", nil)
        ```
    - Example code
        ```go
        package main

        import (
            "fmt"
            "io/ioutil"
            "log"
            "net/http"

            "github.com/eipi10ydz/ustccas"
        )

        func main() {
            var err error
            var client *http.Client
            var resp *http.Response
            // Step 3
            if client, err = ustccas.GetClientAllPara("your student No", "your password"); err != nil {
                log.Fatal(err)
            }

            // Step 4            
            client.Get("https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do")

            resp, err = client.Post("http://mis.teach.ustc.edu.cn/querycjxx.do", "", nil)
            var body []byte
            if body, err = ioutil.ReadAll(resp.Body); err != nil {
                log.Fatal(err)
            }
            fmt.Println(string(body))
        }
        ```