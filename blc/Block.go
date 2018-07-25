package blc

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

//定义区块
type Block struct {

	//1.区块高度，也就是区块的编号，第几个区块
	Height int64

	//2.上一个区块的Hash值
	PreBlockHash []byte

	//3.交易数据（最终都属于transaction 事务）
	Data []byte

	//4.创建时间的时间戳
	Timestamp int64

	//5.当前区块的Hash值
	Hash []byte

	//6.Nonce 随机数，用于验证工作量证明
	Nonce int64

}

func (block *Block) SetHash()  {
	heightBytes := IntToHex(block.Height)

	timeString := strconv.FormatInt(block.Timestamp,2)
	timeBytes := []byte(timeString)
	blockBytes := bytes.Join([][]byte{heightBytes,block.PreBlockHash,block.Data, timeBytes, block.Hash}, []byte{})
	//4.生成Hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]

}

//1. 创建新的区块
func NewBlock(data string, height int64, PreBlockHash []byte) *Block {
	//创建区块
	block := &Block{
		height,
		PreBlockHash,
		[]byte(data),
		time.Now().Unix(),
		nil,
		0,
	}
	//设置Hash
	block.SetHash()
	return block

}

//2.生成创世区块
func CreateGenesisBlock(data string) *Block {

	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

}