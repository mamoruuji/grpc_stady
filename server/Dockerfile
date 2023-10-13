FROM golang:1.21-bookworm

WORKDIR /go/src/github.com/mamoruuji/grpc_stady_api

COPY api/go.mod ./
COPY api/go.sum ./

# VScodeの拡張機能の警告対策
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install github.com/cweill/gotests/...@latest
RUN go install github.com/josharian/impl@latest

# DB操作用ORMライブラリ
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
# RUN go get -v github.com/volatiletech/null/v8

# connect-goの
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

ENV TZ=Asia/Tokyo
