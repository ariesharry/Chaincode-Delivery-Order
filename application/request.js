const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

async function requestDeliveryOrder() {
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

    // Make a request for a delivery order
    const requestDetails = {
      // Request details object
    };

    // Submit the request transaction
    await contract.submitTransaction('RequestDeliveryOrder', JSON.stringify(requestDetails));
    console.log('Request transaction has been submitted successfully.');

    // Disconnect from the gateway
    gateway.disconnect();
  } catch (error) {
    console.error(`Failed to submit request transaction: ${error}`);
    process.exit(1);
  }
}

requestDeliveryOrder();
