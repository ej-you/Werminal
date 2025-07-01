#!/bin/bash

serverPort=8851
domain="fredcv.ru"
nginxPort=8092
userForServer=danil

scriptPath=$(dirname "$(realpath "$0")")

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

# setup files (HTML/CSS/JS) for Nginx
function setupClientFiles() {
    echo "Install client files..."
    mkdir -p /var/www/werminal/client
    cp -r "$scriptPath"/../website/* /var/www/werminal/client

    cd "$scriptPath" || exit 1
    docker build -t werminal_client:latest -f DockerfileClient ..
    docker run --rm --volume /var/www/werminal/client/node_modules:/client/node_modules:rw werminal_client:latest
    docker image rm werminal_client:latest
}

# compile server part and run server as process
function setupServer() {
    echo "Install server..."
    mkdir -p /var/www/werminal/server

    cd "$scriptPath" || exit 1
    echo "Compile server binary..."
    docker build -t werminal_server:latest -f DockerfileServer ..
    docker run --rm --volume /var/www/werminal/server/bin:/server/bin:rw werminal_server:latest
    docker image rm werminal_server:latest

    echo "Start server at :$serverPort..."
    cd "/home/$userForServer" || cd /
    SERVER_PORT="$serverPort" sudo -u "$userForServer" bash -c "/var/www/werminal/server/bin/app &> /var/www/werminal/server/app.log &"
}

# setup nginx
function setupNginx() {
    cd "$scriptPath" || exit 1

    cp ./werminal.conf /etc/nginx/sites-available/werminal.conf
    ln -s /etc/nginx/sites-available/werminal.conf /etc/nginx/sites-enabled/werminal.conf

    nginx -t
    nginx -s reloadserverPort
}

function finish() {
    echo -e "${gc}Installation is finished!${dc}"
    echo "Your app run at http://$domain:$nginxPort"

    echo -e "${yc}Pay attention, please!"
    echo -e "If app is not run, check line \"include /etc/nginx/sites-enabled/*;\" in http directive in main nginx config (/etc/nginx/nginx.conf)"
    echo -e "If this line missing then insert it and restart nginx with \"nginx -s reload\"${dc}"
}

checkRoot
checkDocker
checkNginx
setupClientFiles
setupServer
setupNginx
finish
