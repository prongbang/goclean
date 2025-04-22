package swagger

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewRouter,
)
