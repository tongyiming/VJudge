#!/bin/bash
workspace=$(cd $(dirname $0) && pwd -P)
cd $workspace

app="vjudge"
cfg=cfg/cfg.toml.release
pidfile=var/app.pid
logfile=logs/app.log

function start(){
    mkdir -p var &>/dev/null
    mkdir -p logs &>/dev/null
    
    #check
    check_pid
    local running=$?
    if [[ $running -gt 0 ]];then
	echo "started,pid=$(get_pid)"
	exit 0
    fi

    echo "use cfg file:$cfg"

    #start now
    nohup ./$app -c $cfg >>$logfile 2>&1 &
    local lpid=$!
    sleep 1
    
    check_pid_number $lpid
    local lnum=$?
    if [ "x$lnum" == "x0" ];then
	echo "start failed,pid=$lpid"
	exit 1
    fi

    echo $lpid > $pidfile
    echo "start ok,pid=$lpid"
}

function stop(){
    for(( i=0; i<60; i++));do
	check_pid
	local running=$?
	if [ $running -le 0 ];then
	    echo "stoped,pid=$(get_pid)"
	    return 0
	fi
	# wait to exit
	kill `get_pid` &>/dev/null
	sleep 1
    done

    echo "stop timeout(60s),pid=$(get_pid)"
    return 1
}

function restart(){
    stop
    local lnum=$?
    if [ $lnum == 0 ];then
	start
    fi
}

function status(){
    check_pid
    local running=$?
    if [ $running -gt 0 ];then
	echo -n "running,pid=$(cat $pidfile)"
	return $running
    else
	echo "stoped"
	return 0
    fi
}

## internals
function get_pid(){
    if [ -f $pidfile ];then
	cat $pidfile
    fi
}

function check_pid(){
    local lpid=`get_pid`
    if [ "x_" != "x_$lpid" ];then
	check_pid_number $lpid
	return $?
    fi
    return 0
}

function check_pid_number(){
    local lpid="$1"
    local lprocfile="/proc/$lpid/cmdline"
    if [ ! -f "$lprocfile" ];then
	return 0
    fi

    local lprocnum=$(cat $lprocfile | grep $app | grep -v "PID TTY" | wc -l)
    if [ "x$lprocnum" == "x0" ];then
	return 0
    fi

    local lalivenum=$(ps -p $lpid | grep -v "PID TTY" | wc -l)
    return $lalivenum
}

##########################################################
# main
# action:
#    -start 启动服务
#    -stop 停止服务
#    -status 查看状态(stoped,other)
#########################################################
action=$1
case $action in
	"start" )
		start
		;;
	"stop" )
		stop
		;;
	"restart" )
		restart
		;;
	"status" )
		status
		;;
	* )
		echo "unknow command [$action]"
		echo "usage: [start][stop][restart][status]"
		exit 1
		;;
esac
