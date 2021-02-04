package model_test

import (
	"testing"

	"github.com/LucasGab/imersao-fullstack-fullcycle/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	name := "Lucas"
	email := "teste@gmail.com"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", "lucas@gmail.com")
	require.NotNil(t, err)

	_, err = model.NewUser("Lucas", "")
	require.NotNil(t, err)
}
