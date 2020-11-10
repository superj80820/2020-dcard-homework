# 2020-dcard-homework

## 供 Dcard 人員驗收，謝謝

1. 啟動 Golang Server 與 Redis: `$ docker-compose -f docker-compose.prod.yaml up`
2. 開啟網頁至`localhost:3000`，如果一分鐘內超過 60 次 requests 顯示 Error，否則顯示 requests 數。

- in-memory 資料結構採用 Redis，理由是可以在分佈式系統連結各個微服務，並且有著良好的性能
- 測試:
  - [Unit Test](#unit-test)
  - [E2E](#e2e-test): 可驗證是否有 race-condition
- 整體設計採用[bxcodec](https://github.com/bxcodec)的[Clean Architecture](https://github.com/bxcodec/go-clean-arch)架構，我有針對此架構寫了以下文章:
  - [你的 Backend 可以更有彈性一點 - Clean Architecture 概念篇](https://ithelp.ithome.com.tw/articles/10240228)
  - [你的 Backend 可以更有彈性一點 - Clean Architecture 實作篇](https://ithelp.ithome.com.tw/articles/10241479)
  - [你的 Backend 可以更有彈性一點 - Clean Architecture 測試篇](https://ithelp.ithome.com.tw/articles/10241698)

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
