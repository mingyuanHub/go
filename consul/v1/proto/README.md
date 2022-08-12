
## proto文件 命令

### 自动在当前目录创建 proto 文件夹，并在文件夹种生成go文件

`protoc --go_out=plugins=grpc:. AdxService.proto`


### 自动在当前路径生成go文件

`protoc --go_out=plugins=grpc:.. AdxService.proto`

### proto 文件夹proto文件， 生成到当前路径

`protoc --go_out=plugins=grpc:.. ./proto/AdxService.proto`