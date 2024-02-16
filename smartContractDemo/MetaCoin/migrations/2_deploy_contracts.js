var convertLib = artifacts.require ("./../contract/ConvertLib.sol");
var metaCoin= artifacts.require ("./../contract/MetaCoin.sol");

module.exports = function(deployer) {
      deployer.deploy(convertLib)
      deployer.link(convertLib, metaCoin)
      deployer.deploy(metaCoin);
}