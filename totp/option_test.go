package totp_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/istsh/one-time-password/otpauth"
	"github.com/istsh/one-time-password/totp"
)

func TestOption_SetPeriod(t *testing.T) {
	want := uint(30)

	period := uint(30)
	o := &totp.Option{}
	err := o.SetPeriod(period)
	if err != nil {
		t.Fatalf("SetPeriod(%d)=%#v; want nil, receiver %#v", period, err, o)
	}
	if got := o.Period(); got != want {
		t.Errorf("period: got %d, want %d, receiver %#v", got, want, o)
	}
}

func TestOption_SetPeriod_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOptionIsNil

	period := uint(30)
	var o *totp.Option
	err := o.SetPeriod(period)
	if err == nil {
		t.Fatalf("SetPeriod(%d)=nil; want %v, receiver nil", period, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetPeriod(%d)=%#v; want %v, receiver nil", period, err, wantErr)
	}
}

func TestOption_SetPeriod_InvalidPeriod(t *testing.T) {
	wantErr := errors.New("invalid period. please pass greater than 0")

	period := uint(0)
	o := &totp.Option{}
	err := o.SetPeriod(period)
	if err == nil {
		t.Fatalf("SetPeriod(%d)=nil; want %v, receiver nil", period, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetPeriod(%d)=%#v; want %v, receiver nil", period, err, wantErr)
	}
}

func TestOption_SetDigits(t *testing.T) {
	want := otpauth.DigitsSix

	digits := otpauth.DigitsSix
	o := &totp.Option{}
	err := o.SetDigits(digits)
	if err != nil {
		t.Fatalf("SetDigits(%d)=%#v; want nil, receiver %#v", digits, err, o)
	}
	if got := o.Digits(); got != want {
		t.Errorf("digits: got %d, want %d, receiver %#v", got, want, o)
	}
}

func TestOption_SetDigits_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOptionIsNil

	digits := otpauth.DigitsSix
	var o *totp.Option
	err := o.SetDigits(digits)
	if err == nil {
		t.Fatalf("SetDigits(%d)=nil; want %v, receiver nil", digits, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetDigits(%d)=%#v; want %v, receiver nil", digits, err, wantErr)
	}
}

func TestOption_SetDigits_InvalidDigits(t *testing.T) {
	wantErr := errors.New("invalid digits. please pass 6 or 8")

	digits := otpauth.Digits(0)
	o := &totp.Option{}
	err := o.SetDigits(digits)
	if err == nil {
		t.Fatalf("SetDigits(%d)=nil; want %v, receiver nil", digits, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetDigits(%d)=%#v; want %v, receiver nil", digits, err, wantErr)
	}
}

func TestOption_SetAlgorithm(t *testing.T) {
	want := otpauth.AlgorithmSHA1

	a := otpauth.AlgorithmSHA1
	o := &totp.Option{}
	err := o.SetAlgorithm(a)
	if err != nil {
		t.Fatalf("SetAlgorithm(%d)=%#v; want nil, receiver %#v", a, err, o)
	}
	if got := o.Algorithm(); got != want {
		t.Errorf("algorithm: got %d, want %d, receiver %#v", got, want, o)
	}
}

func TestOption_SetAlgorithm_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOptionIsNil

	a := otpauth.AlgorithmSHA1
	var o *totp.Option
	err := o.SetAlgorithm(a)
	if err == nil {
		t.Fatalf("SetAlgorithm(%d)=nil; want %v, receiver nil", a, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetAlgorithm(%d)=%#v; want %v, receiver nil", a, err, wantErr)
	}
}

func TestOption_SetAlgorithm_InvalidDigits(t *testing.T) {
	wantErr := errors.New("invalid algorithm. please pass any of 0 to 3")

	a := otpauth.Algorithm(4)
	o := &totp.Option{}
	err := o.SetAlgorithm(a)
	if err == nil {
		t.Fatalf("SetAlgorithm(%d)=nil; want %v, receiver nil", a, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetAlgorithm(%d)=%#v; want %v, receiver nil", a, err, wantErr)
	}
}

func TestNewOption(t *testing.T) {
	want := totp.DefaultOption()

	got := totp.NewOption()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewOption()=%#v; want %v", got, want)
	}
}
