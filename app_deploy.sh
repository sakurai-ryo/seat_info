#!/bin/bash

set -eu

function usage {
  cat <<EOM
  -s          Stage
  -h          Show Usage
EOM
  exit 2
}

stage=""

while getopts s:h: OPT
do
  case $OPT in
    "s" ) stage=$OPTARG;;
    '-h'|'--help'|* ) usage;;
  esac
done

echo "DeployStage: ${stage}"

# CHANGESET_OPTION="--no-execute-changeset"

# # 引数に'dploy'文字列がある場合のみデプロイ
# if [ $# = 1 ] && [ $1 = "deploy" ]; then
#   echo "deploy mode"
#   CHANGESET_OPTION=""
# fi

# CFN_TEMPLATE=template.yml
# CFN_STACK_NAME=SeatInfo


# TODO: Deploy時のパラメーターは環境変数とかにする

echo "---------- Network Stack ----------"
# Network Stack
aws cloudformation deploy \
    --stack-name "${stage}-SeatInfo-network" \
    --template-file aws/network/template.yml \
    --parameter-overrides \
        AZa=ap-northeast-1a \
        AZb=ap-northeast-1c \
        VpcCidrBlock=192.168.0.0/16 \
        MainPublicSubnetCidrBlock=192.168.1.0/24 \
        SubPublicSubnetCidrBlock=192.168.2.0/24 \
        Stage=${stage}

echo "---------- WAF Stack ----------"
# WAF Stack
aws cloudformation deploy \
    --stack-name "${stage}-SeatInfo-waf" \
    --template-file aws/waf/template.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameter-overrides \
        Stage=${stage}

echo "---------- SecretManager Stack ----------"
# SecretManager Stack
aws cloudformation deploy \
    --stack-name "${stage}-SeatInfo-secret-manager" \
    --template-file aws/secretManager/template.yml \
    --parameter-overrides \
      Stage=${stage}

echo "---------- Fargate Stack ----------"
DesiredCount=0
ExistService=$(aws ecs list-task-definitions --region ap-northeast-1 | jq '.taskDefinitionArns[] | select(contains("SeatInfo"))')
if [ -n "$ExistService" ]; then
  DesiredCount=2
fi
echo "DesiredCount: ${DesiredCount}"
# Fargate Stack
aws cloudformation deploy \
    --stack-name "${stage}-SeatInfo-fargate" \
    --template-file aws/fargate/template.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameter-overrides \
      ProjectName=SeatInfo \
      DesiredCount=1 \
      Stage=${stage}
