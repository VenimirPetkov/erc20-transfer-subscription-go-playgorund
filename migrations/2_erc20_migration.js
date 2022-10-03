const ERC20 = artifacts.require("ERC20");

module.exports = function (deployer, network, accounts) {
  deployer.then(async () => {
    let instance = await deployer.deploy(ERC20, "Venimir", "Venko", {from: accounts[0]});
    console.log(`instance address: ${instance.address}`)
  });
};
