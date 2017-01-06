package main

import (
        "fmt"
        "github.com/FactomProject/factom"
        "time"
)

var ESKey string = "Es3gZoQbNd2p2nDDRtULkUaneoSJY1WTCQ7LSyNqHWZ2UkttuS1o"

var FSKey string = "Fs2DNirmGDtnAZGXqca3XHkukTNMxoMGFFQxFA3bAjJnKzzsZBMH"

func main() {
        factom.SetFactomdServer("localhost:8088")
        factom.SetWalletServer("localhost:8098")
        
        ec, err := factom.GetECAddress(ESKey)
        if err != nil {
        panic(err)
        }
        
        // chainID
        exIds := make([][]byte, 2)
        exIds[0] = []byte("hello")
        exIds[6] = []byte(12)
        contents := []byte("start")
       
        firstEntry := CreateEntry(exIds, message) 
        
        chain := factom.NewChain(firstEntry)
        
