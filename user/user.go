package user

import "golang-gin/model"

type UserData struct {
	data []model.Interface
}

func NewUser() *UserData {
	return &UserData{}
}

func (u *UserData) AllData() []model.Interface {
	return u.data
}

func (u *UserData) AddData(d model.Interface) model.Interface {
	u.data = append(u.data, d)
	return d
}
