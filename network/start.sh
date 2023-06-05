echo "一、添加环境变量"
export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH

echo "二、清理环境"
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
ProducerPeer0Cli="CORE_PEER_ADDRESS=peer0.producer.com:7051 CORE_PEER_LOCALMSPID=ProducerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/producer.com/users/Admin@producer.com/msp"
ProducerPeer1Cli="CORE_PEER_ADDRESS=peer1.producer.com:7051 CORE_PEER_LOCALMSPID=ProducerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/producer.com/users/Admin@producer.com/msp"
ProcessorPeer0Cli="CORE_PEER_ADDRESS=peer0.processor.com:7051 CORE_PEER_LOCALMSPID=ProcessorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/processor.com/users/Admin@processor.com/msp"
ProcessorPeer1Cli="CORE_PEER_ADDRESS=peer1.processor.com:7051 CORE_PEER_LOCALMSPID=ProcessorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/processor.com/users/Admin@processor.com/msp"
LogtisticsPeer0Cli="CORE_PEER_ADDRESS=peer0.logtistics.com:7051 CORE_PEER_LOCALMSPID=LogtisticsMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/logtistics.com/users/Admin@logtistics.com/msp"
LogtisticsPeer1Cli="CORE_PEER_ADDRESS=peer1.logtistics.com:7051 CORE_PEER_LOCALMSPID=LogtisticsMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/logtistics.com/users/Admin@logtistics.com/msp"
RetailerPeer0Cli="CORE_PEER_ADDRESS=peer0.retailer.com:7051 CORE_PEER_LOCALMSPID=RetailerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/retailer.com/users/Admin@retailer.com/msp"
RetailerPeer1Cli="CORE_PEER_ADDRESS=peer1.retailer.com:7051 CORE_PEER_LOCALMSPID=RetailerMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/retailer.com/users/Admin@retailer.com/msp"
RegulatorPeer0Cli="CORE_PEER_ADDRESS=peer0.regulator.com:7051 CORE_PEER_LOCALMSPID=RegulatorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/regulator.com/users/Admin@regulator.com/msp"
RegulatorPeer1Cli="CORE_PEER_ADDRESS=peer1.regulator.com:7051 CORE_PEER_LOCALMSPID=RegulatorMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/regulator.com/users/Admin@regulator.com/msp"

echo "七、创建通道"
docker exec cli bash -c "$RegulatorPeer0Cli peer channel create -o orderer.regulator0.com:7050 -c appchannel -f /etc/hyperledger/config/appchannel.tx"

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
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode install -n test13 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$ProcessorPeer0Cli peer chaincode install -n test13 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode install -n test13 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$RetailerPeer0Cli peer chaincode install -n test13 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"
docker exec cli bash -c "$RegulatorPeer0Cli peer chaincode install -n test13 -v 1.0.0 -l golang -p github.com/Lingxing-GT/fabric-coldtrain/chaincode"

# 只需要其中一个节点实例化
# -n 对应上一步安装链码的名字
# -v 版本号
# -C 是通道，在fabric的世界，一个通道就是一条不同的链
# -c 为传参，传入init参数
echo "十一、实例化链码"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode instantiate -o orderer.regulator0.com:7050 -C appchannel -n test13 -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"OR ('ProducerMSP.member','ProcessorMSP.member','LogtisticsMSP.member','RetailerMSP.member','RegulatorMSP.member')\""

echo "正在等待链码实例化完成，等待5秒"
sleep 5

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
echo "十二、验证链码"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"hello\"]}'"

