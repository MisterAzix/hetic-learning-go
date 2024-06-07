package utils

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"hetic-learning-go/model"
)

func GenerateOrderPDF(order model.Order) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Order details")
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Order id: %d", order.Id))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("User id: %d", order.UserId))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Product id: %d", order.ProductId))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Quantity: %d", order.Quantity))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Total price: %.2f", order.TotalPrice))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Ordered at: %s", order.OrderAt))
	pdf.Ln(10)

	pdfName := fmt.Sprintf("order_%d.pdf", order.Id)
	err := pdf.OutputFileAndClose(pdfName)
	if err != nil {
		return err
	}

	return nil
}
