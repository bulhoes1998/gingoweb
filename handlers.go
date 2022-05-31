package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getAll(c *gin.Context) {

	jsonFile, err := os.Open("transacoes.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	transactionList := new(Transacoes)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &transactionList)

	canContinue := true

	canContinue = transactionList.FilterById(c)
	if !canContinue {
		return
	}
	canContinue = transactionList.FilterByMoeda(c)
	if !canContinue {
		return
	}

	c.JSON(http.StatusOK, transactionList)
}

func newTransaction(c *gin.Context) {
	var transaction Transacao

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	jsonFile, err := os.Open("transacoes.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	transactionList := new(Transacoes)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &transactionList)

	*transactionList = append(*transactionList, transaction)

	file, err := json.Marshal(transactionList)

	err2 := os.WriteFile("transacoes.json", file, 0666)
	if err2 != nil {
		log.Fatal(err)
	}
}
