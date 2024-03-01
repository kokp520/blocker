var LuckyDraw = artifacts.require("./../contracts/LuckyDraw.sol");

module.exports = function(deployer) {
    deployer.deploy(LuckyDraw)
        .then(() => {
            console.log("deploy done.");
        })
        .catch(error => {
            console.error("deploy fail, error:", error);
        });
};