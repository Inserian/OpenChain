package main

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "strconv"
        "strings"
        "time"
)
type Block struct {
        data         map[string]interface{}
        hash         string
        previousHash string
        timestamp    time.Time
        pow          int
}
type Blockchain struct {
        genesisBlock Block
        chain        []Block
        difficulty   int
}
func (b *Block) mine(difficulty int) {
        for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
                b.pow++
                b.hash = b.calculateHash()
        }
}
func CreateBlockchain(difficulty int) Blockchain {
        genesisBlock := Block{
                hash:      "0",
                timestamp: time.Now(),
        }
        return Blockchain{
                genesisBlock,
                []Block{genesisBlock},
                difficulty,
        }
}
func (b *Blockchain) addBlock(from, to string, amount float64) {
        blockData := map[string]interface{}{
                "from":   from,
                "to":     to,
                "amount": amount,
        }
        lastBlock := b.chain[len(b.chain)-1]
        newBlock := Block{
                data:         blockData,
                previousHash: lastBlock.hash,
                timestamp:    time.Now(),
        }
        newBlock.mine(b.difficulty)
        b.chain = append(b.chain, newBlock)
}
func (b Blockchain) isValid() bool {
        for i := range b.chain[1:] {
                previousBlock := b.chain[i]
                currentBlock := b.chain[i+1]
                if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
                        return false
                }
        }
        return true
}
func main() {
        // create a new blockchain instance
        blockchain := CreateBlockchain(1)

        blockchain.addBlock("", "", 5)
        blockchain.addBlock("", "", 2)

        // check if the blockchain is valid
        fmt.Println(blockchain.isValid())
}
