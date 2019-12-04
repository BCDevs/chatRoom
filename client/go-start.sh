echo "Removing key from key store..."

rm -rf ./hfc-key-store

# Remove chaincode docker image
sudo docker rmi -f dev-peer0.org1.example.com-mycc-1.0-384f11f484b9302df90b453200cfb25174305fce8f53f4e94d45ee3b6cab0ce9
sleep 2

cd ../basic-network
./start.sh

sudo docker ps -a



echo 'Installing chaincode..'
sudo docker exec -it cli peer chaincode install -n mycc -v 1.0 -p "/opt/gopath/src/github.com"

echo 'Instantiating chaincode..'
sudo docker exec -it cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n mycc -v 1.0 -c '{"Args":[]}'

echo 'Getting things ready for Chaincode Invocation..should take only 10 seconds..'

sleep 10

echo 'Adding Marks..'

sudo docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"addMarks","Args":["Alice","67","89","78","92"]}'
sleep 3
echo 'getting Marks..'
sudo docker exec -it cli peer chaincode query -C mychannel -n mycc -c '{"function":"getMarks","Args":["Alice"]}'

sleep 5
echo 'Deleting Marks..'
sudo docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"deleteMarks","Args":["Alice"]}'

# Starting docker logs of chaincode container

sudo docker logs -f dev-peer0.org1.example.com-mycc-1.0


