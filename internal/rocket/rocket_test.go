package rocket

import (
	"context"
	"github.com/golang/mock/gomock"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get rocket by id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().GetRocketByID(id).Return(Rocket{ID: id}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocketByID(context.Background(), id)

		assert.NoError(t, err)
		assert.Equal(t, id, rkt.ID)
	})

	t.Run("tests insert rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().InsertRocket(Rocket{ID: id}).Return(Rocket{ID: id}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.InsertRocket(context.Background(), Rocket{ID: id})

		assert.NoError(t, err)
		assert.Equal(t, id, rkt.ID)
	})

	t.Run("tests delete a rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().DeleteRocket(id).Return(nil)

		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(id)

		assert.NoError(t, err)
	})
}
