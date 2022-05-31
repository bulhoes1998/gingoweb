package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/transacoes", getAll)

	router.Run()
}

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
