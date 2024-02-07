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


