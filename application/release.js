const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

async function releaseDeliveryOrder() {
  try {
    // Load the network configuration and user credentials
    const ccpPath = path.resolve(__dirname, 'path/to/connection-profile.json');
    const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
    const ccp = JSON.parse(ccpJSON);
    const walletPath = path.resolve(__dirname, 'path/to/wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    const identity = 'user1'; // Specify the identity to use

    // Create a new gateway for connecting to the network
    const gateway = new Gateway();
    await gateway.connect(ccp, {
      wallet,
      identity,
      discovery: { enabled: true, asLocalhost: true },
    });

    // Get the network and contract
    const network = await gateway.getNetwork('mychannel');
    const contract = network.getContract('mycc');

    // Release the delivery order
    const deliveryOrderID = '12345'; // Specify the delivery order ID to release

    // Submit the release transaction
    await contract.submitTransaction('ReleaseDeliveryOrder', deliveryOrderID);
    console.log('Release transaction has been submitted successfully.');

    // Disconnect from the gateway
    gateway.disconnect();
  } catch (error) {
    console.error(`Failed to submit release transaction: ${error}`);
    process.exit(1);
  }
}

releaseDeliveryOrder();
