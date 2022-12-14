package helper

import (
	"fmt"
	"os"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

type Recipient struct {
	OrderID      string
	Name         string
	Email        string
	Handphone    string
	TotalPayment string
	PaymentTime  string
}

func SendEmail(data Recipient) {
	temp := `
	<table>
		<tr>
			<td colspan="2">
				<h3>--------- Payment Success --------</h3>
			</td>
		</tr>
		<tr>
			<td>
				ID Order 
			</td>
			<td>
				: %s
			</td>
		</tr>
		<tr>
			<td>
				Name 
			</td>
			<td>
				: %s
			</td>
		</tr>
		<tr>
			<td>
				Total Payment
			</td>
			<td>
			: %s
			</td>
		</tr>
		<tr>
			<td>
				Payment Time 
			</td>
			<td>
				: %s
			</td>
		</tr>
		<tr>
			<td colspan="2">
				<h3>--------- Wisata Foodie -----------</h3>
			</td>
		</tr>
	</table>
	`

	host := os.Getenv("GMAIL_SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("GMAIL_SMTP_PORT"))
	email := os.Getenv("GMAIL_AUTH_EMAIL")
	password := os.Getenv("GMAIL_AUTH_PASSWORD")

	msg := gomail.NewMessage()
	msg.SetHeader("From", "altacapstonegroup3@gmail.com")
	msg.SetHeader("To", data.Email)
	msg.SetHeader("Subject", "Notification Payment")
	msg.SetBody("text/html", fmt.Sprintf(temp, data.OrderID, data.Name, data.TotalPayment, data.PaymentTime))

	n := gomail.NewDialer(host, port, email, password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		fmt.Println(err.Error())
	}
}
