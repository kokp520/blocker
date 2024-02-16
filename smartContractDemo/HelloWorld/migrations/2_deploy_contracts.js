var HelloWorld= artifacts.require ("./../contract/HelloWorld.sol");

module.exports = function(deployer) {
      deployer.deploy(HelloWorld);
}