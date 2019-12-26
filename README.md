# protobuf的发送和接收

client:
1. 构建protobuf对象
2. protobuf对象序列化
3. 分2两次发送序列化字节流，并且在第二次发送的时候增加几次额外的字节，以便测试接收端能不能正确解析

server:
1. 读取接收到的字节，直到发现第一个最高位为0的字节
2. 将该字节之前的所有字节当成数据头，并使用varint解码出原本的数据长度
3. 继续读取字节，直到累积的字节数量达到步骤2解码的长度（未加入超时设置）
4. 使用proto api解码