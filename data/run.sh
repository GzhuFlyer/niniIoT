#!/bin/bash

# 定义服务的目录
ETCD_DIR="./etcd"
MYSQL_DIR="./mysql"
REDIS_DIR="./redis"
NGINX_DIR="./nginx"

# 启动所有服务
start_services() {
    # 启动 etcd 服务
    cd $ETCD_DIR && docker-compose up -d &
    
    # 启动 mysql 服务
    cd $MYSQL_DIR && docker-compose up -d &
    
    # 启动 redis 服务
    cd $REDIS_DIR && docker-compose up -d &

    # 启动 nginx 服务
    cd $NGINX_DIR && docker-compose up -d &
    
    # 等待所有后台进程执行完毕
    wait
}

# 停止所有服务
stop_services() {
    # 停止 etcd 服务
    cd $ETCD_DIR && docker-compose down &
    
    # 停止 mysql 服务
    cd $MYSQL_DIR && docker-compose down &
    
    # 停止 redis 服务
    cd $REDIS_DIR && docker-compose down &

    # 停止 nginx 服务
    cd $NGINX_DIR && docker-compose down &
    
    # 等待所有后台进程执行完毕
    wait
}

# 检查传入的参数
if [ "$1" == "start" ]; then
    echo "Starting services..."
    start_services
elif [ "$1" == "stop" ]; then
    echo "Stopping services..."
    stop_services
else
    echo "Starting services..."
    start_services
    # echo "Usage: ./run.sh {start|stop}"
    exit 1
fi
