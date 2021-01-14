package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    s := `postgres://user:pass@host.com:5432/path?k=v#f`

    // 使用Parse函数，解析地址是否符合url
    u, err := url.Parse(s)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(u)

    // 可以通过属性获取url的分段数据
    fmt.Println(u.Scheme)
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _:= u.User.Password()
    fmt.Println(p)
    fmt.Println(u.Host)
    // 如果存在域名(ip)以及端口号，可以使用SplitHostPort对其进行分割
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // path
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)
    // query数据
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])

    // 多个query数据也可以进行拼接为一个url
    params := url.Values{}
    params.Set("key1","1")
    params.Set("key2","2")
    u.RawQuery = params.Encode()
    newUrl := u.String()
    fmt.Println(newUrl)

}
