//go:generate mockgen -destination=rocket_mocks_text.go -package=rocket github.com/Ekod/go-grpc/internal/rocket Store

package rocket

import "context"

//Rocket - содержит определение рокеты
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

//Service - ответственен за обновление инвентаря рокеты
type Service struct {
	Store Store
}

//New - возвращает инстанс Сервиса
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - достаёт рокету из store по ID
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - Добавляет рокету в store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - удаляет рокету - очень быстрый демонтаж
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
