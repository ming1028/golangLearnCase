## protoc命令

```
protoc -I ./proto --go_out ../../ --go-grpc_out ../../ search.proto
```

* -I 指定.proto文件搜索目录
* --go_out --go-grpc_out 指定生成文件目录
* 默认情况下生成的文件使用import目录，依据go_package生成目录