#!/bin/bash

scriptPath=$(dirname "$(realpath "$0")")

# exit if current user is not root
function checkRoot() {
    echo "Check user..."
    user=$(whoami)

    if [ "$user" != "root" ]; then
        echo "Please, run this script with root privileges"
        echo "Aborted"
        exit 1
    fi
}

# exit if docker is not installed
function checkDocker() {
    echo "Check docker..."

    if [ ! "$(docker --version)" ]; then
        echo "Docker is not installed"
        echo "Install docker manually and run this script again"
        echo "Aborted"
        exit 1
    fi
}

# exit if nginx is not installed
function checkNginx() {
    echo "Check nginx..."

    if [ ! "$(nginx -v)" ]; then
        echo "Nginx is not installed"
        echo "Install nginx manually and run this script again"
        echo "Aborted"
        exit 1
    fi
}

# setup files (HTML/CSS/JS) for Nginx
function setupClientFiles() {
    echo "Install client files..."
    mkdir -p /var/www/werminal/client
    cp -r "$scriptPath"/../website /var/www/werminal/client

    cd "$scriptPath" || exit 1
    docker build -t werminal_client:latest -f DockerfileClient ..
    docker run --rm --volume /var/www/werminal/client/node_modules:/client/node_modules:rw werminal_client:latest
    docker image rm werminal_client:latest
}

# compile server part and create system service
function setupServer() {
    echo "Install server..."
    mkdir -p /var/www/werminal/server

    cd "$scriptPath" || exit 1
    echo "Compile server binary..."
    docker build -t werminal_server:latest -f DockerfileServer ..
    docker run --rm --volume /var/www/werminal/server/bin:/server/bin:rw werminal_server:latest
    docker image rm werminal_server:latest

    echo "SERVER_PORT=8851" > /var/www/werminal/server/env

    echo "Create system service..."
    cp ./server.service /etc/systemd/system/werminal.service

    systemctl daemon-reload
    systemctl start werminal.service
    systemctl enable werminal.service
}

# setup nginx
function setupNginx() {
    cp ./werminal.conf /etc/nginx/sites-available/werminal.conf
    ln -s /etc/nginx/sites-available/werminal.conf /etc/nginx/sites-enabled/werminal.conf

    if [ ! "$(nginx -t)" ]; then
        echo "Invalid Nginx config syntax"
        echo "Check your nginx config files and run this script again"
        echo "Aborted"
        exit 1
    fi

    nginx -s reload
}

function finish() {
    echo "Installation is finished!"
    echo "Your app run at http://localhost:8803"

    echo "Pay attention, please!"
    echo "If app is not run, check line \"include /etc/nginx/sites-enabled/*;\" in http directive in main nginx config (/etc/nginx/nginx.conf)"
    echo "If this line missing then insert it and restart nginx with \"nginx -s reload\""
}

checkRoot
checkDocker
checkNginx
setupClientFiles
setupServer
setupNginx
finish
