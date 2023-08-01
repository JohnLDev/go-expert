package services

import (
	"fmt"
	"os"
	"text/template"
)

func SaveDollarPrice(data DollarPriceResult) error {
	t := template.Must(template.New("dollarPrice").Parse("Dólar: {{.Bid}}\n"))
	data.Bid = fmt.Sprintf("{%s}", data.Bid)

	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
