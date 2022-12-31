## boltdb数据库实例
1.项目地址使用git bash安装：go get github.com/boltdb/bolt/...
2.mod文件加入语句
```go
require (
github.com/boltdb/bolt v1.3.1 // indirect
golang.org/x/sys v0.3.0 // indirect
)
```
3.修改mod文件后，需要运行go mod vendor进行更新