## USTC-CAS 模拟登录
学校在前段时间开始推统一认证系统，有一个好处在于不用输入验证码。
### 用法
- 准备工作
    - 安装
        ```bash
        go get github.com/eipi10ydz/ustccas
        ```
    - 获取[中国科学技术大学校园信息门户](http://i.ustc.edu.cn/)下部"统一认证"处需要模拟登录的页面地址
    - 使用两个 `get client` 函数中其中一个成功获得被 `i.ustc.edu.cn` 认证的 `client`
    - 使用该 `client` 对上面得到的页面地址进行一次 `Get` 请求后，就成功获得对应条目的权限，可以直接请求得到页面
- 举例
    - 想要登录教务系统，所以得到页面地址为 `https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do`
    - 使用 `get client` 函数登录，成功获得被 `i.ustc.edu.cn` 认证的 `client`
    - 请求网址
        ```go
        client.Get("https://passport.ustc.edu.cn/login?service=http%3A%2F%2Fmis.teach.ustc.edu.cn/casLogin.do")
        ```
    - 成功后即可请求页面，比如查询成绩
        ```go
        client.Post("http://mis.teach.ustc.edu.cn/querycjxx.do", "", nil)
        ```
    - 示例代码
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
- 注意事项
    - 编码由于是 `GBK`，如果要转 `UTF-8` 可能需要其它包(例如:[mahonia](https://github.com/axgle/mahonia))