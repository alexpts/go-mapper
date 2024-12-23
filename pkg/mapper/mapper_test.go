package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alexpts/go-mapper/pkg/model"
)

func createMapper[Src any, Dest any]() Mapper[Src, Dest] {
	return Mapper[Src, Dest]{
		TypeManager: NewReflectManager(),
	}
}

// TestSimpleConvertModelToDTO
// Все публичные поля зеркалируются в target модель
func TestSimpleConvertModelToDTO(t *testing.T) {
	mapper := createMapper[model.UserModel, model.UserDto]()

	in := model.NewUserModel("alex", 12)
	actual, err := mapper.Convert(in)
	require.NoError(t, err)

	expected := model.UserDto{
		Age: 12, Name: "alex",
	}
	assert.Equal(t, expected, actual)
}

// TestPartField
// Проверяем, что по дефолту приватные поля в target моделе не заполняются (пропускаются)
func TestPartField(t *testing.T) {
	type DTO struct {
		age  int
		Name string
	}

	mapper := createMapper[model.UserModel, DTO]()

	actual, err := mapper.Convert(
		model.NewUserModel("alex", 12),
	)

	require.NoError(t, err)

	// Default skip private target field
	expected := DTO{age: 0, Name: "alex"}
	assert.Equal(t, expected, actual)
}

// TestReadPrivateProperty
// Проверям чтение приватного поля из ихсодной модели
func TestReadPrivateProperty(t *testing.T) {
	type Model struct {
		Age  int
		name string
	}

	type DTO struct {
		Age  int
		Name string
	}

	mapper := createMapper[Model, DTO]()

	actual, err := mapper.Convert(&Model{name: "alex", Age: 12})

	require.NoError(t, err)

	expected := DTO{Age: 12, Name: "alex"}
	assert.Equal(t, expected, actual)
}
