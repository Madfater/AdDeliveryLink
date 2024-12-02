package enum

type CountryCode string

const (
	CN CountryCode = "CN"
	US CountryCode = "US"
	JP CountryCode = "JP"
	TW CountryCode = "TW"
)

func (c CountryCode) IsValidEnum() bool {
	switch c {
	case CN, US, JP, TW:
		return true
	}
	return false
}
