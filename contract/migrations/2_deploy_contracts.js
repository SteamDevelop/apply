// Scry Info.  All rights reserved.
// license that can be found in the license file.

let Token = artifacts.require("./Token.sol");
let Apply = artifacts.require("./Apply.sol");

let tokenContract;
module.exports = function(deployer, network, accounts) {
  deployer.deploy(Token).then(function(instance) {
    tokenContract = instance;
    console.log(tokenContract.address);

    return deployer.deploy(Apply, tokenContract.address);
  }).then (function(ptl) {
    console.log(ptl.address);
    console.log("account:", accounts.length);
  });
};
