package autowirego

import (
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/slog"
)

var hashingTable = make(map[string]any)

func injectByNameBase[T any](obj *T, name string) {
	hashingTable[name] = obj
}

func injectBase[T any](obj *T) {
	injectByNameBase(obj, "*")
}

func extractByNameBase[T any](name string) *T {
	typeName := reflects.TypeOf((*T)(nil)).String()
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
	return extractByNameBase[T]("*")
}
