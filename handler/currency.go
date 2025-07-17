package handler

import (
	"net/http"
	"strconv"

	"github.com/JeffersonDallalibera/currency-api/model"
	"github.com/JeffersonDallalibera/currency-api/service"
	"github.com/gin-gonic/gin"
)

// ConvertCurrency godoc
// @Summary Converte valores entre moedas
// @Description Recebe moeda de origem, destino e valor. Retorna valor convertido.
// @Tags Conversão
// @Accept json
// @Produce json
// @Param from query string true "Moeda de origem (ex: USD)"
// @Param to query string true "Moeda de destino (ex: BRL)"
// @Param amount query number false "Valor a ser convertido"
// @Success 200 {object} model.ConvertResponse
// @Failure 400 {object} map[string]string
// @Router /convert [get]
func ConvertCurrency(c *gin.Context) {

	from := c.Query("from")
	to := c.Query("to")
	amountStr := c.Query("amount")

	if from == "" || to == ""  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros 'from' e 'to' são obrigatórios"})
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		amount = 0 // Se não for informado, assume 0
	}

	converted, rate, err := service.Convert(from, to, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter taxa de câmbio"})
		return
	}

	response := model.ConvertResponse{
		From:      from,
		To:        to,
		Amount:    amount,
		Converted: converted,
		Rate:      rate,
	}

	c.JSON(http.StatusOK, response)
}
