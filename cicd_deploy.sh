#!/bin/bash

set -eu

function usage() {
  cat <<EOM
  -s          Stage
  -h          Show Usage
EOM
  exit 2
}

Stage=""

while getopts s:h: OPT; do
  case $OPT in
  "s") Stage=$OPTARG ;;
  '-h' | '--help' | *) usage ;;
  esac
done

echo "DeployStage: ${Stage}"

#  CHANGESET_OPTION="--no-execute-changeset"

# # 引数に'dploy'文字列がある場合のみデプロイ
#  if [ $# = 1 ] && [ $1 = "deploy" ]; then
#   echo "deploy mode"
#   CHANGESET_OPTION=""
# fi

# CFN_TEMPLATE=template.yml
# CFN_STACK_NAME=SeatInfo

# TODO: Deploy時のパラメーターは環境変数とかにする

echo "--------- CodeStar Stack ----------"
# CodeStar Stack
aws cloudformation deploy \
  --stack-name "${Stage}-SeatInfo-codeStarConnection" \
  --template-file aws/codeStarConnection/template.yml \
  --parameter-overrides \
  Stage=${stage}

echo "--------- CodePipeline Stack ----------"
# CodePipeline Stack
aws cloudformation deploy \
  --stack-name "${Stage}-SeatInfo-codePipeline" \
  --template-file aws/codePipeline/template.yml \
  --capabilities CAPABILITY_NAMED_IAM \
  --parameter-overrides \
  Stage=${Stage} \
  BranchName=fix/codebuild#7
