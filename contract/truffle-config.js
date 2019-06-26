
var HDWalletProvider = require("truffle-hdwallet-provider");
var mnemonic = "either reduce hawk weekend excess blouse unit suspect issue direct put flat upon balance scrub";

module.exports = {
  // Uncommenting the defaults below 
  // provides for an easier quick-start with Ganache.
  // You can also follow this format for other networks;
  // see <http://truffleframework.com/docs/advanced/configuration>
  // for more details on how to specify configuration options!

  networks: {
    development: {
      host: "127.0.0.1",
      port: 8545,
      network_id: "*"
    },
    ropsten: {
      provider: function() {
        return new HDWalletProvider(mnemonic, "https://ropsten.infura.io/v3/116e0a45dcb144989126684083355eb5")
      },
      network_id: 3
    }
    // test: {
    //   host: "127.0.0.1",
    //   port: 7545,
    //   network_id: "*"
    // }
  }
};
