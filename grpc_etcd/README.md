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