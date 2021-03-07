# 空席確認アプリケーション

## ディレクトリ
```
├── README.md
├── app
│   ├── Dockerfile
│   ├── buildspec.yml
│   ├── codebuild_build.sh
│   ├── controller
│   │   ├── controller.go
│   │   └── request.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── shared
│   │   ├── weakday.go
│   │   └── weakday_test.go
│   ├── static
│   │   ├── css
│   │   │   ├── bootstrap.css
│   │   │   ├── open.css
│   │   │   └── smart.css
│   │   └── images
│   │       ├── 10do_logo.jpg
│   │       └── shop.jpg
│   └── templates
│       ├── close
│       │   └── close.html
│       └── open
│           └── open.html
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
├── cicd_deploy.sh
└── db.json
```

## 使用技術
- Go
- Gin
- AWS
  - Cloud Formation (sls or samでも)
    - Fargate
    - ALB
    - WAF
    - Route53
    - DynamoDB
    - CodePipeline
      - Github
      - CodeBuild
