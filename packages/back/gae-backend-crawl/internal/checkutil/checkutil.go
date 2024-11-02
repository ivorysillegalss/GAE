package checkutil

func CheckString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
