#!/bin/bash

set -eu

function usage {
  cat <<EOM
  -s          Stage
  -h          Show Usage
EOM
  exit 2
}

Stage=""

while getopts s:h: OPT
do
  case $OPT in
    "s" ) Stage=$OPTARG;;
    '-h'|'--help'|* ) usage;;
  esac
done

echo "DeployStage: ${Stage}"

ProjectName="SeatInfo"

# CHANGESET_OPTION="--no-execute-changeset"

# # 引数に'dploy'文字列がある場合のみデプロイ
# if [ $# = 1 ] && [ $1 = "deploy" ]; then
#   echo "deploy mode"
#   CHANGESET_OPTION=""
# fi

# CFN_TEMPLATE=template.yml
# CFN_STACK_NAME=ProjectName


# TODO: Deploy時のパラメーターは環境変数とかにする

echo "---------- Network Stack ----------"
# Network Stack
aws cloudformation deploy \
    --stack-name "${Stage}-${ProjectName}-network" \
    --template-file aws/network/template.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameter-overrides \
        AZa=ap-northeast-1a \
        AZb=ap-northeast-1c \
        VpcCidrBlock=192.168.0.0/16 \
        MainPublicSubnetCidrBlock=192.168.1.0/24 \
        SubPublicSubnetCidrBlock=192.168.2.0/24 \
        Stage=${Stage}

echo "---------- WAF Stack ----------"
# WAF Stack
aws cloudformation deploy \
    --stack-name "${Stage}-${ProjectName}-waf" \
    --template-file aws/waf/template.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameter-overrides \
        Stage=${Stage}

echo "---------- SecretManager Stack ----------"
# SecretManager Stack
aws cloudformation deploy \
    --stack-name "${Stage}-${ProjectName}-secret-manager" \
    --template-file aws/secretManager/template.yml \
    --parameter-overrides \
      Stage=${Stage}

echo "---------- Fargate Stack ----------"
DesiredCount=0
ExistService=$(aws ecs list-task-definitions --region ap-northeast-1 | jq '.taskDefinitionArns[] | select(contains("SeatInfo"))')
if [ -n "$ExistService" ]; then
  DesiredCount=2
fi
echo "DesiredCount: ${DesiredCount}"
# Fargate Stack
aws cloudformation deploy \
    --stack-name "${Stage}-${ProjectName}-fargate" \
    --template-file aws/fargate/template.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameter-overrides \
      ProjectName=${ProjectName} \
      DesiredCount=${DesiredCount} \
      Stage=${Stage}
