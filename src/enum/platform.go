package enum

type Platform string

const (
	Android Platform = "android"
	IOS     Platform = "ios"
	Web     Platform = "web"
)

func (p Platform) IsValidEnum() bool {
	switch p {
	case Android, IOS, Web:
		return true
	}
	return false
}
