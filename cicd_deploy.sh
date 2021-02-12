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
# Network Stack
aws cloudformation deploy \
    --stack-name "${stage}-SeatInfo-network" \
    --template-file aws/network/template.yml \
    --parameter-overrides \
        AZ=ap-northeast-1a \
        VpcCidrBlock=192.168.2.0/24 \
        PublicSubnetCidrBlock=192.168.2.0/25 \
        Stage=${stage}
