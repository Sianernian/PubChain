package chain

import (
	"PublicChain/utils"
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct{
	Height  int64
	Version int64
	PreHash [32]byte
	Hash    [32]byte
	// 默克尔根
	Timestamp int64
	//Difficulty int64
	Nonce int64
	Data  []byte //区块体

}

func (block *Block) SetHash() {
	heightByte ,_ :=utils.IntToByte(block.Height)
	versionByte ,_:=utils.IntToByte(block.Version)
	timeByte ,_:=utils.IntToByte(block.Timestamp)
	nonceByte ,_:=utils.IntToByte(block.Nonce)
	blockbyte :=bytes.Join([][]byte{
		heightByte,
		versionByte,
		timeByte,
		nonceByte,
		block.Data,
		block.PreHash[:],
	},[]byte{})
	hash :=sha256.Sum256(blockbyte)
	block.Hash =hash

}

func CreateBlock(height int64 ,prevHash [32]byte ,data []byte)Block{
	block :=Block{}
	block.Height =height + 1
	block.PreHash = prevHash
	block.Version = 0X00
	block.Timestamp = time.Now().Unix()
	block.Data = data

	block.SetHash() // 计算hash
	return block
}


func CreateGenesisBlock(data []byte)Block{
	genesis :=Block{}
	genesis.Height =0
	genesis.PreHash =[32]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	genesis.Version = 0X01
	genesis.Timestamp = time.Now().Unix()
	genesis.Data =data

	genesis.SetHash() // 计算hash
	return genesis
}