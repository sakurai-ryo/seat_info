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
    - Route53
    - DynamoDB
    - CodePipeline
      - Github
      - CodeBuild

## CodeBuild Local
[参考](https://docs.aws.amazon.com/ja_jp/codebuild/latest/userguide/use-codebuild-agent.html)
1. Gitから当該リポジトリをClone
```shell
$ git clone https://github.com/aws/aws-codebuild-docker-images.git
```
2. ビルド
```shell
$ cd aws-codebuild-docker-images/ubuntu/unsupported_images/golang/1.11
$ docker build -t aws/codebuild/golang:1.11 .
```
3. app以下にbuildspec.yml作成
4. コンテナエージェントのプル
```shell
$ docker pull amazon/aws-codebuild-local:latest --disable-content-trust=false
```
4. local build実行
```shell
$ ❯ ./codebuild_build.sh -i seat-info-repository -a tmp
```