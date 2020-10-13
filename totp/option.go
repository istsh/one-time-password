package totp

import (
	"errors"
	"fmt"

	"github.com/istsh/one-time-password/otpauth"
)

// Option is used when operates the Time-based One Time Password
type Option struct {
	// period is the seconds that a Time-based One Time Password hash is valid
	// The default value is 30 seconds
	period uint
	// skew is the allowed seconds before or after the current time
	// The default value is 0
	skew uint
	// digits is the number of digits
	// The default value is 6
	digits otpauth.Digits
	// algorithm is the hash function to use in the HMAC operation
	// The default value is SHA1
	algorithm otpauth.Algorithm
}

// SetPeriod sets a period that Time-based One Time Password hash is valid
func (opt *Option) SetPeriod(period uint) error {
	if opt == nil {
		return otpauth.ErrOptionIsNil
	}
	if period == 0 {
		return errors.New("invalid period. please pass greater than 0")
	}

	opt.period = period
	return nil
}

// SetSkew sets the allowed seconds before or after the current time
func (opt *Option) SetSkew(skew uint) error {
	if opt == nil {
		return otpauth.ErrOptionIsNil
	}

	opt.skew = skew
	return nil
}

// SetDigits sets the number of digits
func (opt *Option) SetDigits(d otpauth.Digits) error {
	if opt == nil {
		return otpauth.ErrOptionIsNil
	}
	if !d.Enabled() {
		return fmt.Errorf("invalid digits. please pass %d or %d", otpauth.DigitsSix, otpauth.DigitsEight)
	}

	opt.digits = d
	return nil
}

// SetAlgorithm sets the hash algorithm
func (opt *Option) SetAlgorithm(a otpauth.Algorithm) error {
	if opt == nil {
		return otpauth.ErrOptionIsNil
	}
	if !a.Enabled() {
		return fmt.Errorf("invalid algorithm. please pass any of %d to %d", otpauth.AlgorithmSHA1, otpauth.AlgorithmMD5)
	}

	opt.algorithm = a
	return nil
}

// NewOption generates an option with default values
func NewOption() *Option {
	return &Option{
		period:    otpauth.DefaultPeriod,
		digits:    otpauth.DigitsSix,
		algorithm: otpauth.AlgorithmSHA1,
	}
}