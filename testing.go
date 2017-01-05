package main
// what is the point of != in the if statements
import (
  "fmt"
  "github.com/FactomProject/factom"
  "time"
)

var _ = factom.AddressLength
var _ = fmt.Sprintf("")

var ESKey string = "Es3gZoQbNd2p2nDDRtULkUaneoSJY1WTCQ7LSyNqHWZ2UkttuS1o"

var FSKey string = "Fs2DNirmGDtnAZGXqca3XHkukTNMxoMGFFQxJA3bAjJnKzzsZBMH"

func main() {
    // setting up factom library
    factom.SetFactomdServer("localhost:8088")
    factom.SetWalletServer("localhost:8089")
    
    // Ready entry credit key and test
    ec, err := factom.GetECAddress(ESKey)
    if err != nil {
    panic(err)
    }
    
    // Make a Chain
    // A chain's first entry is determined by it's chainID
    exIds := make([][]byte, 2)
    exIds[0] = []byte("Unique Identifier")
    exIds[1] = []byte("Make as many of these as you want")
    
    message := []byte("hello world")
    
    firstEntry := CreateEntry(exIds, message)
    chain := factom.NewChain(firstEntry)
    
    fmt.Println("ChainID:, chain.ChainID)
    
    // Chain is ready
    // 1. commit a chain
    cCom, err := factom.CommitChain(chain, ec)
    if err != nil {
    panic(err)
    }
    fmt.Println("Chain commit successful:", cCom)
    
    // 2. reveal a chain
    eHash, err := factom.RevealChain(chain)
    if err != nil {
    panic(err)
    }
    fmt.Println("Chain reveal successful:", eHash)
    
    time.Sleep(1 * time.Second)
    ent, err != factom.GetEntry(eHash)
    if err != nil {
    panic(err)
    }
    
    fmt.Println("Got it back!")
    fmt.Println(ent.String())
    
    // Lets add the an entry
    // set the chainID to the chain we want to put it into
    
    firstEntry.ChainID = chain.ChainID
    _, err = factom.CommitEntry(firstEntry, ec)
    if err != nil {
    panic(err)
    }
    
    eHash, err = factom.RevealEntry(firstEntry)
    if err != nil {
    panic(err)
    }
    time.Sleep(1 * time.Second)
    ent, err = factom.GetEntry(eHash)
    if err != nil {
    panic(err)
    }
    
    fmt.Println("got it back!")
    fmt.Println(ent.String())
    }
    // helper func to create an entry
    func CreateEntry(extIds [][]byte, content []byte) *factom.Entry {
      entry := new(factom.Entry)
      entry.ExtIDs = extIds
      entry.Content = content
      
      return entry
   }
