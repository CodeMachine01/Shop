package model

type AddCollectionInput struct {
	UserId   uint
	ObjectId int
	Type     int
}

type AddCollectionOutput struct {
	Id uint
}

type DeleteCollectionInput struct {
	Id       uint
	UserId   uint
	ObjectId int
	Type     int
}

type DeleteCollectionOutput struct {
	Id uint
}
