import web3 from '../util/web3util';

import IRMASchemeABI from '../solidity/IRMASchemeABI.json';
import { IRMASchemeBytecode } from '../solidity/IRMAScheme';

import protos from '../proto/IRMASchemeMetadata.pb.js';


const toArray = function(hex) {
    // Find termination


    if (hex.substring(0, 2) === '0x') {
        hex = hex.substr(2);
    }
    var i = 0, l = hex.length;
    console.log(l);
    const array = [];
    // var array = (buffer);
    for (; i < l; i+=2) {
        var code = parseInt(hex.substr(i, 2), 16);
        array.push(code);
    }

    return new Uint8Array(array);
};

export const deployNewScheme = (schemeId, inputMetadata) => {
  const IRMASchemeContract = web3.eth.contract(IRMASchemeABI);

  return new Promise(function(resolve, reject) {
    const metadata = {
      version: inputMetadata.version,
      id: schemeId,
      url: 'http://ethereum-scheme.url',
      name: inputMetadata.name,
      description: [{
        lang: 'en',
        name: 'An IRMA scheme hosted in an Ethereum Smart contract',
      }],
      keyshareserver: 'https://privacybydesign.foundation/tomcat/irma_keyshare_server/api/v1',
      keysharewebsite: 'https://privacybydesign.foundation/mijnirma/',
      keyshareattribute: 'pbdf.pbdf.mijnirma.email',
      contact: 'http://contacturl'
    };
    var err = protos.irmaproto.IRMASchemeMetadata.verify(metadata);
    if (err)
      console.log(err);
    // const IRMASchemeMetadataMessage = protos.irmaproto.IRMASchemeMetadata.create(metadata);
    const metadataUInt8 = protos.irmaproto.IRMASchemeMetadata.encode(metadata).finish();
    console.log(metadataUInt8);
    const metadataBytes = '0x' + Buffer.from(metadataUInt8).toString('hex');
    // const metadataStr = schemeId;
    // console.log('metadata: ', metadataStr);
    // const metadataBytes = web3.fromAscii(metadataHexStr);
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
    const metadataBytes = contract.metadata();


    console.log(metadataBytes);
    // const buffer = ByteBuffer.fromHex(metadata);
    const array = toArray(metadataBytes);
    console.log(array);
    const decoded = protos.irmaproto.IRMASchemeMetadata.decode(array);
    const metadata = protos.irmaproto.IRMASchemeMetadata.toObject(decoded);
    metadata.name = metadata.name.reduce((result, item) => {
      result[item.lang] = item.name;
      return result;
    }, {});
    metadata.description = metadata.description.reduce((result, item) => {
      result[item.lang] = item.name;
      return result;
    }, {});
    console.log();
    console.log('metadata:', metadata);
    resolve({
      metadata,
    });
    // TODO: fail cases
  });
};
