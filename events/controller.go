package events

import (
	"cliTest/common/variables"
	"cliTest/models"
	email "cliTest/services/email"
	sms "cliTest/services/sms"
	"cliTest/utils/log"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func Actions(cmd *cobra.Command, args []string) {

	if variables.UserId < 1 || variables.ProductId < 1 {
		log.Error(errors.New("user id or product id is not supported"))
		os.Exit(1)
	}

	switch args[0] {
	case "sms":
		sendSms()
	case "email":
		sendEmail()
	}
}

func get() (models.User, string) {

	var err error

	defer func() {
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()

	user := models.User{ID: variables.UserId}
	if err = user.GetById(); err != nil {
		return models.User{}, ""
	}
	product := models.Product{ID: variables.ProductId}
	if err = product.GetById(); err != nil {
		return models.User{}, ""
	}

	msg := fmt.Sprintf("Product: %s\nPrice: %d\n", product.Name, product.Price)

	return user, msg
}

func sendSms() {

	user, msg := get()
	if err := sms.Send(user.PhoneNumber, msg); err == nil {
		log.Info("SMS Sent!")
	} else {
		return
	}
}

func sendEmail() {
	user, msg := get()

	if err := email.Send(user.Email, msg); err == nil {
		log.Info("Email Sent!")
	} else {
		log.Error(errors.Wrap(err, "email sent failed"))
		os.Exit(1)
	}

}
