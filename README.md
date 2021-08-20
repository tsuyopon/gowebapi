# gowebapi
go language webapi sample repository using gin web framework

# how to add entrypoint
- 1. vim server/router.go
- 2. append controller
- 3. go run main.go

# reference
- official document
  - https://pkg.go.dev/github.com/gin-gonic/gin
- check context api
  - https://github.com/gin-gonic/gin/blob/master/context.go

# やりたいこと
- 複数エントリポイント
  - chunk
  - swagger
- MVCフレームワーク
  - model作成
- mysql (pool)
- リクエスト時刻を返す
- 指定された値をVaryにして返す


# 参考
https://bsblog.casareal.co.jp/archives/4822
https://bsblog.casareal.co.jp/archives/5121
https://blog.ishikawa.tech/entry/2017/08/28/122327
