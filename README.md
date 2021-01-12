# halia-chat

基于 [halia](https://github.com/halia-group/halia) 开发的聊天室。
采用OOP/组件化开发，零成本扩展

## 术语解释

+ Packet：数据包，请求/响应都视为Packet
+ Processor: 数据包处理器，每个数据包有独立的处理器
+ ProcessorFactory：负责映射`Opcode(数据包命令号)`和`Processor`
+ PacketFactory：负责映射`Opcode`和具体的Packet子类

## 运行

`launcher`目录下依次运行`server`,`client`即可

服务端输出
```text
time="2021-01-12T20:37:35+08:00" level=info msg=started addr=":8080" component=server network=tcp pid=9652
time="2021-01-12T20:37:44+08:00" level=debug msg=connected component=handler peer="127.0.0.1:59313"
```
客户端输出
```text
time="2021-01-12T20:37:44+08:00" level=debug msg=connected component=handler
time="2021-01-12T20:37:44+08:00" level=debug msg="packet Pong{}" component=handler
time="2021-01-12T20:37:44+08:00" level=debug msg="packet RegisterResp{Code=0,Message=注册成功}" component=handler
time="2021-01-12T20:37:44+08:00" level=debug msg="packet LoginResp{Code=0,Message=登录成功}" component=handler
time="2021-01-12T20:37:44+08:00" level=debug msg="packet ChatMessage{Time=2021-01-12 20:37:44 +0800 CST,Publisher=system,MsgType=1,Message=<xialei>已登录}" component=handler
2021-01-12 20:37:44 +0800 CST <system>: <xialei>已登录
time="2021-01-12T20:37:44+08:00" level=debug msg="packet ChatMessage{Time=2021-01-12 20:37:44 +0800 CST,Publisher=xialei,MsgType=1,Message=大家好}" component=handler
2021-01-12 20:37:44 +0800 CST <xialei>: 大家好
```