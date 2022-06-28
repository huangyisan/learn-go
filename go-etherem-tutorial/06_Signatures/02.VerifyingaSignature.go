package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	// 验证签名3个因素. 1.签名  2.原始数据的hash 3.签名者公钥
	// 私钥加密,公钥解密. 私钥签名,公钥验证签名
	// 发送者用私钥进行签名，接收者就只能用发送者的公钥进行验证该信息的发送者就是私钥的持有者。
	// 发送报文时，发送者用一个哈希函数从报文文本中生成摘要信息，将摘要信息用发送者的私钥加密，加密后的摘要信息将作为报文的数字签名与报文一起发送给接收者。接收者先用与发送者一样的哈希函数对收到的原文计算，产生一个摘要信息1，然后用发送者的公钥来解密被加密的数字签名（摘要信息）得到解密后的摘要信息2，这两个摘要信息（1和2）如果相同，那么接收者就能确认该数字签名是发送者的。也说明收到的信息是完整的，在传输过程中没有被修改。
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
	// 从原始数据hash中使用Ecrecover方法恢复公钥,字节格式的公钥
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sigPublicKey)
	fmt.Println(publicKeyBytes)
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	// 也可以用SigToPub方法,他将返回ECDSA类型的公钥
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	// 为方便起见，go-ethereum/crypto包提供了VerifySignature函数，该函数接收原始数据的签名，哈希值和字节格式的公钥。 它返回一个布尔值，如果公钥与签名的签名者匹配，则为true。 一个重要的问题是我们必须首先删除signture的最后一个字节，因为它是ECDSA恢复ID，不能包含它。
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}
