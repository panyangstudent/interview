# 基本概念
我们都知道rpc(Remote Procedure Call，远程过程调用)，是一个计算机通信协议，此协议允许进程间通信。简单来说就是，当A机器上的进程调用B机器上的进程时，A机器上的调用进程会被挂起，而B机器上的进程开始执行。
调用方将参数信息传送给被调用方，然后可以通过被调用方的结果得到返回。市面上目前主流的关于rpc协议实现有grpc和thrift等。而这些rpc框架一般来说都需要解决服务寻址，数据流的序列化和反序列化，网络传输这三个主要问题。

# rpc调用的基本流程
![avater](图片/img.png)
* 服务消费方（Client 客户端）通过本地调用的方式调用服务。
* 客户端存根（Client Stub）接收到调用请求后负责将方法、入参等信息序列化（组装）成能够进行网络传输的消息体。
* 客户端存根（Client Stub）找到远程的服务地址，并且将消息通过网络发送给服务端。
* 服务端存根（Server Stub）收到消息后进行解码（反序列化操作）。
* 服务端存根（Server Stub， 一段代码）根据解码结果调用本地的服务进行相关处理
* 服务端(Server)本地服务业务处理。
* 处理结果返回给服务端存根（Server Stub）。
* 服务端存根（Server Stub）序列化结果。
* 服务端存根（Server Stub）将结果通过网络发送至消费方。
* 客户端存根（Client Stub）接收到消息，并进行解码（反序列化）。
* 服务消费方得到最终结果。

# http和rpc
http和rpc其实不是对立面，我们知道rpc只是一个计算机通信协议框架，通信协议只是其中的一部分。而http协议作为网络七层模型中应用层的协议，他的主要职责是解决如何包装数据。http协议是建立在传输层tcp协议之上，而传输层的tcp协议主要来解决数据传输问题，
但是对于上层应用开发极其不友好，所以就存在了http协议。 除此之外还有常见的socket，socket是针对TCP或UDP的具体接口实现，提供了在传输层进行网络编程的方法。所以对于rpc的具体实现grpc和thrift来说，grpc的底层实现是http2
协议，而thrift的底层实现是tcp协议。因此我们可以认为http和rpc协议是活跃在应用层的网络协议，他们处于同一层级，互相独立且交织。

# thrift架构
![avater](图片/thrift.png)
从上图我们可以看到，thrift仍然是基于rpc的基本调用流程。上图的黄色部分就是用户实现的业务逻辑，接下来service client/write() read()
是thrift根据IDL(接口描述文件)生成的客户端和服务端代码，包括数据的读写部分，对应于rpc调用的基本流程中的client stub和server stub。
Tprotocol用来对数据进行序列化与反序列化，具体方法包括二进制，json或者apache thrift定义的格式。TTransport提供数据数据传输功能

# thrift特性
1. 开发速度快
   编写thrift IDL文件，利用编译器自动生成服务端骨架(skeletons)和客户端桩(stubs)，省去了开发者自定义和维护接口编解码，消息传输，服务器多线程模型等基础工作。
   服务端：只需要按照服务骨架即接口，编写好具体的业务处理逻辑
   客户端：拷贝IDL定义好的客户端桩和服务对象，然后像调用本地方法一样调用远端服务
2. 接口维护简单
3. 学习成本低
4. 多语言支持
5. 稳定且广泛使用


# thrift IDL文件的数据类型
基本类型
   * bool：布尔值
   * byte：8位有符号整数
   * i16：16位有符号整数
   * i32：32位有符号整数
   * i64：:64位有符号整数
   * double：64位浮点数
   * binary：二进制串
结构体类型
   * struct：定义的结构体对象
容器类型
   * list：有序元素列表
   * set：无序无重复元素集合
   * map：有序的key/value键值对
服务类型
   * service：具体对应的服务类
异常类型
   * exception：异常类型

# thrift的传输层
   * TSocket：使用阻塞式IO进行传输
   * TNonblockingTransport：使用非阻塞方式，用于构建异步客户端
   * TFramedTransport：使用非阻塞方式，按块大小进行传输

# thrift的服务端类型
   * TSimpleserver：单线程服务器端，使用标准的阻塞式IO
   * TThreadPoolServer：多线程服务器端，使用标准的阻塞式IO
   * TNonblockingServer：单线程服务器端，使用非阻塞式IO
   * THsHaServer：半同步半异步服务器端，基于非阻塞式IO读写和多线程工作任务处理
   * TThreadedSelectorServer：多线程选择服务器端，对THsHaServer在异步IO模型上进行增强
