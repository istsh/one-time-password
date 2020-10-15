package recovery

func (opt *Option) Letters() string {
	if opt == nil {
		return ""
	}

	return opt.letters
}

func (opt *Option) Length() uint {
	if opt == nil {
		return 0
	}

	return opt.length
}

func (opt *Option) Count() uint {
	if opt == nil {
		return 0
	}

	return opt.count
}

func DefaultOption() *Option {
	return &Option{
		letters: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		length:  8,
		count:   8,
	}
}