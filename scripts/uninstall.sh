#!/bin/bash

rc="\033[31m" # red
yc="\033[33m" # yellow
gc="\033[32m" # green
dc="\033[0m" # default


# exit if current user is not root
function checkRoot() {
    echo "Check user..."
    user=$(whoami)

    if [ "$user" != "root" ]; then
        echo "Please, run this script with root privileges"
        echo "${rc}Aborted${dc}"
        exit 1
    fi
}

# exit if docker is not installed
function checkDocker() {
    echo "Check docker..."

    if [ ! "$(command -v docker)" ]; then
        echo "${rc}ERROR: Docker is not installed${dc}"
        echo "Install docker manually and run this script again"
        echo "${rc}Aborted${dc}"
        exit 1
    fi
}

# exit if nginx is not installed
function checkNginx() {
    echo "Check nginx..."

    if [ ! "$(command -v nginx)" ]; then
        echo "${rc}ERROR: Nginx is not installed${dc}"
        echo "Install nginx manually and run this script again"
        echo "${rc}Aborted${dc}"
        exit 1
    fi
}

# remove client files (HTML/CSS/JS) for Nginx
function removeClient() {
    echo "Remove client files..."
    rm -rf /var/www/werminal/client
}

# stop server process and remove server files
function removeServer() {
    echo "Stop and remove server..."

    echo "Remove server files..."
    rm -rf /var/www/werminal/server

    echo "${yc}Stop server..."
    echo "To stop server process you should find it and kill."
    echo "To find server process use \"pgrep -af /var/www/werminal/server/bin/app\""
    echo "This command will print out the proccess PID and run binary"
    echo "Use \"sudo kill -9 {PID}\" to kill server process${dc}"
}

# remove nginx config for this app
function clearNginx() {
    rm -f /etc/nginx/sites-enabled/werminal.conf
    rm -f /etc/nginx/sites-available/werminal.conf

    nginx -s reload
}

function finish() {
    echo "${gc}Werminal successfully uninstalled!${dc}"
}

checkRoot
checkDocker
checkNginx
clearNginx
removeServer
removeClient
finish
