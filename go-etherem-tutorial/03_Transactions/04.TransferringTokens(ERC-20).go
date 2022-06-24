package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func main() {
	// 如果是需要认证的节点地址.则使用以下方法进行认证  用户名:密码@地址
	client, err := ethclient.Dial("https://u0lvw6c3lk:7vUkST4yxgKT_yjF5laCagbtYNCkw0ZOQIYv-FteCvE@u0avnfe4vh-u0j18z3adv-rpc.us0-aws.kaleido.io/")
	if err != nil {
		log.Fatal(err)
	}
	// 获取发送的nonce值
	privateKey, err := crypto.HexToECDSA("6644cda75a855caa92a634ab1117339dc95ffe5dc10cd0c53b943a939f685c7d")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 因为是其他代币,所以eth不需要,为0即可
	value := big.NewInt(0) // in wei (0 eth)
	// gas price估算
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 发送到代币地址的对端钱包地址
	toAddress := common.HexToAddress("0x601Cd71375D4b8b924dCcf4629EBB2161bbeFa17")
	// 合约代币地址
	tokenAddress := common.HexToAddress("0x9c7cad41c2a2b9aefe3721386743321f098cc6cc")

	// 发送ERC-20代币和eth发送不同,他需要和代币合约进行交互, ERC-20规范里面合约代币发送使用transfer方法
	// https://learnblockchain.cn/docs/eips/eip-20.html#api-%E8%A7%84%E8%8C%83
	// transfer 第一个参数为 接收方钱包地址, 第二个参数为发送的代币数量, 将其字节切片
	transferFnSignature := []byte("transfer(address,uint256)") // 里面的内容不需要空格,和参数名称
	// 通过上面得到的签名,对其用keccak-256hash后, 保留前面4个字节, 则可以得到"方法id".  sha3方法已近迁移,使用NewLegacyKeccak256进行hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	// 将发送代币对端接受的钱包地址填充到32位
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	// 发送代币数量
	amount := new(big.Int)
	// 十进制 格式化为wei, 也就是1个token
	amount.SetString("1000000000000000000", 10)
	// 代币数量也要填充到32个字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	// 将 1.方法ID, 2.填充后的地址  3. 填充后的转账量, 依次拼接成数据段字节切片
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// gas limit限制取决于交易数据大小和智能合约必须执行的计算步骤.客户端可以用EstimateGas方法,它可以估算所需的燃气量
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)

	// 构建事务交易.
	LegacyTx := types.LegacyTx{
		Nonce:    nonce,
		Gas:      uint64(77041),
		GasPrice: gasPrice,
		Data:     data,
		Value:    value,
		// to 给合约地址
		To: &tokenAddress,
	}
	tx := types.NewTx(&LegacyTx)

	// 获取ChainID, 因为需要用这个参数和私钥对交易进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发出交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
