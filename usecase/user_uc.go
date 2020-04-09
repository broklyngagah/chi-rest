package usecase

import (
	"chi-rest/model"

	"github.com/sirupsen/logrus"
)

// GetUser ...
func (uc UC) GetUser() ([]*model.User, error) {
	uc.App.Log.FromDefault().WithFields(logrus.Fields{
		"name": "Broklyn gagah",
		"age":  "33 yo",
	}).Warning("Test warning logs. you can use any level of logs (error, info, panic, fatal, etc)")

	dt, err := model.UserOp.Get(uc.DB)
	return dt, err
}
