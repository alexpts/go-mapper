package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alexpts/go-mapper/pkg/model"
)

func TestSimpleConvertModelToDTO(t *testing.T) {
	mapper := Mapper[model.UserModel, model.UserDto]{
		TypeManager: NewReflectManager(),
	}

	in := model.NewUserModel("alex", 12)
	actual, err := mapper.Convert(in)
	require.NoError(t, err)

	expected := model.UserDto{
		Age: 12, Name: "alex",
	}
	assert.Equal(t, expected, actual)
}
