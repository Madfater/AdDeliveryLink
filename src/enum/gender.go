package enum

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Both   Gender = "B"
)

func (g Gender) IsValidEnum() bool {
	switch g {
	case Male, Female, Both:
		return true
	}
	return false
}
