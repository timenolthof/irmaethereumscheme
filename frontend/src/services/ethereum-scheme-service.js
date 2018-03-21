import web3 from '../util/web3util';

import IRMASchemeABI from '../solidity/IRMASchemeABI.json';
import { IRMASchemeBytecode } from '../solidity/IRMAScheme';

export const deployNewScheme = (schemeId, metadata) => {
  const IRMASchemeContract = web3.eth.contract(IRMASchemeABI);

  return new Promise(function(resolve, reject) {
    const metadata = schemeId;
    console.log('metadata: ', metadata);
    const metadataBytes = web3.fromAscii(metadata);
    console.log(metadataBytes);
    web3.eth.estimateGas({ data: IRMASchemeBytecode }, (err, gasEstimate) => {
      console.log(gasEstimate);
      IRMASchemeContract.new(schemeId, metadataBytes, {
        from: web3.eth.coinbase,
        data: IRMASchemeBytecode,
        gas: gasEstimate*2
      }, function(err, deployedContract) {
        if (err) {
          return reject(err);
        }
        if (!deployedContract.address) {
          // Transaction is in progress
          console.log(deployedContract.transactionHash);
        } else {
          // The contract is deployed
          console.log(deployedContract.address);
          resolve(deployedContract);
        }
      });
    });
  });
};

export const fetchEthereumScheme = (schemeAddress) => {
  const IRMASchemeContract = web3.eth.contract(IRMASchemeABI);

  return new Promise(function(resolve, reject) {
    const contract = IRMASchemeContract.at(schemeAddress);
    const metadata = contract.metadata();
    resolve({
      metadata: {
        id: metadata,
      },
    });
    console.log(metadata);
    console.log('metadata:', web3.toAscii(metadata));
    // TODO: fail cases
  });
};
