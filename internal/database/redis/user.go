package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"hex-example/internal/user"
)

const userTable = "users"

type userRepository struct {
	connection *redis.Client
}

func NewRedisUserRepository(connection *redis.Client) user.UserRepo {
	return &userRepository{
		connection,
	}
}


func (r *userRepository) CreateAccount(account *user.Account) error{

	logrus.Infof("Account %v", account)
	encoded, err := json.Marshal(account)
	if err != nil {
		logrus.Error("Unable to marshal account")
		return err
	}

	cmd := r.connection.HSet(userTable, account.Username, encoded) //Don't expire
	if cmd.Err() != nil {
		logrus.WithField("error", err).Error("Unable to save user account")
		return cmd.Err()
	}
	return nil
}

func (r *userRepository) GetUser(username string) (*user.Account, error){
	b, err := r.connection.HGet(userTable, username).Bytes()

	if err != nil {
		logrus.WithField("username", username).Error("Unable to fetch account")
		return nil, err
	}

	t := new(user.Account)
	err = json.Unmarshal(b, t)

	if err != nil {
		logrus.WithField("username", username).Error("Unable to unmarshal account")
		return nil, err
	}

	return t, nil
}