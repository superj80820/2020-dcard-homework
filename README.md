# 2020-dcard-homework

## 驗收步驟

1. 啟動 Golang Server 與 Redis: `$ docker-compose -f docker-compose.prod.yaml up`
2. 開啟網頁至`localhost:3000`，如果一分鐘內超過 60 次 requests 顯示 Error，否則顯示 requests 數。

- 如需驗證 race-condition 可閱讀 [E2E](#e2e-test)

## 功能

- 每個 IP 每分鐘僅能接受 60 個 requests
- 在首頁顯示目前的 request 量,超過限制的話則顯示 “Error”,例如在一分鐘內第 30 個 request 則顯示 30，第 61 個 request 則顯示 Error
- 可以使用任意資料庫，也可以自行設計 in-memory 資料結構，並在文件中說明理由
- 請附上測試

## 需要安裝

- docker-compose
- docker

## Development

```bash
$ docker-compose -f docker-compose.dev.yaml up
```

## Production

### 使用 K8s

```bash
$ kubectl apply -f ./k8s/redis-deployment.yaml,./k8s/redis-service.yaml,./k8s/server-deployment.yaml,./k8s/server-service.yaml
$ minikube service server
```

### 使用 Docker

```bash
$ docker-compose -f docker-compose.prod.yaml up
```

## 推送至 Dockerhub

```bash
$ ./script/push-to-dockerHub.sh
```

## E2E Test

```bash
$ go run ./test/E2E/race-condition.go
```

## Unit Test

```bash
$ go test ./...
```
