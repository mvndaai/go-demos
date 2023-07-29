package traveler_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mvndaai/go-interface-demo/reuseable/traveler"
	"github.com/mvndaai/go-interface-demo/reuseable/traveler/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTravelerMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	travelerMock := mocks.NewMockTraveler(mockCtrl)
	mockReturn := 10.00
	travelerMock.EXPECT().Travel(gomock.Any()).Return(mockReturn)

	s := struct {
		Traveler traveler.Traveler
	}{
		Traveler: travelerMock,
	}

	actual := s.Traveler.Travel(time.Second)
	assert.Equal(t, mockReturn, actual)
}
