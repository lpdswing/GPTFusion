#!/usr/bin/env bash
# 读取.env版本号
VERSION=$(cat .env | grep -E "^VERSION=" | awk -F '=' '{print $2}')
PREVERSION=$(cat .env | grep -E "^PREVERSION=" | awk -F '=' '{print $2}')
productVersion=$(cat .env | grep -E "^productVersion=" | awk -F '=' '{print $2}')

echo "VERSION: $VERSION"
echo "PREVERSION: $PREVERSION"

export VERSION=$VERSION
export PREVERSION=$PREVERSION

envsubst < template.md > README.md
envsubst < template.en.md > README.en.md
sed -i '.bak' "s/\"productVersion\": \".*\"/\"productVersion\": \"$productVersion\"/g" wails.json
