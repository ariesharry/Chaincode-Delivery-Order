const express = require('express');
const router = express.Router();

const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const { buildCCPOrg1, buildCCPOrg2, buildWallet, prettyJSONString } = require('../../../../test-application/javascript/AppUtil.js');
const myChannel = 'mychannel';
const myChaincodeName = 'delivery';

async function createRequestDO (ccp, wallet, user) {
	try {
		const gateway = new Gateway();
		// connect using Discovery enabled

		await gateway.connect(ccp,
			{ wallet: wallet, identity: user, discovery: { enabled: true, asLocalhost: true } });

		const network = await gateway.getNetwork(myChannel);
		const contract = network.getContract(myChaincodeName);

        // Submit a requestDO transaction
        // const requestor = {
        //     requestorType: requestorType,
        //     urlFile: urlFile,
        //     npwp: npwp,
        //     nib: nib,
        //     requestorName: requestorName,
        //     requestorAddress: requestorAddress,
        // };

		// const statefulTxn = contract.createTransaction('requestDO');

		// console.log('\n--> Submit Transaction: Propose a new auction');
		// await statefulTxn.submit(JSON.stringify({ requestor }));
		// console.log('*** Result: committed');

		console.log('\n--> Evaluate Transaction: query the auction that was just created');
		const result = await contract.evaluateTransaction('queryDO');
		console.log('*** Result: Auction: ' + prettyJSONString(result.toString()));
        const requestInfo = prettyJSONString(result.toString());
		return requestInfo;

		gateway.disconnect();
	} catch (error) {
		console.error(`******** FAILED to submit bid: ${error}`);
		throw error;
	}
}


router.post('/', async function (req, res, next) {
    const requestDetail ={
        org: req.body.org,
        user: req.body.user
    }

    var requestInfo;

    try {
        const org = requestDetail.org;
        const user = requestDetail.user;

        if (org === "Org1" || org === "org1") {
            const ccp = buildCCPOrg1();
            const walletPath = path.join(__dirname, 'wallet/org1');
            const wallet = await buildWallet(Wallets, walletPath);
            requestInfo = await createRequestDO(ccp, wallet, user);
            requestInfo = JSON.parse(requestInfo);
        } else if (org === 'Org2' || org === 'org2') {
            const ccp = buildCCPOrg2();
            const walletPath = path.join(__dirname, 'wallet/org2');
            const wallet = await buildWallet(Wallets, walletPath);
            requestInfo = await createRequestDO(ccp, wallet, user);
            requestInfo = JSON.parse(requestInfo);
        } else {
            console.log('Usage: node createAuction.js org userID auctionID item quantity');
            console.log('Org must be Org1 or Org2');
        }
    } catch (error) {
		console.error(`******** FAILED to run the application: ${error}`);
	}

    res.status(200).json({
        message: 'successfully query request DO',
        requestDetail: requestInfo
    });
});

module.exports = router;