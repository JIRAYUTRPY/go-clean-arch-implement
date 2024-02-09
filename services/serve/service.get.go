package services

type GetServeUseCase interface {
	Gets() error
	GetByServeId(id uint) error
	GetByUserId(id uint) error
	GetByServeName(name string) error
}