import 'babel-polyfill'
import React, { Component } from 'react';
import Web3 from 'web3';
// import NetworkStatus from 'react-web3-network-status';

import GrommetApp from 'grommet/components/App';
import {
  Form, Header, FormFields, Footer, Button, Heading, TextInput
} from 'grommet';

import IRMASchemeABI from '../solidity/IRMASchemeABI.json';
import { IRMASchemeBytecode } from '../solidity/IRMAScheme';


class NewScheme extends Component {

  constructor (props) {
    super(props);
    if (typeof window.web3 !== 'undefined') {
      // this.web3 = window.web3;
      this.web3 = new Web3(window.web3.currentProvider);
    } else {
      this.web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));
    }
    console.log(this.web3.version);
    console.log(IRMASchemeABI);

    // console.log(compiledContract);
    this.state = {
      abi: IRMASchemeABI,
      contract: this.web3.eth.contract(IRMASchemeABI),
      coinbase: this.web3.eth.coinbase,
      fSchemeId: 'new-ethereum-scheme',
      // web3: null,
      // web3loaded: false,
      // lucky: 'n/a',
      // transaction: 'n/a'
    };
    console.log(this.state.abi);

  }

  // componentWillMount() {
  //   window.addEventListener('load', this.initWeb3)
  // }
  //
  // componentWillUnmount() {
  //   window.removeEventListener('load', this.initWeb3)
  // }
  //
  // initWeb3() {
  //   console.log("what is web3 right now: " + window.web3);
  //
  //   if (typeof web3 !== 'undefined') {
  //     this.web3Provider = window.web3.currentProvider;
  //     window.web3 = new Web3(window.web3.currentProvider);
  //   } else {
  //     this.web3Provider = new Web3.providers.HttpProvider('http://localhost:8545');
  //     window.web3 = new Web3(this.web3Provider);
  //   }
  // }

  // componentDidMount = async () => {
  //   // const web3 = await getWeb3Async()
  //   // if(web3.isConnected()) {
  //   //   // const abi = await web3.eth.contract(ABIInterfaceArray)
  //   //   // const instance = instancePromisifier(abi.at(SMART_CONTRACT_INSTANCE))
  //   //
  //   //   // console.log('Interface', ABIInterfaceArray)
  //   //   this.setState({ web3: web3, web3loaded: true })
  //   // }
  //   // initWeb3();
  //   // Get some props
  //   // const coinbase = window.web3.eth.coinbase
  //   // const balance = window.web3.eth.getBalance(coinbase).toString()
  //   //
  //   // // Watch for change
  //   // window.web3.eth.filter('latest').watch(() => {
  //   //   const balance = window.web3.eth.getBalance(coinbase).toString()
  //   //   this.setState({ balance })
  //   // })
  //   //
  //   // // Send some ether every second
  //   // // setInterval(this.doAirDrop, 1000)
  //   //
  //   // this.setState({ loading: false, coinbase, balance })
  // }
  //
  // state = {
  //   coinbase: null
  // };
  // async componentWillMount () {
  //   const { web3 } = this.props;
  //   try {
  //     const coinbase = await web3.eth.getCoinbase();
  //     this.setState({ coinbase });
  //   } catch (e) {
  //     this.setState({
  //       coinbase: `Cannot get coinbase from ${web3.currentProvider.host}`
  //     });
  //   }
  // }

  deployNewScheme() {
    // console.log(this.web3.eth.compile.solidity(IRMAScheme));
    this.web3.eth.estimateGas({ data: IRMASchemeBytecode }, (err, gasEstimate) => {
      console.log(gasEstimate);
      this.state.contract.new(this.state.fSchemeId, ["20160528"], {
        from: this.state.coinbase,
        data: IRMASchemeBytecode,
        gas: gasEstimate*2
      }, function(err, myContract) {
        if(!err) {
           // NOTE: The callback will fire twice!
           // Once the contract has the transactionHash property set and once its deployed on an address.

           // e.g. check tx hash on the first call (transaction send)
           if(!myContract.address) {
               console.log(myContract.transactionHash) // The hash of the transaction, which deploys the contract

           // check address on the second call (contract deployed)
           } else {
               console.log(myContract.address) // the contract address
           }

           // Note that the returned "myContractReturned" === "myContract",
           // so the returned "myContractReturned" object will also get the address set.
        }
      });
    });
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  handleSubmit(event) {
    alert('A name was submitted: ' + this.state.fSchemeId);
    //event.preventDefault();
    this.deployNewScheme();
  }

  render() {
    console.log(this.web3.eth.accounts);
    return (
      <GrommetApp>
        Coinbase: { this.state.coinbase }
        Balance: { this.state.balance }
        <Form pad="medium" >
          <Header>
            <Heading>
              New Scheme
            </Heading>
          </Header>
          <FormFields>
          Scheme Id:<br/>
          <TextInput id='schemeId'
            placeHolder="scheme id"
            value={this.state.fSchemeId} onChange={this.handleInputChange}
            name='fSchemeId'/>
          </FormFields>
          <Footer pad={{"vertical": "medium"}}>
            <Button label='Submit'
              primary={true}
              onClick={() => this.handleSubmit()} />
          </Footer>
        </Form>
      </GrommetApp>
    );
  }
}

export default NewScheme;
