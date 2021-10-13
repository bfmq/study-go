# 云原生
## 使用busybox原因就是比较小，但是busybox实际上网络不工作的，实际情况中可以使用FROM golang或者精简版的centos等等，但是这些元镜像比较大，镜像不好传github了，所以这次元镜像就用busybox了

### 其他作业要求
#### 封镜像命令
#### docker build -t cncamp:1 . --network=host
[root@devops-bj-yz-dx1 k8s-tmp]# docker build -t cncamp:1 . --network=host
Sending build context to Docker daemon  9.474MB
Step 1/6 : FROM busybox
 ---> 16ea53ea7c65
Step 2/6 : EXPOSE 8000
 ---> Running in cc963f426f4b
Removing intermediate container cc963f426f4b
 ---> 0f0a568c8d40
Step 3/6 : ENV VERSION=123456
 ---> Running in e945b7d9b894
Removing intermediate container e945b7d9b894
 ---> 98f36a5757eb
Step 4/6 : COPY homework /root
 ---> 7db903f2d8ea
Step 5/6 : RUN chmod +x /root/homework
 ---> Running in 29c691d2c7c1
Removing intermediate container 29c691d2c7c1
 ---> 3ea871291c82
Step 6/6 : ENTRYPOINT ["/root/homework"]
 ---> Running in 18e30217e2c3
Removing intermediate container 18e30217e2c3
 ---> a183beb27efa
Successfully built a183beb27efa
Successfully tagged cncamp:1
#### 通过 Docker 命令本地启动 httpserver
#### docker run -d -p 30001:8000 a183beb27efa
[root@devops-bj-yz-dx1 k8s-tmp]# docker image ls
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
cncamp              1                   a183beb27efa        28 seconds ago      20.2MB
busybox             latest              16ea53ea7c65        4 weeks ago         1.24MB
openjdk             8-alpine            a3562aa0b991        2 years ago         105MB
[root@devops-bj-yz-dx1 k8s-tmp]# docker run -d -p 30001:8000 a183beb27efa
WARNING: IPv4 forwarding is disabled. Networking will not work.
247db2cad9ed23e0e1c8e14bf89bd2ebab35ee719e0755c451b1dad704144175
[root@devops-bj-yz-dx1 k8s-tmp]# netstat -tpln|grep 30001
tcp6       0      0 :::30001                :::*                    LISTEN      64460/docker-proxy  
[root@devops-bj-yz-dx1 k8s-tmp]# docker exec -it 247db2cad9ed23e0e1c8e14bf89bd2ebab35ee719e0755c451b1dad704144175 /bin/sh
/ # cd /root/
~ # ls
homework
~ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 /root/homework
   10 root      0:00 /bin/sh
   16 root      0:00 ps -ef
~ # netstat -tpln
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 :::8000                 :::*                    LISTEN      1/homework
~ # wget 127.0.0.1:8000/healthz/
Connecting to 127.0.0.1:8000 (127.0.0.1:8000)
saving to 'index.html'
index.html           100% |******************************************************************************************************************************************************************************************************************************|    12  0:00:00 ETA
'index.html' saved
~ # cat index.html 
{"code":200}~ # 
[root@devops-bj-yz-dx1 k8s-tmp]# docker logs -f 247db2cad9ed23e0e1c8e14bf89bd2ebab35ee719e0755c451b1dad704144175
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /healthz/                 --> main.healthzGetHandle (5 handlers)
[GIN-debug] Listening and serving HTTP on :8000
2021/10/13 08:49:31 客户端地址----------------> 127.0.0.1:36725
2021/10/13 08:49:31 给客户端的返回码----------------> 404
[GIN] 2021/10/13 - 08:49:31 | 404 |      326.34µs |       127.0.0.1 | GET      "/"
[GIN-debug] redirecting request 301: /healthz/ --> /healthz/
2021/10/13 08:49:46 客户端地址----------------> 127.0.0.1:36727
2021/10/13 08:49:46 给客户端的返回码----------------> 200
[GIN] 2021/10/13 - 08:49:46 | 200 |     129.098µs |       127.0.0.1 | GET      "/healthz/"
[GIN] 2021/10/13 - 08:50:51 | 200 |     591.788µs |       127.0.0.1 | GET      "/healthz/"
2021/10/13 08:50:51 客户端地址----------------> 127.0.0.1:36737
2021/10/13 08:50:51 给客户端的返回码----------------> 200
2021/10/13 08:51:14 客户端地址----------------> 127.0.0.1:36746
2021/10/13 08:51:14 给客户端的返回码----------------> 200
[GIN] 2021/10/13 - 08:51:14 | 200 |     110.371µs |       127.0.0.1 | GET      "/healthz/"
^C

#### 通过 nsenter 进入容器查看 IP 配置
#### docker inspect xxx 
#### nsenter -t xxx -n ip a
[root@devops-bj-yz-dx1 k8s-tmp]# docker inspect 247db2cad9ed | grep Pid
            "Pid": 66735,
            "PidMode": "",
            "PidsLimit": null,
[root@devops-bj-yz-dx1 k8s-tmp]# nsenter -t 66735 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
97: eth0@if98: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever

