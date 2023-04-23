package autowirego

import "github.com/gookit/slog"

func InjectWithName[T any](obj *T, name string) {
	if name == "*" {
		slog.Errorf("can't inject object with key: *")
		return
	}
	injectByNameBase(obj, name)
}

func Inject[T any](obj *T) {
	injectBase(obj)
}

func ExtractWithName[T any](name string) *T {
	if name == "*" {
		slog.Errorf("can't extract object with key: *")
		return nil
	}
	return extractByNameBase[T](name)
}

func Extract[T any]() *T {
	return extractBase[T]()
}
