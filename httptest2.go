package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	http.HandleFunc("/requestAndResponse", requestAndResponse)

	http.HandleFunc("/getVersion", getVersion)

	http.HandleFunc("/ipAndStatus", statusCode)

	http.HandleFunc("/healthz", healthztest)

	err := http.ListenAndServe(":888", nil)
	if nil != err {
		log.Fatal(err)
	}
}

//功能1
func requestAndResponse(response http.ResponseWriter, request *http.Request) {
	println("调用requestAndResponse接口")
	headers := request.Header //header是Map类型的数据
	println("传入的hander：")
	for header := range headers { //value是[]string
		//println("header的key：" + header)
		values := headers[header]
		for index, _ := range values {
			values[index] = strings.TrimSpace(values[index])
			//println("index=" + strconv.Itoa(index))
			//println("header的value：" + values[index])

		}
		//valueString := strings.Join(values, "")
		//println("header的value：" + valueString)
		println(header + "=" + strings.Join(values, ","))        //打印request的header的k=v
		response.Header().Set(header, strings.Join(values, ",")) // 遍历写入response的Header
		//println()

	}
	fmt.Fprintln(response, "Header全部数据:", headers)
	io.WriteString(response, "succeed")

}

// 功能2
func getVersion(response http.ResponseWriter, request *http.Request) {
	println("调用getVersion接口")
	envStr := os.Getenv("VERSION")
	//envStr := os.Getenv("HADOOP_HOME")
	//println("系统环境变量：" + envStr) //可以看到 C:\soft\hadoop-3.3.1	Win10需要重启电脑才能生效

	response.Header().Set("VERSION", envStr)
	io.WriteString(response, "succeed")

}

// 功能3
func statusCode(response http.ResponseWriter, request *http.Request) {

	form := request.RemoteAddr
	println("Client_ip:port=" + form)
	ipStr := strings.Split(form, ":")
	println("Client_ip=" + ipStr[0])

	println("Client_response_code=" + strconv.Itoa(http.StatusOK))

	io.WriteString(response, "succeed")
}

//功能4
func healthztest(response http.ResponseWriter, request *http.Request) {
	println("调用healthz接口")
	response.WriteHeader(200)
	io.WriteString(response, "succeed")
}
