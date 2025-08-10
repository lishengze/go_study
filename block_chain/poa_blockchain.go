package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"
)

// 交易结构
type Transaction struct {
	ID      string `json:"id"`
	Sender  string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount  float64 `json:"amount"`
}

// 区块结构
type Block struct {
	Index     int           `json:"index"`
	Timestamp string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash  string        `json:"prev_hash"`
	Hash      string        `json:"hash"`
	Miner     string        `json:"miner"` // 出块的权威节点地址
}

// 区块链结构
type Blockchain struct {
	chain        []Block
	pendingTx    []Transaction
	authorities  map[string]bool // 权威节点列表
	lastMinerIndex int           // 记录上一个出块节点索引，用于轮询
}

// 生成交易ID
func generateTxID(tx Transaction) string {
	txData, _ := json.Marshal(tx)
	hash := sha256.Sum256(txData)
	return hex.EncodeToString(hash[:])
}

// 计算区块哈希
func calculateHash(block Block) string {
	blockData, _ := json.Marshal(block)
	hash := sha256.Sum256(blockData)
	return hex.EncodeToString(hash[:])
}

// 创建新交易
func (bc *Blockchain) createTransaction(sender, recipient string, amount float64) {
	tx := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	tx.ID = generateTxID(tx)
	bc.pendingTx = append(bc.pendingTx, tx)
}

// 创建新块
func (bc *Blockchain) createBlock(miner string) Block {
	lastBlock := bc.chain[len(bc.chain)-1]
	newBlock := Block{
		Index:     lastBlock.Index + 1,
		Timestamp: time.Now().String(),
		Transactions: bc.pendingTx,
		PrevHash:  lastBlock.Hash,
		Miner:     miner,
	}
	
	newBlock.Hash = calculateHash(newBlock)
	
	// 清空交易池
	bc.pendingTx = []Transaction{}
	
	return newBlock
}

// 验证区块
func (bc *Blockchain) validateBlock(block Block) bool {
	// 1. 验证出块节点是否为权威节点
	if !bc.authorities[block.Miner] {
		return false
	}
	
	// 2. 验证哈希是否正确
	calculatedHash := calculateHash(block)
	if block.Hash != calculatedHash {
		return false
	}
	
	// 3. 验证前哈希是否匹配
	lastBlock := bc.chain[len(bc.chain)-1]
	if block.PrevHash != lastBlock.Hash {
		return false
	}
	
	return true
}

// 轮询选择下一个出块节点
func (bc *Blockchain) selectNextMiner() string {
	authorities := make([]string, 0, len(bc.authorities))
	for addr := range bc.authorities {
		authorities = append(authorities, addr)
	}
	
	// 循环轮询选择
	bc.lastMinerIndex = (bc.lastMinerIndex + 1) % len(authorities)
	return authorities[bc.lastMinerIndex]
}

// 挖矿（生成新块并添加到链）
func (bc *Blockchain) mineBlock() bool {
	// 选择下一个出块节点
	miner := bc.selectNextMiner()
	
	// 创建新块
	newBlock := bc.createBlock(miner)
	
	// 验证区块
	if !bc.validateBlock(newBlock) {
		return false
	}
	
	// 添加到链
	bc.chain = append(bc.chain, newBlock)
	return true
}

// 创建创世块
func createGenesisBlock(initialMiner string) Block {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Transactions: []Transaction{},
		PrevHash:  "0",
		Miner:     initialMiner,
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}

// 生成新的密钥对（用于节点身份）
func generateKeyPair() (string, *ecdsa.PrivateKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	
	// 公钥作为节点地址
	publicKey := privateKey.PublicKey
	publicKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)
	address := hex.EncodeToString(sha256.Sum256(publicKeyBytes)[:20])
	
	return address, privateKey
}

// 初始化区块链
func initBlockchain(authorities []string) *Blockchain {
	// 初始化权威节点映射
	authorityMap := make(map[string]bool)
	for _, addr := range authorities {
		authorityMap[addr] = true
	}
	
	// 创建创世块，使用第一个权威节点作为初始出块者
	bc := &Blockchain{
		authorities:  authorityMap,
		lastMinerIndex: -1,
	}
	
	genesisBlock := createGenesisBlock(authorities[0])
	bc.chain = append(bc.chain, genesisBlock)
	
	return bc
}

func main() {
	// 生成3个权威节点
	node1Addr, _ := generateKeyPair()
	node2Addr, _ := generateKeyPair()
	node3Addr, _ := generateKeyPair()
	
	fmt.Printf("权威节点1地址: %s\n", node1Addr)
	fmt.Printf("权威节点2地址: %s\n", node2Addr)
	fmt.Printf("权威节点3地址: %s\n", node3Addr)
	
	// 初始化区块链
	bc := initBlockchain([]string{node1Addr, node2Addr, node3Addr})
	fmt.Println("区块链初始化完成，创世块已创建")
	
	// 添加一些交易
	bc.createTransaction("user1", "user2", 10.5)
	bc.createTransaction("user2", "user3", 5.2)
	
	// 生成第一个区块
	if bc.mineBlock() {
		fmt.Printf("区块 %d 已生成，出块节点: %s\n", 
			bc.chain[1].Index, bc.chain[1].Miner)
	}
	
	// 添加更多交易
	bc.createTransaction("user3", "user1", 3.7)
	
	// 生成第二个区块
	if bc.mineBlock() {
		fmt.Printf("区块 %d 已生成，出块节点: %s\n", 
			bc.chain[2].Index, bc.chain[2].Miner)
	}
	
	// 生成第三个区块
	if bc.mineBlock() {
		fmt.Printf("区块 %d 已生成，出块节点: %s\n", 
			bc.chain[3].Index, bc.chain[3].Miner)
	}
	
	// 打印区块链信息
	fmt.Println("\n区块链完整信息:")
	for _, block := range bc.chain {
		blockJson, _ := json.MarshalIndent(block, "", "  ")
		fmt.Println(string(blockJson))
	}
}
