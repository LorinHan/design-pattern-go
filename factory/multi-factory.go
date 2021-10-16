package factory

type ChineseFactory struct {
}

func (cf *ChineseFactory) Create() *chinese {
	return &chinese{
		humanBase{"黄", "中文"},
	}
}

type AmericanFactory struct {
}

func (cf *AmericanFactory) Create() *chinese {
	return &chinese{
		humanBase{"white", "English"},
	}
}
