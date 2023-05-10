package autowirego

import (
	"github.com/gookit/slog"
	"github.com/ohanakogo/ohanakoutilgo"
)

func InjectWithName[T any](obj *T, name string) {
	typeName := ohanakoutilgo.TypeOf[T]().String()
	if name == typeName {
		slog.Errorf("can't inject object with key: *")
		return
	}
	injectByNameBase(obj, name)
}

func Inject[T any](obj *T) {
	injectBase(obj)
}

func ExtractWithName[T any](name string) *T {
	typeName := ohanakoutilgo.TypeOf[T]().String()
	if name == typeName {
		slog.Errorf("can't extract object with key: *")
		return nil
	}
	return extractByNameBase[T](name)
}

func Extract[T any]() *T {
	return extractBase[T]()
}
