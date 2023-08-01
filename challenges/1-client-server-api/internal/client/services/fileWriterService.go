package services

import (
	"bufio"
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

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	err = t.Execute(writer, data)
	if err != nil {
		return err
	}

	return nil
}
