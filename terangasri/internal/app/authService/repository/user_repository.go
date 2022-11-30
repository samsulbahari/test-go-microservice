package repository

import (
	"fmt"
	"terangasri/internal/app/domain"

	"github.com/go-resty/resty/v2"
	"github.com/thedevsaddam/gojsonq/v2"
)

type UserServiceRepository struct {
	Hostname string
}

type Response struct {
}

func NewUserServiceRepository() *UserServiceRepository {
	return &UserServiceRepository{
		Hostname: "http://localhost:8080/",
	}
}

func (user UserServiceRepository) Login(auth domain.Login) (*resty.Response, interface{}, error) {
	endpoint := "login"
	url := fmt.Sprintf("%s%s", user.Hostname, endpoint)

	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(auth).
		SetResult(Response{}).
		Post(url)

	result := gojsonq.New().FromString(resp.String()).Get()

	return resp, result, err

}

func (user UserServiceRepository) Register(auth domain.Register) (*resty.Response, interface{}, error) {
	endpoint := "register"
	url := fmt.Sprintf("%s%s", user.Hostname, endpoint)

	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(auth).
		SetResult(Response{}).
		Post(url)

	result := gojsonq.New().FromString(resp.String()).Get()

	return resp, result, err

}
