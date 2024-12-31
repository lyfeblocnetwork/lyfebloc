// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)
package types

// erc20Bytecode got from a query previous to the upgrade
// eth.getCode("0xc7e56EEc629D3728fE41baCa2f6BFc502096f94E","0x15B5788")
// here were're querying height 22763400 and
// 0xc7e56EEc629D3728fE41baCa2f6BFc502096f94E is the erc20 representation an IBC coin (stSTARS)
const Erc20Bytecode = "0x608060405234801561001057600080fd5b50600436106101da5760003560e01c80635c975abb11610104578063a217fddf116100a2578063d539139311610071578063d53913931461057d578063d547741f1461059b578063dd62ed3e146105b7578063e63ab1e9146105e7576101da565b8063a217fddf146104cf578063a457c2d7146104ed578063a9059cbb1461051d578063ca15c8731461054d576101da565b80638456cb59116100de5780638456cb59146104475780639010d07c1461045157806391d148541461048157806395d89b41146104b1576101da565b80635c975abb146103dd57806370a08231146103fb57806379cc67901461042b576101da565b8063282c51f31161017c578063395093511161014b578063395093511461036b5780633f4ba83a1461039b57806340c10f19146103a557806342966c68146103c1576101da565b8063282c51f3146102f75780632f2ff15d14610315578063313ce5671461033157806336568abe1461034f576101da565b806318160ddd116101b857806318160ddd1461025d5780631cf2c7e21461027b57806323b872dd14610297578063248a9ca3146102c7576101da565b806301ffc9a7146101df57806306fdde031461020f578063095ea7b31461022d575b600080fd5b6101f960048036038101906101f49190612216565b610605565b604051610206919061225e565b60405180910390f35b61021761067f565b6040516102249190612312565b60405180910390f35b610247600480360381019061024291906123c8565b610711565b604051610254919061225e565b60405180910390f35b61026561072f565b6040516102729190612417565b60405180910390f35b610295600480360381019061029091906123c8565b610739565b005b6102b160048036038101906102ac9190612432565b6107b7565b6040516102be919061225e565b60405180910390f35b6102e160048036038101906102dc91906124bb565b6108af565b6040516102ee91906124f7565b60405180910390f35b6102ff6108ce565b60405161030c91906124f7565b60405180910390f35b61032f600480360381019061032a9190612512565b6108f2565b005b61033961091b565b604051610346919061256e565b60405180910390f35b61036960048036038101906103649190612512565b610932565b005b610385600480360381019061038091906123c8565b6109b5565b604051610392919061225e565b60405180910390f35b6103a3610a61565b005b6103bf60048036038101906103ba91906123c8565b610adb565b005b6103db60048036038101906103d69190612589565b610b59565b005b6103e5610b6d565b6040516103f2919061225e565b60405180910390f35b610415600480360381019061041091906125b6565b610b84565b6040516104229190612417565b60405180910390f35b610445600480360381019061044091906123c8565b610bcd565b005b61044f610c48565b005b61046b600480360381019061046691906125e3565b610cc2565b6040516104789190612632565b60405180910390f35b61049b60048036038101906104969190612512565b610cf1565b6040516104a8919061225e565b60405180910390f35b6104b9610d5b565b6040516104c69190612312565b60405180910390f35b6104d7610ded565b6040516104e491906124f7565b60405180910390f35b610507600480360381019061050291906123c8565b610df4565b604051610514919061225e565b60405180910390f35b610537600480360381019061053291906123c8565b610edf565b604051610544919061225e565b60405180910390f35b610567600480360381019061056291906124bb565b610efd565b6040516105749190612417565b60405180910390f35b610585610f21565b60405161059291906124f7565b60405180910390f35b6105b560048036038101906105b09190612512565b610f45565b005b6105d160048036038101906105cc919061264d565b610f6e565b6040516105de9190612417565b60405180910390f35b6105ef610ff5565b6040516105fc91906124f7565b60405180910390f35b60007f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610678575061067782611129565b5b9050919050565b60606005805461068e906126bc565b80601f01602080910402602001604051908101604052809291908181526020018280546106ba906126bc565b80156107075780601f106106dc57610100808354040283529160200191610707565b820191906000526020600020905b8154815290600101906020018083116106ea57829003601f168201915b5050505050905090565b600061072561071e6111a3565b84846111ab565b6001905092915050565b6000600454905090565b61076a7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8486107656111a3565b610cf1565b6107a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a090612760565b60405180910390fd5b6107b38282611376565b5050565b60006107c484848461154f565b6000600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061080f6111a3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561088f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610886906127f2565b60405180910390fd5b6108a38561089b6111a3565b8584036111ab565b60019150509392505050565b6000806000838152602001908152602001600020600101549050919050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b6108fb826108af565b61090c816109076111a3565b6117d3565b6109168383611870565b505050565b6000600760019054906101000a900460ff16905090565b61093a6111a3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146109a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099e90612884565b60405180910390fd5b6109b182826118a4565b5050565b6000610a576109c26111a3565b8484600360006109d06111a3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610a5291906128d3565b6111ab565b6001905092915050565b610a927f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a8d6111a3565b610cf1565b610ad1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac89061299b565b60405180910390fd5b610ad96118d8565b565b610b0c7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6610b076111a3565b610cf1565b610b4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4290612a2d565b60405180910390fd5b610b55828261197a565b5050565b610b6a610b646111a3565b82611376565b50565b6000600760009054906101000a900460ff16905090565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000610be083610bdb6111a3565b610f6e565b905081811015610c25576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1c90612abf565b60405180910390fd5b610c3983610c316111a3565b8484036111ab565b610c438383611376565b505050565b610c797f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610c746111a3565b610cf1565b610cb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610caf90612b51565b60405180910390fd5b610cc0611adb565b565b6000610ce98260016000868152602001908152602001600020611b7e90919063ffffffff16565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b606060068054610d6a906126bc565b80601f0160208091040260200160405190810160405280929190818152602001828054610d96906126bc565b8015610de35780601f10610db857610100808354040283529160200191610de3565b820191906000526020600020905b815481529060010190602001808311610dc657829003601f168201915b5050505050905090565b6000801b81565b60008060036000610e036111a3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610ec0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb790612be3565b60405180910390fd5b610ed4610ecb6111a3565b858584036111ab565b600191505092915050565b6000610ef3610eec6111a3565b848461154f565b6001905092915050565b6000610f1a60016000848152602001908152602001600020611b98565b9050919050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b610f4e826108af565b610f5f81610f5a6111a3565b6117d3565b610f6983836118a4565b505050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6110238282610cf1565b6110f557600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061109a6111a3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000611121836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611bad565b905092915050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061119c575061119b82611c1d565b5b9050919050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561121b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121290612c75565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561128b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161128290612d07565b60405180910390fd5b80600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516113699190612417565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156113e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113dd90612d99565b60405180910390fd5b6113f282600083611c87565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611479576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147090612e2b565b60405180910390fd5b818103600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008282546114d19190612e4b565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516115369190612417565b60405180910390a361154a83600084611c97565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156115bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115b690612ef1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561162f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162690612f83565b60405180910390fd5b61163a838383611c87565b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156116c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116b890613015565b60405180910390fd5b818103600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461175691906128d3565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516117ba9190612417565b60405180910390a36117cd848484611c97565b50505050565b6117dd8282610cf1565b61186c576118028173ffffffffffffffffffffffffffffffffffffffff166014611c9c565b6118108360001c6020611c9c565b604051602001611821929190613109565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118639190612312565b60405180910390fd5b5050565b61187a8282611019565b61189f81600160008581526020019081526020016000206110f990919063ffffffff16565b505050565b6118ae8282611ed8565b6118d38160016000858152602001908152602001600020611fb990919063ffffffff16565b505050565b6118e0610b6d565b61191f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119169061318f565b60405180910390fd5b6000600760006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6119636111a3565b6040516119709190612632565b60405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156119ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119e1906131fb565b60405180910390fd5b6119f660008383611c87565b8060046000828254611a0891906128d3565b9250508190555080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611a5e91906128d3565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051611ac39190612417565b60405180910390a3611ad760008383611c97565b5050565b611ae3610b6d565b15611b23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b1a90613267565b60405180910390fd5b6001600760006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611b676111a3565b604051611b749190612632565b60405180910390a1565b6000611b8d8360000183611fe9565b60001c905092915050565b6000611ba682600001612014565b9050919050565b6000611bb98383612025565b611c12578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050611c17565b600090505b92915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b611c92838383612048565b505050565b505050565b606060006002836002611caf9190613287565b611cb991906128d3565b67ffffffffffffffff811115611cd257611cd16132e1565b5b6040519080825280601f01601f191660200182016040528015611d045781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110611d3c57611d3b613310565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110611da057611d9f613310565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002611de09190613287565b611dea91906128d3565b90505b6001811115611e8a577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110611e2c57611e2b613310565b5b1a60f81b828281518110611e4357611e42613310565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080611e839061333f565b9050611ded565b5060008414611ece576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ec5906133b5565b60405180910390fd5b8091505092915050565b611ee28282610cf1565b15611fb557600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550611f5a6111a3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000611fe1836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6120a0565b905092915050565b600082600001828154811061200157612000613310565b5b9060005260206000200154905092915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b6120538383836121b4565b61205b610b6d565b1561209b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161209290613447565b60405180910390fd5b505050565b600080836001016000848152602001908152602001600020549050600081146121a85760006001826120d29190612e4b565b90506000600186600001805490506120ea9190612e4b565b905081811461215957600086600001828154811061210b5761210a613310565b5b906000526020600020015490508087600001848154811061212f5761212e613310565b5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b8560000180548061216d5761216c613467565b5b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506121ae565b60009150505b92915050565b505050565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6121f3816121be565b81146121fe57600080fd5b50565b600081359050612210816121ea565b92915050565b60006020828403121561222c5761222b6121b9565b5b600061223a84828501612201565b91505092915050565b60008115159050919050565b61225881612243565b82525050565b6000602082019050612273600083018461224f565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156122b3578082015181840152602081019050612298565b838111156122c2576000848401525b50505050565b6000601f19601f8301169050919050565b60006122e482612279565b6122ee8185612284565b93506122fe818560208601612295565b612307816122c8565b840191505092915050565b6000602082019050818103600083015261232c81846122d9565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061235f82612334565b9050919050565b61236f81612354565b811461237a57600080fd5b50565b60008135905061238c81612366565b92915050565b6000819050919050565b6123a581612392565b81146123b057600080fd5b50565b6000813590506123c28161239c565b92915050565b600080604083850312156123df576123de6121b9565b5b60006123ed8582860161237d565b92505060206123fe858286016123b3565b9150509250929050565b61241181612392565b82525050565b600060208201905061242c6000830184612408565b92915050565b60008060006060848603121561244b5761244a6121b9565b5b60006124598682870161237d565b935050602061246a8682870161237d565b925050604061247b868287016123b3565b9150509250925092565b6000819050919050565b61249881612485565b81146124a357600080fd5b50565b6000813590506124b58161248f565b92915050565b6000602082840312156124d1576124d06121b9565b5b60006124df848285016124a6565b91505092915050565b6124f181612485565b82525050565b600060208201905061250c60008301846124e8565b92915050565b60008060408385031215612529576125286121b9565b5b6000612537858286016124a6565b92505060206125488582860161237d565b9150509250929050565b600060ff82169050919050565b61256881612552565b82525050565b6000602082019050612583600083018461255f565b92915050565b60006020828403121561259f5761259e6121b9565b5b60006125ad848285016123b3565b91505092915050565b6000602082840312156125cc576125cb6121b9565b5b60006125da8482850161237d565b91505092915050565b600080604083850312156125fa576125f96121b9565b5b6000612608858286016124a6565b9250506020612619858286016123b3565b9150509250929050565b61262c81612354565b82525050565b60006020820190506126476000830184612623565b92915050565b60008060408385031215612664576126636121b9565b5b60006126728582860161237d565b92505060206126838582860161237d565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806126d457607f821691505b602082108114156126e8576126e761268d565b5b50919050565b7f45524332304d696e7465724275726e6572446563696d616c733a206d7573742060008201527f68617665206275726e657220726f6c6520746f206275726e0000000000000000602082015250565b600061274a603883612284565b9150612755826126ee565b604082019050919050565b600060208201905081810360008301526127798161273d565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b60006127dc602883612284565b91506127e782612780565b604082019050919050565b6000602082019050818103600083015261280b816127cf565b9050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b600061286e602f83612284565b915061287982612812565b604082019050919050565b6000602082019050818103600083015261289d81612861565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006128de82612392565b91506128e983612392565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561291e5761291d6128a4565b5b828201905092915050565b7f45524332304d696e7465724275726e6572446563696d616c733a206d7573742060008201527f686176652070617573657220726f6c6520746f20756e70617573650000000000602082015250565b6000612985603b83612284565b915061299082612929565b604082019050919050565b600060208201905081810360008301526129b481612978565b9050919050565b7f45524332304d696e7465724275726e6572446563696d616c733a206d7573742060008201527f68617665206d696e74657220726f6c6520746f206d696e740000000000000000602082015250565b6000612a17603883612284565b9150612a22826129bb565b604082019050919050565b60006020820190508181036000830152612a4681612a0a565b9050919050565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f7760008201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b6000612aa9602483612284565b9150612ab482612a4d565b604082019050919050565b60006020820190508181036000830152612ad881612a9c565b9050919050565b7f45524332304d696e7465724275726e6572446563696d616c733a206d7573742060008201527f686176652070617573657220726f6c6520746f20706175736500000000000000602082015250565b6000612b3b603983612284565b9150612b4682612adf565b604082019050919050565b60006020820190508181036000830152612b6a81612b2e565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000612bcd602583612284565b9150612bd882612b71565b604082019050919050565b60006020820190508181036000830152612bfc81612bc0565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000612c5f602483612284565b9150612c6a82612c03565b604082019050919050565b60006020820190508181036000830152612c8e81612c52565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000612cf1602283612284565b9150612cfc82612c95565b604082019050919050565b60006020820190508181036000830152612d2081612ce4565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000612d83602183612284565b9150612d8e82612d27565b604082019050919050565b60006020820190508181036000830152612db281612d76565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000612e15602283612284565b9150612e2082612db9565b604082019050919050565b60006020820190508181036000830152612e4481612e08565b9050919050565b6000612e5682612392565b9150612e6183612392565b925082821015612e7457612e736128a4565b5b828203905092915050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000612edb602583612284565b9150612ee682612e7f565b604082019050919050565b60006020820190508181036000830152612f0a81612ece565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000612f6d602383612284565b9150612f7882612f11565b604082019050919050565b60006020820190508181036000830152612f9c81612f60565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000612fff602683612284565b915061300a82612fa3565b604082019050919050565b6000602082019050818103600083015261302e81612ff2565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000613076601783613035565b915061308182613040565b601782019050919050565b600061309782612279565b6130a18185613035565b93506130b1818560208601612295565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b60006130f3601183613035565b91506130fe826130bd565b601182019050919050565b600061311482613069565b9150613120828561308c565b915061312b826130e6565b9150613137828461308c565b91508190509392505050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b6000613179601483612284565b915061318482613143565b602082019050919050565b600060208201905081810360008301526131a88161316c565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006131e5601f83612284565b91506131f0826131af565b602082019050919050565b60006020820190508181036000830152613214816131d8565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b6000613251601083612284565b915061325c8261321b565b602082019050919050565b6000602082019050818103600083015261328081613244565b9050919050565b600061329282612392565b915061329d83612392565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156132d6576132d56128a4565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061334a82612392565b9150600082141561335e5761335d6128a4565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b600061339f602083612284565b91506133aa82613369565b602082019050919050565b600060208201905081810360008301526133ce81613392565b9050919050565b7f45524332305061757361626c653a20746f6b656e207472616e7366657220776860008201527f696c652070617573656400000000000000000000000000000000000000000000602082015250565b6000613431602a83612284565b915061343c826133d5565b604082019050919050565b6000602082019050818103600083015261346081613424565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220c3d4a4231a6c94cfb03623ea4b77df2c9ccfa487132bebf43620219e3dc2f4cf64736f6c63430008090033"