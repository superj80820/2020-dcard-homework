# 2020-dcard-homework

## Feature

- 每個 IP 每分鐘僅能接受 60 個 requests
- 在首頁顯示目前的 request 量,超過限制的話則顯示 “Error”,例如在一分鐘內第 30 個 request 則顯示 30，第 61 個 request 則顯示 Error
- 可以使用任意資料庫，也可以自行設計 in-memory 資料結構，並在文件中說明理由
- 請附上測試

## TODO

- [x] 設計 `docker-compose`: Golang Server, Redis 配置
- [x] 設計 `Redis repository`: 讀寫 Redis
- [x] 設計 `limit-rate usecase`: 每分鐘只能接受 60 個 requests
  - 設計 SET: IP:requestsCount
  - SET 必須要有 60s expire time，每次更新都重置 expire time
- [x] 撰寫首頁，使用 `limit-rate usecase`，如果 usecase 成功就顯示 requests count，失敗就顯示 Error
- [ ] 撰寫 README
- [ ] 撰寫 k8s yaml 檔案
- [ ] Unit Test

## Development

```bash
$ docker-compose -f docker-compose.dev.yaml up
```

## Production

```bash
$ docker-compose -f docker-compose.prod.yaml up --build
```

## Push to image repository

```bash
$ ./script/push-to-dockerHub.sh
```

## E2E testing

```bash
$ go run ./test/E2E/race-condition.go
```
