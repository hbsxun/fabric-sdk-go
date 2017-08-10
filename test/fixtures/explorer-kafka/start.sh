#!/bin/bash
function rmAllContainers() {
    docker stop $(docker ps -a -q)
    docker rm -f $(docker ps -a -q)
}

function rmImages() {
    if [ "$#" -lt 1 ]
    then
	    echo 'No docker image name key word specified, please provide one';
	    exit
    fi
    keyword=$1;
    echo "The key word specified is $keyword, will delete docker images whose name contain the key word provided"
    docker rmi -f $(docker images|grep $keyword|awk '{print $3}')
}

echo "deleting all docker containers..."
rmAllContainers
# remove all docker images that built for chaincodes
echo "removing all docker images whose name contains dev"
rmImages dev
echo "removing fabric-client-kvs"
rm -rf /tmp/fabric-client-kvs_peerOrg1
echo "begin to remove all existing crypto-config"

#./network_setup.sh -s -n mychannel -f docker-compose-e2e.yaml up
./network_setup.sh  -n mychannel -t 1000 -f docker-compose-cli.yaml up
rm -rf ~/.hfc*
