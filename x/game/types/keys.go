package types

const (
	ModuleName = "game"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

var (
	KeyOddsAmount  = []byte{0x11}
	KeyEvensAmount = []byte{0x12}
)
