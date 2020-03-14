package email

import (
	"context"
	"errors"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"net/http"
)

var ErrSendingEmail = errors.New("email: could not send email")

type Sender struct {
	Client                  *sendgrid.Client
	FromEmail               *mail.Email
	ActivationTemplateID    string
	ActivationURL           string
	ResetPasswordTemplateID string
	ResetPasswordURL        string
	ChangeEmailTemplateID   string
	ChangeEmailURL          string
	Logger                  *zap.SugaredLogger
}

func New(
	client *sendgrid.Client,
	from *mail.Email,
	activationTemplateID, activationURL, changeEmailTemplateID, changeEmailTemplateURL, resetPasswordTemplateID, resetPasswordTemplateURL string,
	logger *zap.SugaredLogger,
) *Sender {
	return &Sender{
		Client:                  client,
		FromEmail:               from,
		ActivationTemplateID:    activationTemplateID,
		ActivationURL:           activationURL,
		ChangeEmailTemplateID:   changeEmailTemplateID,
		ChangeEmailURL:          changeEmailTemplateURL,
		ResetPasswordTemplateID: resetPasswordTemplateID,
		ResetPasswordURL:        resetPasswordTemplateURL,
		Logger:                  logger,
	}
}

func (s *Sender) SendActivationEmail(ctx context.Context, to, id, hash string) error {
	_, span := trace.StartSpan(ctx, "SendActivationEmail")
	defer span.End()

	logger := s.Logger.With("action", "SendActivationEmail", "to", to, "id", id)

	res, err := s.Client.Send(&mail.SGMailV3{
		From: s.FromEmail,
		Personalizations: []*mail.Personalization{{
			To:                  []*mail.Email{{Address: to}},
			DynamicTemplateData: map[string]interface{}{"url": fmt.Sprintf(s.ActivationURL, id, hash)},
		}},
		TemplateID: s.ActivationTemplateID,
	})

	if err != nil {
		logger.With("error", err).Error("could not send email")
		return fmt.Errorf("%w: %s", ErrSendingEmail, err)
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		logger.With("status_code", res.StatusCode).Error("could not send email")
		return fmt.Errorf("%w: invalid status code", ErrSendingEmail)
	}

	logger.Debug("email sent")

	return nil
}

func (s *Sender) SendResetPassword(ctx context.Context, to, id, hash string) error {
	_, span := trace.StartSpan(ctx, "SendResetPassword")
	defer span.End()

	logger := s.Logger.With("action", "SendResetPassword", "to", to, "id", id)

	res, err := s.Client.Send(&mail.SGMailV3{
		From: s.FromEmail,
		Personalizations: []*mail.Personalization{{
			To:                  []*mail.Email{{Address: to}},
			DynamicTemplateData: map[string]interface{}{"url": fmt.Sprintf(s.ResetPasswordURL, id, hash)},
		}},
		TemplateID: s.ResetPasswordTemplateID,
	})

	if err != nil {
		logger.With("error", err).Error("could not send email")
		return fmt.Errorf("%w: %s", ErrSendingEmail, err)
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		logger.With("status_code", res.StatusCode).Error("could not send email")
		return fmt.Errorf("%w: invalid status code", ErrSendingEmail)
	}

	logger.Debug("email sent")

	return nil
}

func (s *Sender) SendChangeEmail(ctx context.Context, to, id, hash string) error {
	_, span := trace.StartSpan(ctx, "SendChangeEmail")
	defer span.End()

	logger := s.Logger.With("action", "SendChangeEmail", "to", to, "id", id)

	res, err := s.Client.Send(&mail.SGMailV3{
		From: s.FromEmail,
		Personalizations: []*mail.Personalization{{
			To:                  []*mail.Email{{Address: to}},
			DynamicTemplateData: map[string]interface{}{"url": fmt.Sprintf(s.ChangeEmailURL, id, hash)},
		}},
		TemplateID: s.ChangeEmailTemplateID,
	})

	if err != nil {
		logger.With("error", err).Error("could not send email")
		return fmt.Errorf("%w: %s", ErrSendingEmail, err)
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		logger.With("status_code", res.StatusCode).Error("could not send email")
		return fmt.Errorf("%w: invalid status code", ErrSendingEmail)
	}

	logger.Debug("email sent")

	return nil
}
