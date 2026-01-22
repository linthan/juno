#!/bin/bash

set -exo pipefail
basePath=$(dirname $(dirname $(dirname $(dirname $(readlink -f $0)))))

echo -e "\nBuilding binaries in ${basePath}"
declare -a bins=${@:2}

# 设置 Go proxy 代理（可按需修改）
export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

# 检查 Go 版本并设置 GOTOOLCHAIN，自动适配 go1.21/1.22 及以上
go_version=$(go version | awk '{print $3}' | sed 's/go//') # e.g. 1.22.1
go_major=$(echo "$go_version" | cut -d. -f1)
go_minor=$(echo "$go_version" | cut -d. -f2)

# 只在 go >= 1.21 时设置
if { [ "$go_major" -gt 1 ] || { [ "$go_major" -eq 1 ] && [ "$go_minor" -ge 21 ]; }; }; then
    export GOTOOLCHAIN=local
    echo "Set GOTOOLCHAIN=local for Go $go_version"
else
    echo "Go version is $go_version, not setting GOTOOLCHAIN"
fi


cd ${basePath}
rm -rf ./bin/*
for task in ${bins[@]}
do
    echo building ${task}...
    read pkgPath  <<< $(echo ${task} | awk -F ":" '{ print $1 }')
    read binName  <<< $(echo ${task} | awk -F ":" '{  print $2 }')
    echo pkgPath:${pkgPath}, binName:${binName}
    go build -ldflags "$1" -o ./bin/$(echo $(echo ${binName} -v | awk -F/ '{print $NF}') ${pkgPath})
    echo -e "\n"
done

if [ -d "${basePath}/cmd/task" ];then
    for file in ${basePath}/cmd/*; do
        echo building ${file}...
        echo pkgPath:${file}, binName:$(basename  $file)
        cd ${file}
        go build -ldflags "$1" -o ${basePath}/bin/job/$(echo $(basename  $file))
        echo -e "\n"
    done
fi
