package client

import (
	"bytes"
	"github.com/mrchar/seedpod/server/handler"
	"github.com/mrchar/seedpod/utils"
	"github.com/pkg/errors"
	"net/http"
)

type Client struct {
	baseURL string
	c       *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		c:       &http.Client{},
	}
}
func (c *Client) Register(accountName, password string) error {
	request := handler.RegisterRequest{AccountName: accountName, Password: password}
	response, err := c.c.Post(c.baseURL+"/register", "application/json", bytes.NewReader(request.ToJSON()))
	if err != nil {
		return err
	}

	registerResponse := handler.RegisterResponse{}
	err = utils.UnmarshalResponse(response, &registerResponse)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.Wrap(
			ErrorServerReported{response.StatusCode, registerResponse.Message},
			"注册失败",
		)
	}

	return nil
}

func (c *Client) Login(accountName, password string) error {
	request := handler.RegisterRequest{AccountName: accountName, Password: password}
	response, err := c.c.Post(c.baseURL+"/login", "application/json", bytes.NewReader(request.ToJSON()))
	if err != nil {
		return err
	}

	loginResponse := handler.LoginResponse{}
	err = utils.UnmarshalResponse(response, &loginResponse)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.Wrap(
			ErrorServerReported{response.StatusCode, loginResponse.Message},
			"登录失败",
		)
	}

	return nil
}
