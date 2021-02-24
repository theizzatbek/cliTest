package cmd

import (
	"cliTest/models"
	email "cliTest/utils/email"
	sms "cliTest/utils/sms"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func actions(cmd *cobra.Command, args []string) {

	if userId < 1 || productId < 1 {
		print("user id or product id is not supported", nil)
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
		print("user is not found", errors.Wrap(err, "get user by id"))
		os.Exit(1)
	}
	product := models.Product{ID: productId}
	if err := product.GetById(); err != nil {
		print("product is not found", errors.Wrap(err, "get product by id"))
		os.Exit(1)
	}

	return user, product
}

func sendSms() {
	user, product := get()
	if err := sms.Send(user.PhoneNumber, fmt.Sprintf("Product: %s\nPrice: %d\n", product.Name, product.Price)); err == nil {
		print("SMS Sent!", nil)
	} else {
		print("SMS sent failed", errors.Wrap(err, "sms sent"))
		os.Exit(1)
	}
}

func sendEmail() {
	user, product := get()
	if err := email.Send(user.Email,
		fmt.Sprintf("Product: %s\nPrice: %d\n", product.Name, product.Price)); err == nil {
		print("Email Sent!", nil)
	} else {
		print("Email Sent!", errors.Wrap(err, "email sent"))
		os.Exit(1)
	}

}
