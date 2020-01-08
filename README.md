# protobuf的发送和接收

代码: cmd/customProtocolClient.go cmd/customProtocolServer.go 

client:
1. 构建protobuf对象
2. 序列化两个protobuf对象，p1和p2
3. 分多次发送序列化字节流，第一次发送p1部分（测试拆包），第二次发送p1剩余+p2部分（测试粘包），第三次发送p2剩余+额外字节

server:
1. 读取接收到的字节，直到发现第一个最高位为0的字节
2. 将该字节之前的所有字节当成数据头，并使用varint解码出原本的数据长度
3. 继续读取字节，直到累积的字节数量达到步骤2解码的长度（未加入超时设置）
4. 使用proto api解码

# protobuf在grpc中的应用
代码：  
simpleGRPCClient.go simpleGRPCServer.go  
serverSideRPCClient.go serverSideRPCServer.go  
clientSideRPCClient.go clientSideRPCServer.go  
bidirectionalGRPCClient.go bidirectionalGRPCServer.go  

代理：  
目前只发现通过设置环境变脸的方式来通过代理连接到server端

`
os.Setenv("http_proxy", "socks5://127.0.0.1:1080")
`

