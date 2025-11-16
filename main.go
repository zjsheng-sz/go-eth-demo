package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/local/go-eth-demo/simpleCounter"
)

const account1 = "0x744d6ce80290a885479f031b640868b00dcb6b50"
const account2 = "0x6eb4ac69f05fef3d1507fd60349ecae33505cbf1"
const alchemyKey = "mBW0uE5HiVGi8W0Uk_0O2"
const alchemyurl = "https://eth-sepolia.g.alchemy.com/v2/"
const wsAlchemyurl = "wss://eth-sepolia.g.alchemy.com/v2/"
const myTokeAddress = "0x9C4CfecD4Ef00822F15aA31dA828E7B3eea96851"
const privateKeyStr = "61e37a97c81bd9de585b6f74d9e00a9d7050c34f628aa0246c7867dcc795f905"

func main() {

	// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
	client, err := ethclient.Dial(alchemyurl + alchemyKey)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// //查询区块
	// queryBlockInfo(client)

	// // // 发送交易
	// amount := big.NewInt(1e17)
	// sendTransaction(client, privateKeyStr, account2, amount)

	// // 查询交易
	// queryTxInfo(client, "0xd4f51ee77b9dc2824dc58c7dd94f78031301c551a0bafaf321b18e98debf25a7")

	// 部署
	// deployContract(client, privateKeyStr)

	// 加载
	// loadContract(client, "0x7133CD329D9930A1D227498CC2FEBf14a80A0203")

	// 使用合约
	useContract(client, "0x7133CD329D9930A1D227498CC2FEBf14a80A0203", privateKeyStr)

	// 订阅区块头
	// subscribeBlockHeaders()

	//查询事件
	// querylogs(client, "0x7133CD329D9930A1D227498CC2FEBf14a80A0203")

	//订阅事件
	// subscribeLogsRealtime("0x7133CD329D9930A1D227498CC2FEBf14a80A0203", nil)
}

// 查询区块信息
func queryBlockInfo(client *ethclient.Client) *types.Block {
	ctx := context.Background()

	// 获取最新区块号
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("最新区块号: %d\n", blockNumber)

	// 获取特定区块信息
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("区块时间: %d\n", block.Time())
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	return block
}

func sendTransaction(client *ethclient.Client, privateKeyStr, toAddress string, amount *big.Int) string {
	ctx := context.Background()

	// hex 转 私钥
	// 私钥 获取 公钥
	// 公钥 转 ecdsa公钥
	// ecdsa 获取 地址
	// pendingNonce nonce
	// gasprice
	// to Address
	// chainid
	// 生成交易
	// 对交易签名
	// 发送交易

	// 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress", fromAddress)

	client.PendingNonceAt(ctx, fromAddress)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 获取 gas 价格
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 构建交易
	to := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, to, amount, 21000, gasPrice, nil)

	// 获取链 ID
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal(err)
	}

	txHex := signedTx.Hash().Hex()

	fmt.Printf("交易已发送: %s\n", txHex)

	return txHex

}

func queryTxInfo(client *ethclient.Client, txStr string) {

	ctx := context.Background()

	tx, _, err := client.TransactionByHash(ctx, common.HexToHash(txStr))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx.Hash().Hex(): ", tx.Hash().Hex())
	fmt.Println("tx.Value().String(): ", tx.Value().String())
	fmt.Println("tx.Gas(): ", tx.Gas())
	fmt.Println("tx.GasPrice().Uint64(): ", tx.GasPrice().Uint64())
	fmt.Println("tx.Nonce(): ", tx.Nonce())
	fmt.Println("tx.Data(): ", tx.Data())
	fmt.Println("tx.To().Hex():", tx.To().Hex())

	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	sender, err := types.Sender(types.NewEIP155Signer(chainId), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sender: ", sender.Hex())

}

// 部署

func deployContract(client *ethclient.Client, privateKeyStr string) *simpleCounter.Counter {

	ctx := context.Background()

	privateKey, err := crypto.HexToECDSA(privateKeyStr)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0))
	auth.GasPrice = gasPrice
	// auth.GasLimit = uint64(300000)

	// 动态估算 Gas Limit
	estimatedGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     auth.From,
		To:       nil, // 合约部署时 To 为 nil
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     common.FromHex(simpleCounter.CounterBin), // 合约的字节码
	})
	if err != nil {
		log.Fatal("估算Gas失败:", err)
	}

	// 在估算值基础上增加安全边际（10-20%）
	safetyMargin := uint64(float64(estimatedGas) * 1.2)
	auth.GasLimit = safetyMargin

	fmt.Printf("估算Gas: %d, 实际设置: %d\n", estimatedGas, auth.GasLimit)

	address, tx, instance, err := simpleCounter.DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("合约部署成功")
	fmt.Println("合约地址:", address.Hex())   //0x7133CD329D9930A1D227498CC2FEBf14a80A0203
	fmt.Println("交易哈希:", tx.Hash().Hex()) //0x3ba699dfb405a59f79cd235189cd85bf6d875d21b24743e11df74d471b240215

	// 等待交易确认
	fmt.Println("等待交易确认...")
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status != 1 {
		log.Fatal("合约部署失败")
	}

	fmt.Println("合约部署成功，已确认")
	return instance
}

// 加载
func loadContract(client *ethclient.Client, contractAddressStr string) *simpleCounter.Counter {

	contractAddress := common.HexToAddress(contractAddressStr)

	simpleCounter, err := simpleCounter.NewCounter(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("合约加载成功:", contractAddressStr)

	return simpleCounter

}

func useContract(client *ethclient.Client, contractAddressStr, privateKeyStr string) {

	simpleCounter := loadContract(client, contractAddressStr)

	// 读取合约数据
	count, err := simpleCounter.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("当前计数值:", count.Uint64())

	// 修改合约数据
	ctx := context.Background()

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	auth.GasLimit = getGasLimitByOperation("simple_contract_call")

	tx, err := simpleCounter.Increment(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("计数值已增加，交易哈希:", tx.Hash().Hex())

	// 等待交易确认
	fmt.Println("等待交易确认...")
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status != 1 {
		log.Fatal("交易执行失败")
	}

	fmt.Println("交易已确认")

	// 读取更新后的合约数据
	updatedCount, err := simpleCounter.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("更新后的计数值:", updatedCount.Uint64())
}

func getGasLimitByOperation(operationType string) uint64 {
	switch operationType {
	case "contract_deploy":
		return 2000000 // 合约部署：通常需要较多 Gas
	case "simple_transfer":
		return 21000 // 简单转账：固定 21000
	case "token_transfer":
		return 65000 // ERC20 代币转账：约 45000-65000
	case "simple_contract_call":
		return 50000 // 简单合约调用
	case "complex_contract_call":
		return 200000 // 复杂合约调用
	case "swap_operation":
		return 300000 // DEX 兑换操作
	default:
		return 300000 // 默认值
	}
}

// 订阅区块头
func subscribeBlockHeaders() {
	client, err := ethclient.Dial(wsAlchemyurl + alchemyKey)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("订阅区块头成功，等待新块...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:

			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}
}

// 查询交易
func querylogs(client *ethclient.Client, contractAddressStr string) {
	contractAddress := common.HexToAddress(contractAddressStr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	query.FromBlock = header.Number
	query.ToBlock = header.Number

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println("Log Block Number: ", vLog.BlockNumber)
		fmt.Println("Log Index: ", vLog.Index)
		fmt.Println("Log Tx Hash: ", vLog.TxHash.Hex())
		fmt.Println("Log Address: ", vLog.Address.Hex())
		fmt.Println("Log Data: ", vLog.Data)
		fmt.Println("Log Topics: ", vLog.Topics)
	}
}

// 使用WebSocket实时订阅，不受批量限制
func subscribeLogsRealtime(contractAddressStr string, topics [][]common.Hash) {

	contractAddress := common.HexToAddress(contractAddressStr)

	client, err := ethclient.Dial(wsAlchemyurl + alchemyKey)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    topics,
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	query.FromBlock = header.Number
	query.ToBlock = header.Number

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("订阅失败，可能需要WebSocket连接: ", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Printf("订阅错误: %v", err)
		case log := <-logs:
			fmt.Printf("新区块 %d: 交易哈希 %s\n", log.BlockNumber, log.TxHash.Hex())

			// 创建合约过滤器实例
			filterer, err := simpleCounter.NewCounterFilterer(contractAddress, client)
			if err != nil {
			}

			// 解析日志
			counterIncrementedEvents, err := filterer.ParseCountChanged(log)
			if err != nil {
			}

			fmt.Printf("CounterIncremented 事件 - 新计数值: %d\n", counterIncrementedEvents.NewValue.Uint64())

			// 处理日志...
		}
	}
}
