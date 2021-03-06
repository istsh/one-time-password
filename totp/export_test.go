package totp

import "github.com/butterv/one-time-password/otpauth"

func (opt *Option) Period() uint {
	if opt == nil {
		return 0
	}

	return opt.period
}

func (opt *Option) Skew() uint {
	if opt == nil {
		return 0
	}

	return opt.skew
}

func (opt *Option) Digits() otpauth.Digits {
	if opt == nil {
		return 0
	}

	return opt.digits
}

func (opt *Option) Algorithm() otpauth.Algorithm {
	if opt == nil {
		return 0
	}

	return opt.algorithm
}

func DefaultOption() *Option {
	return &Option{
		period:    30,
		skew:      1,
		digits:    6,
		algorithm: 0,
	}
}
