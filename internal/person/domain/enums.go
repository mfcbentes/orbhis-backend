package domain

// MaritalStatus representa o estado civil.
type MaritalStatus string

const (
	MaritalSingle      MaritalStatus = "single"
	MaritalMarried     MaritalStatus = "married"
	MaritalDivorced    MaritalStatus = "divorced"
	MaritalWidowed     MaritalStatus = "widowed"
	MaritalSeparated   MaritalStatus = "separated"
	MaritalStableUnion MaritalStatus = "stable_union"
	MaritalNotInformed MaritalStatus = "not_informed"
)

// Race representa cor/raça no padrão IBGE/SUS.
type Race string

const (
	RaceWhite       Race = "white"
	RaceBlack       Race = "black"
	RaceBrown       Race = "brown"
	RaceYellow      Race = "yellow"
	RaceIndigenous  Race = "indigenous"
	RaceNotInformed Race = "not_informed"
)
