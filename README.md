# 空席確認アプリケーション

## ディレクトリ
```
.
├── controller
├── middleware
├── model
    ├── seatInfo.go
├── service
├── shared
└── main.go
```

## 使用技術
- Go
- Gin
- AWS
  - Cloud Formation (sls or samでも)
    - Fargate
    - DynamoDB
    - CodePipeline (code Buildしか使わないからなくてもいいかも)
      - CodeBuild