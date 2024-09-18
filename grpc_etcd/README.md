## protoc命令

```
protoc -I ./proto --go_out ../../ --go-grpc_out ../../ search.proto
```

* -I 指定.proto文件搜索目录
* --go_out --go-grpc_out 指定生成文件目录
* 默认情况下生成的文件使用import目录，依据go_package生成目录

## protobuf优势

* 二进制数据格式
* 变长编码，整数类型根据大小动态调整所需的大小
* 字段标识符和类型使用一个字节表示，序列化时不会有字段名称，使用字段标识符和类型来确定
* 可选字段，未被设置，序列化之后就不会出现

## go kit

* server端，不同方法可能有不能的数据格式，可以针对每个方法定义自己的压缩、解压方法，方法返回值必须是endpoint.Endpoint函数类型
* client端
* 轮询：根据请求次数（原子性）求余获得

## etcd 

* 使用etcd的resolver.NewBuilder作为grpc的解析器，底层默认将target当为前缀去检索服务，并将数据json反序列化
* zrpc使用etcd服务发现，服务注册不支持manager格式