package block

import (
    "bytes"
    "crypto/sha256"
    "encoding/gob"
    "log"
    "time"
)

// 区块结构体
type Block struct {
    Version    uint64 // 版本号
    PrevHash   []byte // 前一个区块 hash
    MerKleRoot []byte // 梅特根
    Timestamp  uint64 // 时间戳
    Difficulty uint64 // 挖矿难度值
    Nonce      uint64 // 随机数
    // Data       []byte // 数据，后续使用交易替代
    Transactions []*Transaction
    Hash         []byte // 当前区块 hash
}

// 创建区块函数
func NewBlock(txs []*Transaction, prevHash []byte) *Block {
    block := Block{
        Version:    00,
        PrevHash:   prevHash,
        MerKleRoot: []byte{}, // 先填写空
        Timestamp:  uint64(time.Now().Unix()),
        Difficulty: Bits,
        Nonce:      0,
        Hash:       []byte{}, // 先填充为空
        // Data:       []byte(data),
        Transactions: txs,
    }
    // block.setHash()

    block.HashTransactions()

    // 创建工作量证明
    pow := NewProofOfWork(&block)
    hash, nonce := pow.Run()
    block.Hash = hash
    block.Nonce = nonce
    return &block
}

// gob 序列化
func (block *Block) ToBytes() []byte {
    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    err := encoder.Encode(block)
    if err != nil {
        log.Panic(err)
    }
    return buffer.Bytes()
}

// gob 反序列化
func (block *Block) ToBlock(data []byte) {
    decoder := gob.NewDecoder(bytes.NewReader(data))
    err := decoder.Decode(block)
    if err != nil {
        log.Panic(err)
    }
}

// 模拟梅特尔根
func (block *Block) HashTransactions() {
    // 将交易 id 拼接，做一次 hash 运算作为梅特尔根
    var txIds []byte
    for _, tx := range block.Transactions {
        txIds = append(txIds, tx.TxId...)
    }
    hash := sha256.Sum256(txIds)
    block.MerKleRoot = hash[:]
}

// 计算当前区块 hash
//func (block *Block) setHash() {
//    block.Hash = sha256Bytes[:]
//}
