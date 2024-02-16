# blocker
Create new chain.

## makefile

make build 建立所有bin file
make run 執行所有 會先執行build 之後看要不要clean run: clean build 
make clean =>  rm all
make test => 測試當前所有子資料夾檔案

## log 


gothereumbook : Book 
Ganache : 本地ETH環境


```go
go mod init blocker; go mod tidy;
// go get -u github.com/ethereum/go-ethereum
```



#### data 

```geth -version```

./data/ 自己建立私有練
genesis.json - geth config 
~我先用ganache試試 就好 之後再研究自己刻礦工~

### main



### other info

reply attack for ETH: ETH曾受到reply attack過

Rollup : 


### 待解決問題

P1
docker run --network host image 
本地環境開了一個repo/go-server server.go 
虛擬環境可以打到該server 但ganache開的port無法打到

後來放棄使用docker直接在本地環境中載go export切換go版本
go.work 改成 go.work.disable

ps. 後來想到
[ganache-cli](https://www.npmjs.com/package/ganache-cli)
也可以直接抓docker到虛擬環境中執行



### private key 

xyz 座標
一個橢圓曲線加密（Elliptic Curve Cryptography, ECC）中的公鑰以及私鑰。

PublicKey 是一個struct，其中包含了一對坐標 (X, Y)，這些坐標表示了橢圓曲線上的一個點，即公鑰的座標。

Curve 是指橢圓曲線的類型或參數，這裡是以十六進制表示的一個數字。
X 和 Y 則分別是該點在橢圓曲線上的 X 軸和 Y 軸的坐標。
D 則是私鑰的值，也就是橢圓曲線加密中的私鑰，同樣以十六進制表示。

--- 

block prev hash
blockchain
token
coinbase

--- 

truffle demo 


truffle unbox metacoin ./ - download demo

truffle init -- create new repo truffle env
truffle create contract -- create contrace 

truffle complie  -- complie to develop env, localhost:7545, own blockchain 

truffle migrate -- node js 
truffle console -- to console 

truffle console :
# get instance 


truffle(development)> let instance = await MetaCoin.deployed()
truffle(development)> let accounts = await web3.eth.getAccounts()

truffle(development)> let balance = await instance.getBalance(accounts[0])
truffle(development)> balance.toNumber()

# eth

truffle(development)> let ether = await instance.getBalanceInEth(accounts[0])
truffle(development)> ether.toNumber()


# sendcoin

truffle(development)> instance.sendCoin(accounts[1], 500)


truffle(development)> let received = await instance.getBalance(accounts[1])
truffle(development)> received.toNumber()


truffle(development)> let newBalance = await instance.getBalance(accounts[0])
truffle(development)> newBalance.toNumber()

--- 



truffle config要調整到我的blockchain

ganache 就是一個鏈 address 
remix 會自己建立一個vm有自己的鏈
truffle sol框架 
lafura在線上環境上的執行dapp的環境

--- 

todo question

- 部署多個合約到同個區塊鏈上會怎樣
ans: 名稱相通但地址不同 只是會造成使用者混淆

- 有些合約不用gas?

目前觀察
要對合約上的變數有異動才會有gas fee
call and transacation的差別

- remix為什麼有版本限制?

remix sol 8.20之後有個func只能夠透過以太call

- 當我在truffle 執行合約 有進行proof of work嗎

不會，沒有證明 truffle 預設拿第一個帳戶當作gas消耗 也就是執行者
但沒有礦工 礦工就是truffle本人
