package main

import (
    "encoding/base64"
    "fmt"
)

func main() {

    data := "hello world"
    // 使用StdEncoding.EncodeToString方法进行标准base64编码
    sEnc := base64.StdEncoding.EncodeToString([]byte(data))

    fmt.Println(sEnc)

    // 使用DecodeString对base64内容进行解码
    sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))

    // base64编码还有一种为url的编码，使用URLEncoding.EncodeToString, 使用DecodeString对base64内容进行解码
    uEnc := base64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)

    uDec, _ := base64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))

    // 标准 base64 编码和 URL base64 编码的 编码字符串存在稍许不同（后缀可能为 + 和 -）， 但是两者都可以正确解码为原始字符串。

}
