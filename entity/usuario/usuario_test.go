package usuario

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohanalexander/deezefy-music/entity"
)

func TestNewUsuario(t *testing.T) {

	t.Run("Usuario criado com sucesso", func(t *testing.T) {
		u, err := NewUsuario("steve.jobs@apple.com", "new_password", "2006-01-02")
		assert.Nil(t, err)
		assert.Equal(t, u.Email, "steve.jobs@apple.com")
		assert.NotEqual(t, u.Password, "new_password")
	})

}

func TestValidatePassword(t *testing.T) {

	t.Run("Senha correta", func(t *testing.T) {
		u, _ := NewUsuario("steve.jobs@apple.com", "new_password", "2006-01-02")
		err := u.ValidatePassword("new_password")
		assert.Nil(t, err)
	})

	t.Run("Senha incorreta", func(t *testing.T) {
		u, _ := NewUsuario("steve.jobs@apple.com", "new_password", "2006-01-02")
		err := u.ValidatePassword("wrong_password")
		assert.NotNil(t, err)
	})

}

func TestUsuario_Validate(t *testing.T) {

	type test struct {
		name     string
		email    string
		password string
		birthday string
		want     error
	}

	tests := []test{
		{
			name:     "Campos válidos",
			email:    "steve.jobs@apple.com",
			password: "new_password",
			birthday: "2006-01-02",
			want:     nil,
		},
		{
			name:     "Email inválido (user@company.com)",
			email:    "",
			password: "new_password",
			birthday: "2006-01-02",
			want:     entity.ErrInvalidEntity,
		},
		{
			name:     "Password inválida (12345678)",
			email:    "steve.jobs@apple.com",
			password: "",
			birthday: "2006-01-02",
			want:     entity.ErrInvalidEntity,
		},
		{
			name:     "Birthday inválido (2006-01-02)",
			email:    "steve.jobs@apple.com",
			password: "new_password",
			birthday: "2006/01/02",
			want:     entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewUsuario(tc.email, tc.password, tc.birthday)
			assert.Equal(t, err, tc.want)
		})
	}

}
