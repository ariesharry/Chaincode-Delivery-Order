
## Deploy the chaincode

Change into the test network directory.
```
cd fabric-samples/test-network
```

If the test network is already running, run the following command to bring the network down and start from a clean initial state.
```
./network.sh down
```

You can then run the following command to deploy a new network.
```
./network.sh up createChannel -ca
```

```
./network.sh deployCC -ccn delivery-order -ccp ../delivery-order/chaincode/ -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -ccl go
```

Update
./network.sh deployCC -ccn delivery-order -ccp ../delivery-order/chaincode/ -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -ccl go -ccs 2

