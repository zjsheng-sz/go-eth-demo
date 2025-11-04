package goethdemo

func main() {
    // 连接节点
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_PROJECT_ID")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // 获取最新区块
    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("最新区块: %d\n", header.Number.Int64())

    // 查询 Vitalik 的余额
    vitalik := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
    balance, err := client.BalanceAt(context.Background(), vitalik, nil)
    if err != nil {
        log.Fatal(err)
    }

    ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
    fmt.Printf("Vitalik 余额: %f ETH\n", ethBalance)
}