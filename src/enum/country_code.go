package enum

type CountryCode string

const (
	CA CountryCode = "CA"
	US CountryCode = "US"
	JP CountryCode = "JP"
	TW CountryCode = "TW"
)

func (c CountryCode) IsValidEnum() bool {
	switch c {
	case CA, US, JP, TW:
		return true
	}
	return false
}
