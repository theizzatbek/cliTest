package cmd

import (
	"cliTest/models"
	email "cliTest/service/email"
	sms "cliTest/service/sms"
	"cliTest/utils/log"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func actions(cmd *cobra.Command, args []string) {

	if userId < 1 || productId < 1 {
		log.Info("user id or product id is not supported")
		os.Exit(1)
	}

	switch args[0] {
	case "sms":
		sendSms()
	case "email":
		sendEmail()
	}
}

func get() (models.User, models.Product) {
	user := models.User{ID: userId}
	if err := user.GetById(); err != nil {
		log.Error("user is not found", errors.Wrap(err, "get user by id"))
		os.Exit(1)
	}
	product := models.Product{ID: productId}
	if err := product.GetById(); err != nil {
		log.Error("product is not found", errors.Wrap(err, "get product by id"))
		os.Exit(1)
	}

	return user, product
}

func sendSms() {
	user, product := get()
	if err := sms.Send(user.PhoneNumber, fmt.Sprintf("Product: %s\nPrice: %d\n", product.Name, product.Price)); err == nil {
		log.Info("SMS Sent!")
	} else {
		log.Error("SMS sent failed", errors.Wrap(err, "sms sent"))
		os.Exit(1)
	}
}

func sendEmail() {
	user, product := get()
	if err := email.Send(user.Email,
		fmt.Sprintf("Product: %s\nPrice: %d\n", product.Name, product.Price)); err == nil {
		log.Info("Email Sent!")
	} else {
		log.Error("Email Sent!", errors.Wrap(err, "email sent"))
		os.Exit(1)
	}

}
