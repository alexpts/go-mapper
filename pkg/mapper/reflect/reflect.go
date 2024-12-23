package reflect

import "reflect"

// StructType Обертка над нативным reflect.Type
// @not-tread-safe
type StructType struct {
	reflect.Type
	Fields map[string]reflect.StructField // @todo кеш (bench сравнить по получению из числового индекса)
}

func NewStructType(model any) StructType {
	native := reflect.TypeOf(model)
	if native.Kind() == reflect.Ptr {
		native = native.Elem()
	}

	if native.Kind() != reflect.Struct {
		panic("model is not ~struct") // @todo error
	}

	return StructType{
		Type:   native,
		Fields: make(map[string]reflect.StructField),
	}
}

func (rt *StructType) FillFields() bool {
	if len(rt.Fields) != 0 {
		return false
	}

	count := rt.Type.NumField()
	for i := 0; i < count; i++ {
		field := rt.Type.Field(i)
		rt.Fields[field.Name] = field
	}

	return true
}

func (rt *StructType) GetFullName() string {
	return rt.PkgPath() + "." + rt.Name()
}

func (rt *StructType) FieldByName(name string) (reflect.StructField, bool) {
	field, isExists := rt.Fields[name]

	if !isExists {
		field, isExists = rt.Type.FieldByName(name)
		if isExists {
			rt.Fields[name] = field
		}
	}

	return field, isExists
}
