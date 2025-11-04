package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接到以太坊节点（可以使用 Infura 或本地节点）
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/mBW0uE5HiVGi8W0Uk_0O2")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	queryBlockHeader(client)

}

func queryBlockHeader(client *ethclient.Client) {
	ctx := context.Background()
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功连接到以太坊网络")
	fmt.Printf("最新区块号: %s\n", header.Number.String())

	fmt.Println(header.Number.String())
}

// 查询区块信息
func queryBlockInfo(client *ethclient.Client) {
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
}

// 查询账户余额
func queryBalance(client *ethclient.Client, address string) {
	ctx := context.Background()

	// 将地址字符串转换为 common.Address 类型
	account := common.HexToAddress(address)

	// 查询余额（返回的是 Wei 单位）
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 将 Wei 转换为 ETH
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("地址 %s 的余额: %s ETH\n", address, ethBalance.String())
}

// 查询交易信息
func queryTransaction(client *ethclient.Client, txHash string) {
	ctx := context.Background()

	hash := common.HexToHash(txHash)
	tx, isPending, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		log.Fatal(err)
	}

	if isPending {
		fmt.Println("交易仍在等待确认")
	} else {
		fmt.Println("交易已确认")
	}

	// 获取发送方地址（从交易签名恢复）
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	from, err := types.Sender(types.NewEIP155Signer(chainID), tx)
	if err != nil {
		log.Fatal(err)
	}

	var to string
	if tx.To() != nil {
		to = tx.To().Hex()
	} else {
		to = "contract creation"
	}

	var gasPrice string
	if tx.GasPrice() != nil {
		gasPrice = tx.GasPrice().String()
	} else {
		gasPrice = "nil"
	}

	fmt.Printf("发送方: %s\n", from.Hex())
	fmt.Printf("接收方: %s\n", to)
	fmt.Printf("金额: %s Wei\n", tx.Value().String())
	fmt.Printf("Gas 价格: %s Wei\n", gasPrice)
}

// 5. 监听新区块
func watchNewBlocks(client *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Printf("新区块: %s, 时间: %d\n",
				header.Number.String(), header.Time)
		}
	}
}

// ERC20 标准 ABI 的部分定义
const erc20ABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"}]`

// 查询 ERC-20 代币余额
func queryTokenBalance(client *ethclient.Client, tokenAddress string, walletAddress string) {
	ctx := context.Background()

	// 解析合约 ABI
	contractAbi, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Fatal(err)
	}

	// 编码 balanceOf 函数调用
	data, err := contractAbi.Pack("balanceOf", common.HexToAddress(walletAddress))
	if err != nil {
		log.Fatal(err)
	}

	// 调用合约
	tokenAddr := common.HexToAddress(tokenAddress)
	msg := ethereum.CallMsg{
		To:   &tokenAddr,
		Data: data,
	}

	result, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 解码结果
	var balance *big.Int
	err = contractAbi.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("代币余额: %s\n", balance.String())
}

func sendTransaction(client *ethclient.Client, privateKeyStr, toAddress string, amount *big.Int) {
	ctx := context.Background()

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

	fmt.Printf("交易已发送: %s\n", signedTx.Hash().Hex())
}
