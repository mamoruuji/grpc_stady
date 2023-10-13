# 0ベースから作成する場合の実装手順
0. go.mod, go.sumファイル作成
- go.modの項目 *module* と *go* を追記
1. docker起動
2. bufライブラリの設定
- ``` $ buf mod init ```
  - ``` buf lint ```の警告対策でファイルの項目 *lint* の最後尾に追加 ```rpc_allow_google_protobuf_empty_requests: true```
- buf.gen.yamlファイル作成、以下のように追記
```
version: v1
plugins:
  - plugin: buf.build/protocolbuffers/go
    out: gen
    opt: paths=source_relative
  - plugin: buf.build/bufbuild/connect-go
    out: gen
    opt: paths=source_relative
```
3. connect-goのコードを生成
- todo.protoファイル作成
  - ``` $ buf lint ``` エラー発生しないことを確認
    - ``` $ buf generate ``` connect-goのコード生成

4. ORMツール（SQLBoiler）でDB操作可能にする
- sqlboiler.tomlファイル作成
  - 接続設定追記
    - ``` $ sqlboiler psql ``` モデルのコード生成

5. 動作確認
- ``` $ go mod tidy ``` ライブラリの依存関係を修正
  - ``` $ go run ./cmd/server/main.go ``` サーバ起動
    - リクエスト
```
curl \
--header "Content-Type: application/json" \
--data '{}' \
http://localhost:8080/proto.todo.v1.TodoService/ListTodos
```
※別コンテナからcurlを実行する際
Dockerの仕様で *localhost* を *grpc_stady_server* に変えれば接続可能になる
