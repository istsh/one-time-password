package totp

import (
	"math"
	"time"

	"github.com/butterv/one-time-password/hotp"
)

// Validate validates a Time-based One Time Password with using default value of option
func Validate(passcode, secret string, t time.Time) (bool, error) {
	opt := NewOption()
	return ValidateWithOption(passcode, secret, t, opt)
}

// ValidateWithOption validates a Time-based One Time Password
// This function can pass custom value of option
// See: https://tools.ietf.org/html/rfc6238#section-4.2
func ValidateWithOption(passcode, secret string, t time.Time, opt *Option) (bool, error) {
	hotpOpt := hotp.NewOption()
	_ = hotpOpt.SetDigits(opt.digits)
	_ = hotpOpt.SetAlgorithm(opt.algorithm)

	c := uint64(math.Floor(float64(t.Unix()) / float64(opt.period)))

	var cs []uint64
	cs = append(cs, c)

	for i := uint64(1); i <= uint64(opt.skew); i++ {
		cs = append(cs, c+i)
		cs = append(cs, c-i)
	}

	for _, c := range cs {
		ok, err := hotp.ValidateWithOption(passcode, secret, c, hotpOpt)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	return false, nil
}
