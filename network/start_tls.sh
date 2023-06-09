echo "一、添加环境变量"
export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH

echo "二、清理环境"
docker rm -f $(docker ps)
rm -r config
rm -r crypto-config
docker-compose down -v

echo "三、生成证书和秘钥（ MSP 材料），生成结果将保存在 crypto-config 文件夹中"
cryptogen generate --config=./crypto-config.yaml

echo "四、创建排序通道创世区块"
mkdir config
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block -channelID firstchannel

echo "五、生成通道配置事务'appchannel.tx'"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/appchannel.tx -channelID appchannel

echo "六、为各组织定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/ProducerAnchor.tx -channelID appchannel -asOrg Producer
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/ProcessorAnchor.tx -channelID appchannel -asOrg Processor
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/LogtisticsAnchor.tx -channelID appchannel -asOrg Logtistics
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/RetailerAnchor.tx -channelID appchannel -asOrg Retailer
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/RegulatorAnchor.tx -channelID appchannel -asOrg Regulator



echo "区块链 ： 启动"
docker-compose up -d
echo "正在等待节点的启动完成，等待10秒"
sleep 10

echo "配置Cli环境变量"
ProducerPeer0Cli="CORE_PEER_ADDRESS=peer0.producer.com:7051 CORE_PEER_LOCALMSPID=ProducerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/producer.com/users/Admin@producer.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/producer.com/peers/peer0.producer.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/producer.com/peers/peer0.producer.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/producer.com/peers/peer0.producer.com/tls/ca.crt"
ProducerPeer1Cli="CORE_PEER_ADDRESS=peer1.producer.com:7051 CORE_PEER_LOCALMSPID=ProducerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/producer.com/users/Admin@producer.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/producer.com/peers/peer1.producer.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/producer.com/peers/peer1.producer.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/producer.com/peers/peer1.producer.com/tls/ca.crt"
ProcessorPeer0Cli="CORE_PEER_ADDRESS=peer0.processor.com:7051 CORE_PEER_LOCALMSPID=ProcessorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/processor.com/users/Admin@processor.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/processor.com/peers/peer0.processor.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/processor.com/peers/peer0.processor.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/processor.com/peers/peer0.processor.com/tls/ca.crt"
ProcessorPeer1Cli="CORE_PEER_ADDRESS=peer1.processor.com:7051 CORE_PEER_LOCALMSPID=ProcessorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/processor.com/users/Admin@processor.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/processor.com/peers/peer1.processor.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/processor.com/peers/peer1.processor.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/processor.com/peers/peer1.processor.com/tls/ca.crt"
LogtisticsPeer0Cli="CORE_PEER_ADDRESS=peer0.logtistics.com:7051 CORE_PEER_LOCALMSPID=LogtisticsMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/logtistics.com/users/Admin@logtistics.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/logtistics.com/peers/peer0.logtistics.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/logtistics.com/peers/peer0.logtistics.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/logtistics.com/peers/peer0.logtistics.com/tls/ca.crt"
LogtisticsPeer1Cli="CORE_PEER_ADDRESS=peer1.logtistics.com:7051 CORE_PEER_LOCALMSPID=LogtisticsMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/logtistics.com/users/Admin@logtistics.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/logtistics.com/peers/peer1.logtistics.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/logtistics.com/peers/peer1.logtistics.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/logtistics.com/peers/peer1.logtistics.com/tls/ca.crt"
RetailerPeer0Cli="CORE_PEER_ADDRESS=peer0.retailer.com:7051 CORE_PEER_LOCALMSPID=RetailerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/retailer.com/users/Admin@retailer.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/retailer.com/peers/peer0.retailer.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/retailer.com/peers/peer0.retailer.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/retailer.com/peers/peer0.retailer.com/tls/ca.crt"
RetailerPeer1Cli="CORE_PEER_ADDRESS=peer1.retailer.com:7051 CORE_PEER_LOCALMSPID=RetailerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/retailer.com/users/Admin@retailer.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/retailer.com/peers/peer1.retailer.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/retailer.com/peers/peer1.retailer.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/retailer.com/peers/peer1.retailer.com/tls/ca.crt"
RegulatorPeer0Cli="CORE_PEER_ADDRESS=peer0.regulator.com:7051 CORE_PEER_LOCALMSPID=RegulatorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/regulator.com/users/Admin@regulator.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/regulator.com/peers/peer0.regulator.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/regulator.com/peers/peer0.regulator.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/regulator.com/peers/peer0.regulator.com/tls/ca.crt"
RegulatorPeer1Cli="CORE_PEER_ADDRESS=peer1.regulator.com:7051 CORE_PEER_LOCALMSPID=RegulatorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/regulator.com/users/Admin@regulator.com/msp  CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/peer/regulator.com/peers/peer1.regulator.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/peer/regulator.com/peers/peer1.regulator.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/peer/regulator.com/peers/peer1.regulator.com/tls/ca.crt"

echo "七、创建通道"
docker exec cli bash -c "$RegulatorPeer0Cli peer channel create -o orderer.regulator0.com:7050 -c appchannel -f /etc/hyperledger/config/appchannel.tx --tls true --cafile /etc/hyperledger/orderer/regulator0.com/msp/tlscacerts/tlsca.regulator0.com-cert.pem"

echo "八、将所有节点加入通道"
docker exec cli bash -c "$ProducerPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$ProducerPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$ProcessorPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$ProcessorPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$LogtisticsPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$LogtisticsPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$RetailerPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$RetailerPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$RegulatorPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$RegulatorPeer1Cli peer channel join -b appchannel.block"

echo "九、更新锚节点"
docker exec cli bash -c "$ProducerPeer0Cli peer channel update -o orderer.regulator0.com:7050 -c appchannel -f /etc/hyperledger/config/ProducerAnchor.tx"
docker exec cli bash -c "$ProcessorPeer0Cli peer channel update -o orderer.regulator0.com:7050 -c appchannel -f /etc/hyperledger/config/ProcessorAnchor.tx"
docker exec cli bash -c "$LogtisticsPeer0Cli peer channel update -o orderer.regulator0.com:7050 -c appchannel -f /etc/hyperledger/config/LogtisticsAnchor.tx"
docker exec cli bash -c "$RetailerPeer0Cli peer channel update -o orderer.regulator0.com:7050 -c appchannel -f /etc/hyperledger/config/RetailerAnchor.tx"
docker exec cli bash -c "$RegulatorPeer0Cli peer channel update -o orderer.regulator0.com:7050 -c appchannel -f /etc/hyperledger/config/RegulatorAnchor.tx"



# -n 链码名，可以自己随便设置
# -v 版本号
# -p 链码目录，在 /opt/gopath/src/ 目录下
echo "十、安装链码"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode install -n test10 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$ProcessorPeer0Cli peer chaincode install -n test10 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode install -n test10 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$RetailerPeer0Cli peer chaincode install -n test10 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$RegulatorPeer0Cli peer chaincode install -n test10 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"

# 只需要其中一个节点实例化
# -n 对应上一步安装链码的名字
# -v 版本号
# -C 是通道，在fabric的世界，一个通道就是一条不同的链
# -c 为传参，传入init参数
echo "十一、实例化链码"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode instantiate -o orderer.regulator0.com:7050 -C appchannel -n test10 -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"AND ('ProducerMSP.member','ProcessorMSP.member','LogtisticsMSP.member','RetailerMSP.member','RegulatorMSP.member')\""

echo "正在等待链码实例化完成，等待5秒"
sleep 5

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
echo "十二、验证链码"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"hello\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"addCattle\",\"FM1C1\",\"FM001\",\"FM1B1\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"deleteCattle\",\"FM1C1\",\"CSS\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"addBeff\",\"BF001\",\"FM1C1\",\"1\",\"10.5\"]}'"
docker exec cli bash -c "$ProcessorPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"frozenProcess\",\"BF001\",\"FT1OP1\",\"FT001\",\"12.0\"]}'"
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"createWaybill\",\"SF001\",\"陕A4B103\",\"ChangAn, XiAn\",\"BF001\"]}'"
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"addWaybillInfo\",\"SF001\",\"Chencang, Baoji\",\"-18.7\",\"1\"]}'"
docker exec cli bash -c "$RetailerPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"addRetailBeff\",\"BF001\",\"MK001\",\"SF001\",\"600\"]}'"
docker exec cli bash -c "$RetailerPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"addRetailBeff\",\"BF001\",\"MK001\",\"SF001\",\"600\"]}'"
docker exec cli bash -c "$RetailerPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"addSaleBill\",\"BF001\",\"MB001\"]}'"
docker exec cli bash -c "$RegualatorPeer0Cli peer chaincode invoke -C appchannel -n test10 -c '{\"Args\":[\"queryByBeffID\",\"BF001\"]}'"

