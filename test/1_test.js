const ERC20 = artifacts.require("ERC20");
contract("Access control tests", (accounts) => {
    let ERC20Instance;
    let amount = new web3.utils.BN(web3.utils.toWei('1', 'ether'));
    let firstHolder = accounts[0];

    beforeEach(async () => {
        ERC20Instance = await ERC20.deployed();
    });

    describe("Transfer", async () => {
        it("Transfer tokens", async () => {
            for(let i = 1; i < accounts.length; i++){
                let tx = await ERC20Instance.transfer(accounts[i], amount, {from: firstHolder});
                console.log(`${i}, tx: ${tx.tx}`)
            }
        });
    });
});
      