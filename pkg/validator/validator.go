package validator

type Validator struct {
	// maybe it will be contain something in fututre
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) IsFileHasTxtExtension(filename string) bool {
	if len(filename) <= 4 {
		return false
	}
	return filename[len(filename)-4:] == ".txt"
}

func (v *Validator) IsArrayContainsString(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

func (v *Validator) IsTextAscii(text string) bool {
	for _, ch := range text {
		if (int(ch) < 32 || int(ch) > 126) && int(ch) != 10 && int(ch) != 9 {
			return false
		}
	}
	return true
}
