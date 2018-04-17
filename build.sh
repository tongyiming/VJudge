#!/bin/bash
workspace=$(cd $(dirname $0) && pwd -P)
cd $workspace

## const
app="vjudge"
cfg="./cfg"


## funcation
function build(){
    local go="/usr/local/go"
    if [[ -d "$go" ]]; then
         export GOROOT="$go"
         export PATH=$GOROOT/bin:$PATH
	 export GOPATH="$(pwd):$(pwd)/deps"
    else
	echo "Go文件不存在"
	exit 1
    fi
    echo "`go version`"
    go build -o vjudge
}

function make_output(){
    local output="./output"
    rm -rf $output &>/dev/null
    mkdir -p $output &>/dev/null
    (
	cp -vrf $app $output && #拷贝二进制文件至output目录
	cp -vrf $cfg $output && #拷贝cfg配置文件至output目录
	cp -vrf control.sh $output && # 拷贝cotrol.sh至output目录
	echo "make output ok" 
    ) || { echo "make output error";rm -rf "./output";exit 2;} 
}

build
make_output

echo "build down"
exit 0
