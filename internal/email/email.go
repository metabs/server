package email

import (
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/unprogettosenzanomecheforseinizieremo/server/email"
	"go.uber.org/zap"
	"os"
)

var apiKey = os.Getenv("SENDGRID_API_KEY")
var name = os.Getenv("EMAIL_NAME")
var address = os.Getenv("EMAIL_ADDRESS")
var activationTemplateID = os.Getenv("EMAIL_ACTIVATION_TEMPLATE_ID")
var activationTemplateURL = os.Getenv("EMAIL_ACTIVATION_TEMPLATE_URL")

func New(logger *zap.SugaredLogger) (*email.Sender, error) {
	if apiKey == "" {
		return nil, errors.New("internal.email: could not get api key")
	}

	if name == "" {
		return nil, errors.New("internal.email: could not get name")
	}

	if address == "" {
		return nil, errors.New("internal.email: could not get address")
	}

	if activationTemplateID == "" {
		return nil, errors.New("internal.email: could not get activationTemplateID")
	}

	if activationTemplateURL == "" {
		return nil, errors.New("internal.email: could not get activationTemplateURL")
	}

	return email.New(
		sendgrid.NewSendClient(apiKey),
		&mail.Email{Address: address, Name: name},
		activationTemplateID,
		activationTemplateURL,
		logger,
	), nil
}
