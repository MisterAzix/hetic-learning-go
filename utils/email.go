package utils

import (
	"fmt"
	"hetic-learning-go/model"
	"log"
	"net/smtp"
)

var smtpHost = "localhost"
var smtpPort = "1025"

func SendOrderEmail(order model.Order) error {
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	auth := smtp.PlainAuth("", "", "", smtpHost)

	from := "sender@example.com"
	to := []string{"recipient@example.com"}

	subject := "Order confirmation"
	body := fmt.Sprintf("Thanks for your order! \n\nCommand detail:\nID: %d\nQuantity: %d\nTotal price: %.2f\nOrder at: %s",
		order.Id, order.Quantity, order.TotalPrice, order.OrderAt)
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to[0], subject, body))

	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Printf("smtp error: %s", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Print("Email sent successfully!")
	return nil
}
