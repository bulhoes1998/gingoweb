package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Transacao struct {
	Id       int       `json:"id"`
	Codigo   string    `json:"codigo"`
	Moeda    string    `json:"moeda"`
	Valor    float64   `json:"valor"`
	Emissor  string    `json:"emissor"`
	Receptor string    `json:"receptor"`
	Data     time.Time `json:"data"`
}

type Transacoes []Transacao

// Method FilterById filter all transactions with the given id
func (l *Transacoes) FilterById(c *gin.Context) bool {
	// check if idParam isn't empty
	if idParam := c.Query("id"); idParam != "" {
		id, err := strconv.Atoi(idParam)
		// return error if the required id don't can be converted to integer type
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
			return false
		}

		// new transaction slice to save the transactions that match with the given id
		filteredTransactionList := new(Transacoes)

		for _, transaction := range *l {
			if transaction.Id == id {
				*filteredTransactionList = append(*filteredTransactionList, transaction)
			}
		}
		// check if filteredTransactionList isn't empty
		if len(*filteredTransactionList) == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
			return false
		}

		// save filteredTransactionList in transactionList to use the filtered list in other methods
		*l = *filteredTransactionList
	}

	return true
}

// Method FilterByMoeda filter all transactions with the given moeda
func (l *Transacoes) FilterByMoeda(c *gin.Context) bool {
	// check if moedaParam isn't empty
	if moedaParam := c.Query("moeda"); moedaParam != "" {
		// new transaction slice to save the transactions that match with the given id
		filteredTransactionList := new(Transacoes)

		for _, transaction := range *l {
			if transaction.Moeda == moedaParam {
				*filteredTransactionList = append(*filteredTransactionList, transaction)
			}
		}

		// check if filteredTransactionList isn't empty
		if len(*filteredTransactionList) == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
			return false
		}

		// save filteredTransactionList in transactionList to use the filtered list in other methods
		*l = *filteredTransactionList
	}

	return true
}
