{
  dotenv: '../../../scripts/.env',
  'lyfebloc_1775-1': {
    'account-prefix': 'lyfebloc',
    'coin-type': 60,
    cmd: 'lyfeblocd',
    'start-flags': '--trace',
    'app-config': {
      'app-db-backend': 'goleveldb',
      'minimum-gas-prices': '0alyfe',
      'index-events': ['ethereum_tx.ethereumTxHash'],
      'json-rpc': {
        address: '127.0.0.1:{EVMRPC_PORT}',
        'ws-address': '127.0.0.1:{EVMRPC_PORT_WS}',
        api: 'eth,net,web3,debug',
        'feehistory-cap': 100,
        'block-range-cap': 10000,
        'logs-cap': 10000,
        'fix-revert-gas-refund-height': 1,
        enable: true,
      },
      api: {
        enable: true
      }
    },
    validators: [{
      coins: '10001000000000000000000alyfe',
      staked: '1000000000000000000alyfe',
      mnemonic: '${VALIDATOR1_MNEMONIC}',
    }, {
      coins: '10001000000000000000000alyfe',
      staked: '1000000000000000000alyfe',
      mnemonic: '${VALIDATOR2_MNEMONIC}',
    }],
    accounts: [{
      name: 'community',
      coins: '10000000000000000000000alyfe',
      mnemonic: '${COMMUNITY_MNEMONIC}',
    }, {
      name: 'signer1',
      coins: '20000000000000000000000alyfe',
      mnemonic: '${SIGNER1_MNEMONIC}',
    }, {
      name: 'signer2',
      coins: '30000000000000000000000alyfe',
      mnemonic: '${SIGNER2_MNEMONIC}',
    }],
    genesis: {
      consensus_params: {
        block: {
          max_bytes: '1048576',
          max_gas: '81500000',
        },
      },
      app_state: {
        evm: {
          params: {
            evm_denom: 'alyfe',
          },
        },
        crisis: {
          constant_fee: {
            denom: 'alyfe',
          },
        },
        staking: {
          params: {
            bond_denom: 'alyfe',
          },
        },
        inflation: {
          params: {
            mint_denom: 'alyfe',
          },
        },
        gov: {
          deposit_params: {
            max_deposit_period: '10s',
            min_deposit: [
              {
                denom: 'alyfe',
                amount: '1',
              },
            ],
          },
          params: {
            min_deposit: [
              {
                denom: 'alyfe',
                amount: '1',
              },
            ],
            max_deposit_period: '10s',
            voting_period: '10s',         
            expedited_voting_period: '5s',   
          },
        },
        transfer: {
          params: {
            receive_enabled: true,
            send_enabled: true,
          },
        },
        feemarket: {
          params: {
            no_base_fee: false,
            base_fee: '100000000000',
          },
        },
      },
    },
  },
}
