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
#------------------addCattle-------------------#
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addCattle\",\"FM1C1\",\"FM001\",\"FM1B1\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addCattle\",\"FM1C2\",\"FM001\",\"FM1B1\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addCattle\",\"FM2C1\",\"FM002\",\"FM2B1\"]}'"
#---------------addCattleGrowInfo--------------#
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addCattleGrowInfo\",\"FM1C1\",\"36.5\",\"Good\",\"317.8\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addCattleGrowInfo\",\"FM1C1\",\"37.0\",\"Good\",\"325.4\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addCattleGrowInfo\",\"FM1C1\",\"36.3\",\"Good\",\"330.1\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addCattleGrowInfo\",\"FM1C1\",\"37.2\",\"Good\",\"337.1\",\"即将出栏\"]}'"
#-----------------deleteCattle-----------------#
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"deleteCattle\",\"FM1C1\",\"CSS\"]}'"
#--------------------addBeff-------------------#
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addBeff\",\"BF001\",\"FM1C1\",\"1\",\"10.5\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addBeff\",\"BF002\",\"FM1C1\",\"1\",\"20.8\"]}'"
docker exec cli bash -c "$ProducerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addBeff\",\"BF003\",\"FM1C1\",\"2\",\"37.4\"]}'"
#-----------------frozenProcess-----------------#
docker exec cli bash -c "$ProcessorPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"frozenProcess\",\"BF001\",\"FT1OP1\",\"FT001\",\"12.0\"]}'"
docker exec cli bash -c "$ProcessorPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"frozenProcess\",\"BF002\",\"FT2OP1\",\"FT002\",\"22.8\"]}'"
docker exec cli bash -c "$ProcessorPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"frozenProcess\",\"BF003\",\"FT2OP1\",\"FT002\",\"38.2\"]}'"
#-------------------createWaybill---------------#
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"createWaybill\",\"SF001\",\"陕A4B103\",\"ChangAn, XiAn\",\"BF001\"]}'"
#-------------------addWaybillInfo--------------#
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addWaybillInfo\",\"SF001\",\"Guancheng, Zhengzhou\",\"-18.7\"]}'"
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addWaybillInfo\",\"SF001\",\"Jiangxia, Wuhan\",\"-20.3\"]}'"
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addWaybillInfo\",\"SF001\",\"Xinjian, Nanshan\",\"-19.5\"]}'"
docker exec cli bash -c "$LogtisticsPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addWaybillInfo\",\"SF001\",\"Xinfeng, Ganzhou\",\"-23.4\",\"1\"]}'"
#-------------------addRetailBeff--------------#
docker exec cli bash -c "$RetailerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addRetailBeff\",\"BF001\",\"MK001\",\"SF001\",\"600\"]}'"
#-------------------addSaleBill----------------#
docker exec cli bash -c "$RetailerPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"addSaleBill\",\"BF001\",\"MB001\"]}'"
#-------------------queryByBeffID----------------#
docker exec cli bash -c "$RegulatorPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"queryByBeffID\",\"BF001\"]}'"
#---------------- -queryByWaybillNo--------------#
docker exec cli bash -c "$RegulatorPeer0Cli peer chaincode invoke -C appchannel -n test13 -c '{\"Args\":[\"queryByWaybillNo\",\"SF001\"]}'"
