package mapper

import (
	nativeReflect "reflect"

	"github.com/alexpts/go-mapper/pkg/mapper/reflect"
)

// Mapper
// @todo Менеджер нужен сверху будет, который по паре может вернут generic из пула
type Mapper[Src any, Dest any] struct {
	TypeManager *ReflectManager
}

type IMapper[Src comparable, Dest comparable] interface {
	Convert(model *Src) (Dest, error)
}

// Convert поиграться бенчмарк тестами указатель или нет
func (m *Mapper[Src, Dest]) Convert(model *Src) (Dest, error) {
	dest := new(Dest)

	srcReflectType := m.TypeManager.Add(model)
	destReflectType := m.TypeManager.Add(dest)

	srcReflectValue := nativeReflect.Indirect(nativeReflect.ValueOf(model))
	destReflectValue := nativeReflect.Indirect(nativeReflect.ValueOf(dest))

	for name, srcField := range srcReflectType.Fields {
		// @todo стратегия разная
		destField, isOk := destReflectType.Fields[name]
		_ = destField
		if !isOk {
			continue
		}

		srcFieldValue := srcReflectValue.FieldByIndex(srcField.Index)
		destFieldValue := destReflectValue.FieldByIndex(destField.Index)

		destFieldValue.Set(srcFieldValue)
		_ = 1
		//srcValueField := srcValue.Field(1)
		//fieldValue := srcValue.FieldByName(name).Ind

		// поле есть в обоих моделятх, нужно из model скопировать значение в dest
	}

	return *dest, nil
}

// ReflectManager - in memory
type ReflectManager struct {
	types map[string]reflect.StructType
}

func NewReflectManager() *ReflectManager {
	return &ReflectManager{types: make(map[string]reflect.StructType)}
}

// Add - @todo Бенчмарк тексты без кеша в in memory / FieldByName / vs custom map
func (rm *ReflectManager) Add(model any) reflect.StructType {
	structType := reflect.NewStructType(model)
	typeFullName := structType.GetFullName()

	// @todo потокобезопасность добавить
	val, isExists := rm.types[typeFullName]
	if isExists {
		return val
	}

	rm.types[typeFullName] = structType
	_ = structType.FillFields()

	return structType
}
