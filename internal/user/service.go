package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"hex-example/internal/env"
	"time"
)

type UserService interface {
	CreateAccount(account *Account) error
	Login(username, password string) (*Login, error)
}

type userService struct {
	repo UserRepo
	key []byte
}

func NewUserService(repo UserRepo) UserService {
	key := env.EnvString("SECRET", "secret")
	return &userService{
		repo,
		[]byte(key),
	}
}

func (s *userService) CreateAccount(account *Account) error {

	logrus.Infof("Account password: %s", account.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	if err != nil {
		logrus.WithField("error", err).Error("Unable to hash password: ")
		return err
	}

	account.Password = string(hashedPassword)
	account.ID = uuid.New().String()

	err = s.repo.CreateAccount(account)
	if err != nil {
		logrus.WithField("error", err).Error("Unable save password")
		return err
	}

	account.Password = ""

	return err
}

func (s *userService) Login(username, password string) (*Login, error) {

	account, err := s.repo.GetUser(username)

	if err != nil {
		logrus.WithField("username", username).Error("Unable to fetch account")
		return nil, err
	}

	if account == nil {
		return nil, errors.New("Invalid Username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		logrus.WithFields(logrus.Fields{"username": username, "error": err.Error()}).Error("Invalid login")
		return nil, err
	}

	token, err := s.getToken(account)
	if err != nil {
		logrus.WithFields(logrus.Fields{"username": username, "error": err}).Error("Unable to generate token")
	}
	login := &Login{
		Username: username,
		Token: token,
	}
	return login, nil
}

func (s *userService) getToken(account *Account) (string, error) {

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["sub"] = account.ID
	claims["type"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	return token.SignedString(s.key)
}
