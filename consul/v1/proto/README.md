

#proto 常用命令

## proto文件 命令

### 自动在当前目录创建文件夹(go_package 属性值)，并在文件夹种生成go文件

`protoc --go_out=plugins=grpc:. AdxService.proto`


### 自动在当前路径生成go文件

`protoc --go_out=plugins=grpc:.. AdxService.proto`

### proto 文件夹proto文件， 生成到当前路径

`protoc --go_out=plugins=grpc:.. ./proto/AdxService.proto`





## swagger 文档生成

### 1 获取包

`go get "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"`

### 2 生成 swagger json 文件

`protoc --openapiv2_out . --openapiv2_opt allow_merge=true,merge_file_name=dms --openapiv2_opt logtostderr=true ./*.proto`

### 3 安装 swagger 命令

#### 3.1 下载命令

`https://github.com/go-swagger/go-swagger/releases`

#### 3.2 重名名，设置环境变量(可直接复制到 $GOPATH/bin 路径)

`swagger_windows_amd64.exe` => `swagger.exe` 

#### 4 swagger 通过 json 文件 生成服务

`swagger serve -F=swagger ./dms.swagger.json --host=localhost -p=5000`