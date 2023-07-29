package database_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mvndaai/go-interface-demo/testable/database"
	"github.com/mvndaai/go-interface-demo/testable/database/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseMock(t *testing.T) {
	expectedErr := errors.New("foo")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbMock := mocks.NewMockIDatabase(mockCtrl)
	dbMock.EXPECT().Save(gomock.Any()).Return(expectedErr)

	s := struct {
		db database.IDatabase
	}{
		db: dbMock,
	}
	err := s.db.Save("")
	assert.Equal(t, expectedErr, err)
}
