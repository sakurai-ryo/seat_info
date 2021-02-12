# 空席確認アプリケーション

## ディレクトリ
```
.
├── README.md
├── app
│   ├── Dockerfile
│   ├── controller
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── middleware
│   ├── model
│   ├── service
│   ├── shared
│   │   ├── weakday.go
│   │   └── weakday_test.go
│   └── templates
│       ├── close
│       │   └── inde.html
│       └── open
│           └── index.html
├── app_deploy.sh
├── aws
│   ├── codePipeline
│   │   └── template.yml
│   ├── codeStarConnection
│   │   └── template.yml
│   ├── fargate
│   │   └── template.yml
│   ├── network
│   │   └── template.yml
│   ├── secretManager
│   │   └── template.yml
│   └── waf
│       └── template.yml
└── cicd_deploy.sh
```

## 使用技術
- Go
- Gin
- AWS
  - Cloud Formation (sls or samでも)
    - Fargate
    - ALB
    - WAF
    - DynamoDB
    - CodePipeline
      - Github
      - CodeBuild