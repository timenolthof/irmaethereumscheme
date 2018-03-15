import web3 from '../util/web3util';

import IRMASchemeABI from '../solidity/IRMASchemeABI.json';
import { IRMASchemeBytecode } from '../solidity/IRMAScheme';

export const deployNewScheme = (schemeId, metadata) => {
  const IRMASchemeContract = web3.eth.contract(IRMASchemeABI);

  return new Promise(function(resolve, reject) {
    console.log(schemeId, metadata);
    const str = web3.toHex('aomedata');
    const metadataBytes = [str];
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
  console.log(schemeAddress);
  const IRMASchemeContract = web3.eth.contract(IRMASchemeABI);
  const contract = IRMASchemeContract.at(schemeAddress);

  return new Promise(function(resolve, reject) {
    const metadata = contract.metadata(null);
    console.log('id', web3.toAscii(metadata));

  });
};
