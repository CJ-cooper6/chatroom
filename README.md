# 聊天室



## 介绍

简易多人聊天室，使用go语言net包实现server端，客户端使用linux nc命令。

## 功能实现

- 群聊

- 查看在线用户

- 用户改名

- 主动退出（ctrl + c）

- 超时退出

  



## 主要流程图



![2021-09-09 16-34-48 的屏幕截图](https://gitee.com/CJ-cooper6/picgo/raw/master/2021-09-09%2016-34-48%20%E7%9A%84%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE.png)

## 具体功能代码实现

##### 监听全局的message通道

![image-20210909204850621](https://gitee.com/CJ-cooper6/picgo/raw/master/image-20210909204850621.png)



##### 将用户通道里的信息写回给客户端

![image-20210909205124114](https://gitee.com/CJ-cooper6/picgo/raw/master/image-20210909205124114.png)



##### 查看在线人数功能 （who命令）

![image-20210909212300286](https://gitee.com/CJ-cooper6/picgo/raw/master/image-20210909212300286.png)



##### 修改用户名称

![image-20210909214400030](https://gitee.com/CJ-cooper6/picgo/raw/master/image-20210909214400030.png)



##### 用户退出和超时退出

通过select语句和time包实现

![image-20210910225847092](https://gitee.com/CJ-cooper6/picgo/raw/master/image-20210910225847092.png)
