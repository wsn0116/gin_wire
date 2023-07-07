# gin_wire

gin と google wire でつくる REST API のサンプルです。

## gin と wire が入っていなければ入れる
- https://github.com/gin-gonic/gin
- https://github.com/google/wire

## go mod init
```
go mod init example.com/gin_wire
```

## go.mod 生成
```
go mod tidy
```

## wire_gen を生成する
```
make wire
```

## ビルドする
```
make build
```

## サーバ実行
```
./bin/gin_wire
```

## クライアント接続
http://localhost:8080/sample にアクセスすると JSON でレスポンスが返ってきます。
```
{"message":"sample repository"}
```