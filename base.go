package autowirego

import (
	"github.com/gookit/slog"
	"github.com/ohanakogo/ohanakoutilgo"
)

var hashingTable = make(map[string]any)

func injectByNameBase[T any](obj *T, name string) {
	hashingTable[name] = obj
}

func injectBase[T any](obj *T) {
	typeName := ohanakoutilgo.TypeOf[T]().String()
	injectByNameBase(obj, typeName)
}

func extractByNameBase[T any](name string) *T {
	typeName := ohanakoutilgo.TypeOf[T]().String()
	if hashingTable[name] != nil {
		val, ok := hashingTable[name].(*T)
		if !ok {
			slog.Errorf("can't extract object with key \"%s\" by type \"%s\"", name, typeName)
			return nil
		}
		return val
	}
	return nil
}

func extractBase[T any]() *T {
	typeName := ohanakoutilgo.TypeOf[T]().String()
	return extractByNameBase[T](typeName)
}
