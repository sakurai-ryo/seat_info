#!/bin/bash

set -eu

function usage() {
  cat <<EOM
  -s          Stage
  -t          TargetStack
  -h          Show Usage
EOM
  exit 2
}

Stage=""
Target=""

while getopts s:t:h: OPT; do
  case $OPT in
  "s") Stage=$OPTARG ;;
  "t") Target=$OPTARG ;;
  '-h' | '--help' | *) usage ;;
  esac
done

echo "DeployStage: ${Stage}"
echo "DeployStack: ${Target}"

if [ -z ${Target} ]; then
  echo "TargetStack(-t)を指定してください"
  exit 1
fi

#  CHANGESET_OPTION="--no-execute-changeset"

# # 引数に'dploy'文字列がある場合のみデプロイ
#  if [ $# = 1 ] && [ $1 = "deploy" ]; then
#   echo "deploy mode"
#   CHANGESET_OPTION=""
# fi

# CFN_TEMPLATE=template.yml
# CFN_STACK_NAME=SeatInfo

# TODO: Deploy時のパラメーターは環境変数とかにする

if [ ${Target} = "all" ] || [ ${Target} = "connections" ]; then
  echo "--------- CodeStar Stack ----------"
  # CodeStar Stack
  aws cloudformation deploy \
    --stack-name "${Stage}-SeatInfo-codeStarConnection" \
    --template-file aws/codeStarConnection/template.yml \
    --parameter-overrides \
    stage=${Stage}
fi

if [ ${Target} = "all" ] || [ ${Target} = "codepipeline" ]; then
  echo "--------- CodePipeline Stack ----------"
  # CodePipeline Stack
  aws cloudformation deploy \
    --stack-name "${Stage}-SeatInfo-codePipeline" \
    --template-file aws/codePipeline/template.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameter-overrides \
    stage=${Stage} \
    BranchName=fix/codebuild#7
fi
