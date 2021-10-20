package factory

func CreateHuman(t string) humanI {
	switch t {
	case "chinese":
		return &chinese{humanBase{"黄", "中文"}}
	case "american":
		return &american{humanBase{"white", "English"}}
	default:
		return nil
	}
}
