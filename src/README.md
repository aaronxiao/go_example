#题目要求

## 题目一要求
1. 设计开发一个 tcp server 服务
2. 设计开发一个 client 测试，与 tcp server 通讯
3. tcp server 服务收到 client 连接的时候，
   保存该连接的信息，如：连接的ip，端口，用户uid(自定整形不重复即可)
   数据保存在(内存 或 文件 或 自己实现简单的cache服务)
4. tcp server 服务支持可以获取 client 连接数及服务本身信息，
   并可以支持重启(如果 client 正在连接操作此服务，保证不影响 client)，停服(如果有数据操作，请保证数据完整性)等操作
5. tcp server 与 client 通讯协议自行设计即可
6. 模拟 10/100 个 client 同时连接tcp server场景
7. 把 tcp server 多个节点部署 (可选)
8. tcp server 多个节点部署的情况下，实现节点的负载均衡 (可选)
9. client 模拟 tcp server 负载均衡的场景 (可选)
10. client 连接信息的数据保存，如果连接数1万，10万，100万该如何设计cache服务保存 (可选)
11. 请对这个tcp server服务进行压测，计算出最大并发连接数和CPS(每秒新建连接数) (可选)


## 题目二要求
1. 设计开发一个数据服务
2. 在数据服务内存(或自行设计cache), mysql, redis, mongodb 插入相同的模拟数据(自行设计，可只选其中两种，mysql，redis，mongodb预留接口即可，无需实现)
3. 通过数据服务进行增删查改操作
4. 数据服务通讯协议自行设计即可
5. 设计开发一个 client 操作数据服务
6. 如果数据服务是多节点部署，如何来保证幂等 (可选)
7. 数据源数据量比较大的场景下，该如何设计 (可选)



#目录说明
公共目录：
* common:实现一致性hash功能 + 获取当前工作路径
* proto: 协议以及状态码
* conf: 配置文件的读取

题一目录：
* userCache: tcp连接的ip和端口号信息 
* tcpServer: 实现的一个tcp服务端
* client: 实现的一个与tcpServer交互的客户端
* clientHttp: 实现负载均衡的tcp客户端
* proxy: 基于一致性hash实现的负载均衡，通过http方式取得tcp服务节点，同时基于tcp的方式进行健康检查


题二目录：
* clientHttp: 实现数据操作的http客户端
* cacheServer: token缓存节点，token类似校验，用于数据操作时的校验，防止出现幂等性问题
* dataServer: 基于http方式实现的数据操作服务端，包括db结构和数据models



其他：
* dataStruct:网上摘录的b树和红黑树的实现
* test: 自己测试用的