package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// GetURL 获取URL的内容
func GetURL() {
	for _, url := range os.Args[1:] {
		// 练习1.8 增加 http:// 前缀
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err: %v\n", err)
			os.Exit(1)
		}
		// 练习1.7 使用io.Copy
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "io.Copy err: %v\n", err)
			os.Exit(1)
		}
		// 练习1.9 新增http响应状态输出
		respStatus := "http resp status code: " + strconv.Itoa(resp.StatusCode)
		_, err = os.Stdout.WriteString(respStatus)
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.Stdout write err: %v\n", err)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch read err :%v\n", err)
		// 	os.Exit(1)
		// }
		// fmt.Println(string(b))
	}
}
