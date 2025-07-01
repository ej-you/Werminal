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
        echo -e "${rc}Aborted${dc}"
        exit 1
    fi
}

# exit if docker is not installed
function checkDocker() {
    echo "Check docker..."

    if [ ! "$(command -v docker)" ]; then
        echo -e "${rc}ERROR: Docker is not installed${dc}"
        echo "Install docker manually and run this script again"
        echo -e "${rc}Aborted${dc}"
        exit 1
    fi
}

# exit if nginx is not installed
function checkNginx() {
    echo "Check nginx..."

    if [ ! "$(command -v nginx)" ]; then
        echo -e "${rc}ERROR: Nginx is not installed${dc}"
        echo "Install nginx manually and run this script again"
        echo -e "${rc}Aborted${dc}"
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

    echo -e "${yc}Stop server..."
    echo -e "To stop server process you should find it and kill."
    echo -e "To find server process use \"pgrep -af /var/www/werminal/server/bin/app\""
    echo -e "This command will print out the proccess PID and run binary"
    echo -e "Use \"sudo kill -15 {PID}\" to gracefully shutdown server process${dc}"
    echo -e "If the command above is executed for more than 10 seconds use \"sudo kill -9 {PID}\" to kill server process${dc}"
}

# remove nginx config for this app
function clearNginx() {
    rm -f /etc/nginx/sites-enabled/werminal.conf
    rm -f /etc/nginx/sites-available/werminal.conf

    nginx -s reload
}

function finish() {
    rm -rf /var/www/werminal
    echo -e "${gc}Werminal successfully uninstalled!${dc}"
}

checkRoot
checkDocker
checkNginx
clearNginx
removeServer
removeClient
finish
