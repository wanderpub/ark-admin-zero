### 自动生成API代码

`goctl api go -api doc.api -dir ../`

`goctl api go -api admin.api -dir ../app/admin`

#### 自动生成RPC代码

`goctl rpc proto -src user.proto -dir ../rpc/user`

#### model生成

`goctl model mysql datasource -url="root:123456@tcp(172.16.16.188:3306)/task" -table="*" -dir ../model -c`

#### 运和
go run index.go -f etc/index.yaml