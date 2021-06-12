package store_test

import (
	"github.com/stretchr/testify/assert"
	"http-rest-api/internal/app/model"
	"http-rest-api/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email: "user@test.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@test.com"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@test.com",
	})

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
