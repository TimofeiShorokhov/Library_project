package email

import (
	"Library_project/model"
	"Library_project/repo"
	"fmt"
	"net/smtp"
	"strconv"
	"time"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendEmail(email string, message string) {

	// Sender data.
	from := "limjkeee@mail.ru"
	password := "Xfp6Gmb12VsrYGpP5HvD"
	// Receiver email address.
	to := []string{
		email,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.mail.ru", port: "587"}
	// Message.
	mes := []byte(message)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, mes)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CheckForSend() {
	var mes1 string
	var mes2 string
	var documents []repo.Document
	var readers []repo.Reader

	for _, i := range model.GetDocuments(documents) {
		t1 := i.Date
		dt1, _ := time.Parse("2006-01-02", t1)
		t2 := time.Now()

		for _, c := range model.GetReaders(readers) {

			if i.ReaderSurname == c.Surname {
				mes1 = "Вам необходимо вернуть книгу: " + i.BookName + ". Если не вернете в течении 5 дней с момента окончания срока возврата, то будет выставлен штраф"

				if (t2.Sub(dt1).Hours())/24 <= 6 {
					SendEmail(c.Email, mes1)
				} else if (t2.Sub(dt1).Hours())/24 > 6 {
					days := (t2.Sub(dt1).Hours()) / 24
					fine := float64(i.Price) * (0.01) * (days)
					fineInString := strconv.Itoa(int(fine))
					daysInString := strconv.Itoa(int(days))
					mes2 = "Вы просрочили сдачу, прошло уже " + daysInString + " дней, сумма штрафа:" + fineInString
					SendEmail(c.Email, mes2)
				}

			}
		}

	}
}
