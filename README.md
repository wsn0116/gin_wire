# gin_wire

gin と google wire でつくる REST API のサンプルです。

## セットアップ

### 1. gin と wire が入っていなければ入れる
- https://github.com/gin-gonic/gin
- https://github.com/google/wire

### 2. go mod init
```
go mod init example.com/gin_wire
```

### 3. go.mod 生成
```
go mod tidy
```

### 4. wire_gen を生成する
```
make wire
```

### 5. ビルドする
```
make build
```

### 6. サーバ実行
```
./bin/gin_wire
```

### 7. クライアント接続
http://localhost:8080/sample にアクセスすると JSON でレスポンスが返ってきます。
```
{"message":"sample repository"}
```