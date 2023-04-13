# golang TCP

## 什么样的操作会导致 close_wait和 time_wait状态的出现

close_wait: 客户端发起关闭连接的操作，在服务端在应用程序中未正确处理关闭。（如：没有调用conn.Close）

time_wait: 客户端/者服务端 主动发起关闭的一端。在tcp完全关闭后会进入这个状态，客户端出现这个状态是没有关系的，由于服务端需要监听端口，处理大量的链接，如果服务端有大量的time_wait将影响后续链接的建立。所以，一般由客户端发起关闭操作。

程序模拟了服务端产生大量time_wait：
````markdown
1. 启动sever
2. 启动`justconnect`文件夹下的main.go程序
````

使用shell查看Linux的网络端口：
```shell
$ netstat -p tcp -n | grep 'TIME*'

tcp4       0      0  127.0.0.1.11123        127.0.0.1.52315        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52316        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52313        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52317        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52318        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52319        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52320        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52321        TIME_WAIT
tcp4       0      0  127.0.0.1.11123        127.0.0.1.52322        TIME_WAIT
```

程序模拟了客户端产生大量close_wait：
````markdown
1. 启动sever
2. 启动`neverclose`文件夹下的main.go程序
````

使用shell查看Linux的网络端口：
```shell
$ netstat -p tcp -n | grep -E 'FIN*|CLOSE*'

tcp4       0      0  127.0.0.1.54118        127.0.0.1.11123        CLOSE_WAIT
tcp4       0      0  127.0.0.1.54117        127.0.0.1.11123        CLOSE_WAIT
tcp4       0      0  127.0.0.1.54116        127.0.0.1.11123        CLOSE_WAIT
tcp4       0      0  127.0.0.1.54115        127.0.0.1.11123        CLOSE_WAIT
tcp4       0      0  127.0.0.1.54114        127.0.0.1.11123        CLOSE_WAIT
.
.
.
tcp4       0      0  127.0.0.1.11123        127.0.0.1.55722        FIN_WAIT_2
tcp4       0      0  127.0.0.1.11123        127.0.0.1.55721        FIN_WAIT_2
tcp4       0      0  127.0.0.1.11123        127.0.0.1.55720        FIN_WAIT_2
tcp4       0      0  127.0.0.1.11123        127.0.0.1.55719        FIN_WAIT_2
```