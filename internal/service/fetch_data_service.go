package service

type FetchDataService interface {
	GetUsername(username string) error
}
