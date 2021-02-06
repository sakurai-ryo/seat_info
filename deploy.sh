#!/bin/bash

set -eu

CHANGESET_OPTION="--no-execute-changeset"

# 引数に'dploy'文字列がある場合のみデプロイ
if [ $# = 1 ] && [ $1 = "deploy" ]; then
  echo "deploy mode"
  CHANGESET_OPTION=""
fi

CFN_TEMPLATE=template.yml
CFN_STACK_NAME=SeatInfo

# テンプレートの実行
aws cloudformation deploy   
    --stack-name ${CFN_STACK_NAME} \
    --template-file ${CFN_TEMPLATE} ${CHANGESET_OPTION} \
    --parameter-overrides \
        AZ=ap-northeast-1a \
        VpcCidrBlock=192.168.2.0/24 \
        PublicSubnetCidrBlock=192.168.2.0/25
