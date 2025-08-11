package main

import (
	"crypto/ecdsa"    // 用于生成椭圆曲线密钥对
	"crypto/elliptic" // 提供椭圆曲线算法支持
	"crypto/rand"     // 提供随机数生成器
	"crypto/sha256"   // 提供SHA-256哈希算法
	"encoding/hex"    // 提供十六进制编码解码功能
	"encoding/json"   // 提供JSON序列化反序列化功能
	"fmt"             // 提供格式化输出功能
	"log"             // 提供日志记录功能

	// 提供大整数支持
	"time" // 提供时间相关功能
)

// Transaction 表示区块链中的一笔交易
// 包含发送者地址、接收者地址、交易金额和唯一交易ID
type Transaction struct {
	ID        string  `json:"id"`        // 交易唯一标识符
	Sender    string  `json:"sender"`    // 发送者地址
	Recipient string  `json:"recipient"` // 接收者地址
	Amount    float64 `json:"amount"`    // 交易金额
}

// Block 表示区块链中的一个区块
// 包含区块索引、时间戳、交易列表、前一区块哈希、当前区块哈希和出块矿工地址
type Block struct {
	Index        int           `json:"index"`        // 区块索引，从0开始递增
	Timestamp    string        `json:"timestamp"`    // 区块创建时间戳
	Transactions []Transaction `json:"transactions"` // 区块中包含的交易列表
	PrevHash     string        `json:"prev_hash"`    // 前一区块的哈希值
	Hash         string        `json:"hash"`         // 当前区块的哈希值
	Miner        string        `json:"miner"`        // 出块矿工的地址
}

// Blockchain 表示整个区块链系统
// 管理区块链表、待处理交易池、权威节点集合和轮询索引
type Blockchain struct {
	chain          []Block         // 区块链主链，存储所有区块
	pendingTx      []Transaction   // 待处理交易池，存储尚未打包的交易
	authorities    map[string]bool // 权威节点集合，key为节点地址，value为是否为权威节点
	lastMinerIndex int             // 上一个出块节点的索引，用于轮询选择下一个出块节点
}

// generateTxID 为交易生成唯一ID
// 通过对交易进行JSON序列化后计算SHA-256哈希值，再转换为十六进制字符串得到交易ID
// 参数:
//
//	tx - 需要生成ID的交易对象
//
// 返回值:
//
//	交易ID的十六进制字符串表示
func generateTxID(tx Transaction) string {
	// 将交易对象序列化为JSON字节流
	txData, _ := json.Marshal(tx)
	// 计算JSON字节流的SHA-256哈希值
	hash := sha256.Sum256(txData)
	// 将哈希值转换为十六进制字符串并返回
	return hex.EncodeToString(hash[:])
}

// calculateHash 计算区块的哈希值
// 通过对区块进行JSON序列化后计算SHA-256哈希值，再转换为十六进制字符串
// 参数:
//
//	block - 需要计算哈希的区块对象
//
// 返回值:
//
//	区块哈希的十六进制字符串表示
func calculateHash(block Block) string {
	// 将区块对象序列化为JSON字节流
	blockData, _ := json.Marshal(block)
	// 计算JSON字节流的SHA-256哈希值
	hash := sha256.Sum256(blockData)
	// 将哈希值转换为十六进制字符串并返回
	return hex.EncodeToString(hash[:])
}

// createTransaction 创建新交易并添加到交易池
// 生成交易ID并将交易添加到区块链的待处理交易列表中
// 参数:
//
//	sender - 发送者地址
//	recipient - 接收者地址
//	amount - 交易金额
func (bc *Blockchain) createTransaction(sender, recipient string, amount float64) {
	// 创建交易对象
	tx := Transaction{
		Sender:    sender,    // 发送者地址
		Recipient: recipient, // 接收者地址
		Amount:    amount,    // 交易金额
	}
	// 生成交易ID
	tx.ID = generateTxID(tx)
	// 将交易添加到待处理交易池
	bc.pendingTx = append(bc.pendingTx, tx)
}

// createBlock 创建新区块
// 使用当前交易池中的所有交易创建新块，并清空交易池
// 参数:
//
//	miner - 出块矿工的地址
//
// 返回值:
//
//	新创建的区块对象
func (bc *Blockchain) createBlock(miner string) Block {
	// 获取区块链中的最后一个区块
	lastBlock := bc.chain[len(bc.chain)-1]
	// 创建新区块对象
	newBlock := Block{
		Index:        lastBlock.Index + 1, // 区块索引为上一个区块索引+1
		Timestamp:    time.Now().String(), // 当前时间戳
		Transactions: bc.pendingTx,        // 包含所有待处理交易
		PrevHash:     lastBlock.Hash,      // 前一区块的哈希值
		Miner:        miner,               // 出块矿工地址
	}

	// 计算新区块的哈希值
	newBlock.Hash = calculateHash(newBlock)

	// 清空交易池，所有交易已被打包到新块中
	bc.pendingTx = []Transaction{}

	// 返回新创建的区块
	return newBlock
}

// validateBlock 验证区块的合法性
// 检查区块是否由权威节点创建、哈希值是否正确以及前哈希是否匹配
// 参数:
//
//	block - 需要验证的区块对象
//
// 返回值:
//
//	区块是否合法的布尔值，true表示合法，false表示不合法
func (bc *Blockchain) validateBlock(block Block) bool {
	// 1. 验证出块节点是否为权威节点
	if !bc.authorities[block.Miner] {
		return false
	}

	// 2. 验证区块哈希是否正确
	calculatedHash := calculateHash(block)
	if block.Hash != calculatedHash {
		return false
	}

	// 3. 验证前哈希是否与区块链中最后一个区块的哈希匹配
	lastBlock := bc.chain[len(bc.chain)-1]
	if block.PrevHash != lastBlock.Hash {
		return false
	}

	// 所有验证通过，返回true
	return true
}

// selectNextMiner 通过轮询方式选择下一个出块节点
// 从权威节点列表中按顺序选择下一个节点，实现轮流出块机制
// 返回值:
//
//	选中的出块节点地址
func (bc *Blockchain) selectNextMiner() string {
	// 将权威节点映射转换为切片，方便按索引访问
	authorities := make([]string, 0, len(bc.authorities))
	for addr := range bc.authorities {
		authorities = append(authorities, addr)
	}

	// 循环轮询选择下一个出块节点索引
	// 初始值为-1，第一次调用时会变为0（第一个节点）
	bc.lastMinerIndex = (bc.lastMinerIndex + 1) % len(authorities)
	// 返回选中的节点地址
	return authorities[bc.lastMinerIndex]
}

// mineBlock 挖矿主流程
// 选择出块节点、创建区块、验证区块并添加到区块链
// 返回值:
//
//	区块是否成功添加到链的布尔值，true表示成功，false表示失败
func (bc *Blockchain) mineBlock() bool {
	// 选择下一个出块节点
	miner := bc.selectNextMiner()

	// 创建新块
	newBlock := bc.createBlock(miner)

	// 验证区块合法性
	if !bc.validateBlock(newBlock) {
		return false
	}

	// 将区块添加到区块链
	bc.chain = append(bc.chain, newBlock)
	return true
}

// createGenesisBlock 创建创世块
// 区块链的第一个区块，没有前序区块，哈希值为"0"
// 参数:
//
//	initialMiner - 创世块的出块矿工地址
//
// 返回值:
//
//	创建的创世块对象
func createGenesisBlock(initialMiner string) Block {
	// 创建创世块对象
	genesisBlock := Block{
		Index:        0,                   // 创世块索引为0
		Timestamp:    time.Now().String(), // 当前时间戳
		Transactions: []Transaction{},     // 创世块不包含交易
		PrevHash:     "0",                 // 前哈希为"0"，表示没有前序区块
		Miner:        initialMiner,        // 创世块出块矿工地址
	}
	// 计算创世块哈希值
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}

// generateKeyPair 生成ECDSA密钥对并返回节点地址
// 使用P256椭圆曲线生成密钥对，公钥哈希后作为节点地址
// 返回值:
//
//	节点地址（公钥哈希的前20字节的十六进制表示）
//	生成的私钥（用于签名交易）
func generateKeyPair() (string, *ecdsa.PrivateKey) {
	// 使用P256椭圆曲线生成ECDSA私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		// 生成密钥失败时记录日志并退出程序
		log.Fatal(err)
	}

	// 从私钥中提取公钥
	publicKey := privateKey.PublicKey
	// 将公钥的X和Y坐标字节流拼接
	publicKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)
	// 计算公钥字节流的SHA-256哈希值，并取前20字节作为节点地址
	hash := sha256.Sum256(publicKeyBytes)
	address := hex.EncodeToString(hash[:20])

	// 返回节点地址和私钥
	return address, privateKey
}

// initBlockchain 初始化区块链
// 创建包含权威节点集合的区块链实例，并添加创世块
// 参数:
//
//	authorities - 权威节点地址列表
//
// 返回值:
//
//	初始化后的区块链实例指针
func initBlockchain(authorities []string) *Blockchain {
	// 初始化权威节点映射
	authorityMap := make(map[string]bool)
	for _, addr := range authorities {
		authorityMap[addr] = true
	}

	// 创建区块链实例
	bc := &Blockchain{
		authorities:    authorityMap, // 设置权威节点集合
		lastMinerIndex: -1,           // 初始化轮询索引为-1
	}

	// 创建创世块并添加到区块链
	genesisBlock := createGenesisBlock(authorities[0])
	bc.chain = append(bc.chain, genesisBlock)

	// 返回区块链实例
	return bc
}

// main 函数是程序的入口点
// 生成权威节点、初始化区块链、添加交易、挖矿并打印区块链信息
func main() {
	// 生成3个权威节点的密钥对和地址
	node1Addr, _ := generateKeyPair()
	node2Addr, _ := generateKeyPair()
	node3Addr, _ := generateKeyPair()

	// 打印权威节点地址
	fmt.Printf("权威节点1地址: %s\n", node1Addr)
	fmt.Printf("权威节点2地址: %s\n", node2Addr)
	fmt.Printf("权威节点3地址: %s\n", node3Addr)

	// 初始化区块链，设置3个权威节点
	bc := initBlockchain([]string{node1Addr, node2Addr, node3Addr})
	fmt.Println("区块链初始化完成，创世块已创建")

	// 添加一些测试交易到交易池
	bc.createTransaction("user1", "user2", 10.5) // user1向user2转账10.5
	bc.createTransaction("user2", "user3", 5.2)  // user2向user3转账5.2

	// 挖第一个区块
	if bc.mineBlock() {
		fmt.Printf("区块 %d 已生成，出块节点: %s\n",
			bc.chain[1].Index, bc.chain[1].Miner)
	}

	// 添加更多测试交易
	bc.createTransaction("user3", "user1", 3.7) // user3向user1转账3.7

	// 挖第二个区块
	if bc.mineBlock() {
		fmt.Printf("区块 %d 已生成，出块节点: %s\n",
			bc.chain[2].Index, bc.chain[2].Miner)
	}

	// 挖第三个区块（此时交易池为空，但仍会生成空块）
	if bc.mineBlock() {
		fmt.Printf("区块 %d 已生成，出块节点: %s\n",
			bc.chain[3].Index, bc.chain[3].Miner)
	}

	// 打印完整区块链信息
	fmt.Println("\n区块链完整信息:")
	for _, block := range bc.chain {
		// 将区块对象转换为格式化的JSON字符串
		blockJson, _ := json.MarshalIndent(block, "", "  ")
		// 打印区块JSON
		fmt.Println(string(blockJson))
	}
}
