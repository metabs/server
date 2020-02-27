package jwt

import (
	"errors"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/google/uuid"
	"github.com/unprogettosenzanomecheforseinizieremo/server/customer"
	"os"
	"time"
)

var alg = jwt.NewHS256([]byte(os.Getenv("JWT_SECRET")))
var issuer = os.Getenv("JWT_ISSUER")
var audience = os.Getenv("JWT_AUDIENCE")
var subject = os.Getenv("JWT_SUBJECT")

type SignerVerifier struct {
	Alg      jwt.Algorithm
	Issuer   string
	Audience string
	Subject  string
}

type payload struct {
	jwt.Payload
	ID        customer.ID
	Email     customer.Email
	Status    customer.Status
	Created   time.Time
	Activated time.Time
	Updated   time.Time
}

func New() (*SignerVerifier, error) {

	if alg == nil {
		return nil, errors.New("internal.jwt: could not use invalid alg")
	}
	if issuer == "" {
		return nil, errors.New("internal.jwt: could not use invalid issuer")
	}
	if audience == "" {
		return nil, errors.New("internal.jwt: could not use invalid audience")
	}
	if subject == "" {
		return nil, errors.New("internal.jwt: could not use invalid subject")
	}

	return &SignerVerifier{
		Alg:      alg,
		Issuer:   issuer,
		Audience: audience,
		Subject:  subject,
	}, nil
}

func (s *SignerVerifier) Sign(
	ID customer.ID,
	Email customer.Email,
	Status customer.Status,
	Created,
	Activated,
	Updated time.Time,
) ([]byte, error) {

	pl := payload{
		Payload: jwt.Payload{
			Issuer:         issuer,
			Subject:        subject,
			Audience:       jwt.Audience{audience},
			ExpirationTime: jwt.NumericDate(time.Now().Add(12 * time.Hour)),
			NotBefore:      jwt.NumericDate(time.Now()),
			IssuedAt:       jwt.NumericDate(time.Now()),
			JWTID:          uuid.New().String(),
		},
		ID:        ID,
		Email:     Email,
		Status:    Status,
		Created:   Created,
		Activated: Activated,
		Updated:   Updated,
	}

	token, err := jwt.Sign(pl, alg)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *SignerVerifier) Verify(token string) (customer.ID, customer.Status, error) {
	var pl payload
	_, err := jwt.Verify([]byte(token), s.Alg, &pl)
	if err != nil {
		return "", "", err
	}

	return pl.ID, pl.Status, nil
}
