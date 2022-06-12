

# 模块三：docker



- ## 构建本地镜像

  ### 1.将上一节中go源码上传至云主机

  ![image-20220612193006160](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612193006160.png)

  

  `源码如下`

  ```
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
     //println("系统环境变量：" + envStr) //可以看到 C:\soft\hadoop-3.3.1  Win10需要重启电脑才能生效
  
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
  ```

  ### 2.go build 构建源码测试在云主机环境是否可用：

  ![image-20220612193152724](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612193152724.png)

  在源码的上级目录上新建dockerfile：

  ![image-20220612192904946](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612192904946.png)

  

  ### 3.构建本地镜像：docker build -t gohttpserver:v1 ，结果如图：

  ![image-20220612193527195](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612193527195.png)

  

- ## 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化

  ### Dockerfile如下：

  `FROM ubuntu
  RUN echo '这是一个基于ubuntu且应用为httpserver的镜像'
  COPY go/gohttpserver /root/cloud-native/httpserver/

  ENTRYPOINT ["/root/dockerstudy/httpserver/gohttpserver"]`

  

- ## 将镜像推送至 docker 官方镜像仓库

  ### 链接：https://hub.docker.com/repository/docker/patrick090747/httpserver

- ## 通过 docker 命令本地启动 httpserver

  ### 启动httpserver容器：

  ![image-20220612193745319](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612193745319.png)

  无问题，与本地go构建完运行结果一致

- ## 通过 nsenter 进入容器查看 IP 配置

  ### 1.docker ps查看当前在跑的容器，得到容器id：

  ![image-20220612193853316](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612193853316.png)

### 2.通过docker inspect -f {{.State.Pid}}命令找到对应容器id的pid

![image-20220612194847114](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612194847114.png)

### 3.通过pid进入网络namespace，查看网络配置

![image-20220612194824178](C:\Users\12894\AppData\Roaming\Typora\typora-user-images\image-20220612194824178.png)

