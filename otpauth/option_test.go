package otpauth_test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"reflect"
	"testing"

	"github.com/butterv/one-time-password/otpauth"
)

func TestHost_Enabled(t *testing.T) {
	tests := []struct {
		in   otpauth.Host
		want bool
	}{
		{in: 0, want: true},
		{in: 1, want: true},
		{in: -1, want: false},
		{in: 2, want: false},
	}

	for _, tt := range tests {
		got := otpauth.ExportHostEnabled(tt.in)
		if got != tt.want {
			t.Errorf("ExportHostEnabled(%d)=%v; want %v", tt.in, got, tt.want)
		}
	}
}

func TestHost_Name(t *testing.T) {
	tests := []struct {
		in   otpauth.Host
		want string
	}{
		{in: 0, want: "hotp"},
		{in: 1, want: "totp"},
	}

	for _, tt := range tests {
		got := otpauth.ExportHostName(tt.in)
		if got != tt.want {
			t.Errorf("ExportHostName(%d)=%s; want %s", tt.in, got, tt.want)
		}
	}
}

func TestHost_Name_Panic(t *testing.T) {
	defer func() {
		want := "invalid host"

		err := recover()
		if err == nil {
			t.Fatalf("no panic is detected. panic should be called. want=%v", want)
		}
		if !reflect.DeepEqual(err, want) {
			t.Errorf("unexpected error is returned; got=%#v, want=%v", err, want)
		}
	}()

	in := otpauth.Host(2)
	_ = otpauth.ExportHostName(in)
}

func TestDigits_Enabled(t *testing.T) {
	tests := []struct {
		in   otpauth.Digits
		want bool
	}{
		{in: otpauth.DigitsSix, want: true},
		{in: otpauth.DigitsEight, want: true},
		{in: 0, want: false},
	}

	for _, tt := range tests {
		got := tt.in.Enabled()
		if got != tt.want {
			t.Errorf("Enabled()=%v; want %v, receiver %#v", got, tt.want, tt.in)
		}
	}
}

func TestDigits_Length(t *testing.T) {
	tests := []struct {
		in   otpauth.Digits
		want int
	}{
		{in: otpauth.DigitsSix, want: 6},
		{in: otpauth.DigitsEight, want: 8},
	}

	for _, tt := range tests {
		got := tt.in.Length()
		if got != tt.want {
			t.Errorf("Length()=%d; want %d, receiver %#v", got, tt.want, tt.in)
		}
	}
}

func TestDigits_Format(t *testing.T) {
	tests := []struct {
		in   otpauth.Digits
		want string
	}{
		{in: otpauth.DigitsSix, want: "001234"},
		{in: otpauth.DigitsEight, want: "00001234"},
	}

	in := int32(1234)
	for _, tt := range tests {
		got := tt.in.Format(in)
		if got != tt.want {
			t.Errorf("Format(%d)=%s; want %s, receiver %#v", in, got, tt.want, tt.in)
		}
	}
}

func TestDigits_Format_Panic(t *testing.T) {
	defer func() {
		want := "invalid digits"

		err := recover()
		if err == nil {
			t.Fatalf("no panic is detected. panic should be called. want=%v", want)
		}
		if !reflect.DeepEqual(err, want) {
			t.Errorf("unexpected error is returned; got=%#v, want=%v", err, want)
		}
	}()

	in := int32(1234)
	d := otpauth.Digits(10)
	_ = d.Format(in)
}

func TestAlgorithm_Enabled(t *testing.T) {
	tests := []struct {
		in   otpauth.Algorithm
		want bool
	}{
		{in: otpauth.AlgorithmSHA1, want: true},
		{in: otpauth.AlgorithmSHA256, want: true},
		{in: otpauth.AlgorithmSHA512, want: true},
		{in: otpauth.AlgorithmMD5, want: true},
		{in: -1, want: false},
		{in: 4, want: false},
	}

	for _, tt := range tests {
		got := tt.in.Enabled()
		if got != tt.want {
			t.Errorf("Enabled()=%v; want %v, receiver %#v", got, tt.want, tt.in)
		}
	}
}

func TestAlgorithm_Name(t *testing.T) {
	tests := []struct {
		in   otpauth.Algorithm
		want string
	}{
		{in: otpauth.AlgorithmSHA1, want: "SHA1"},
		{in: otpauth.AlgorithmSHA256, want: "SHA256"},
		{in: otpauth.AlgorithmSHA512, want: "SHA512"},
		{in: otpauth.AlgorithmMD5, want: "MD5"},
	}

	for _, tt := range tests {
		got := otpauth.ExportAlgorithmName(tt.in)
		if got != tt.want {
			t.Errorf("ExportAlgorithmName(%d)=%s; want %s", tt.in, got, tt.want)
		}
	}
}

func TestAlgorithm_Name_Panic(t *testing.T) {
	defer func() {
		want := "invalid algorithm"

		err := recover()
		if err == nil {
			t.Fatalf("no panic is detected. panic should be called. want=%v", want)
		}
		if !reflect.DeepEqual(err, want) {
			t.Errorf("unexpected error is returned; got=%#v, want=%v", err, want)
		}
	}()

	a := otpauth.Algorithm(4)
	_ = otpauth.ExportAlgorithmName(a)
}

func TestAlgorithm_Hash(t *testing.T) {
	tests := []struct {
		in   otpauth.Algorithm
		want hash.Hash
	}{
		{in: otpauth.AlgorithmSHA1, want: sha1.New()},
		{in: otpauth.AlgorithmSHA256, want: sha256.New()},
		{in: otpauth.AlgorithmSHA512, want: sha512.New()},
		{in: otpauth.AlgorithmMD5, want: md5.New()},
	}

	for _, tt := range tests {
		got := tt.in.Hash()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Hash()=%v; want %v, receiver %d", got, tt.want, tt.in)
		}
	}
}

func TestAlgorithm_Hash_Panic(t *testing.T) {
	defer func() {
		want := "invalid algorithm"

		err := recover()
		if err == nil {
			t.Fatalf("no panic is detected. panic should be called. want=%v", want)
		}
		if !reflect.DeepEqual(err, want) {
			t.Errorf("unexpected error is returned; got=%#v, want=%v", err, want)
		}
	}()

	a := otpauth.Algorithm(4)
	_ = a.Hash()
}

func TestOption_SetPeriod(t *testing.T) {
	want := uint(30)

	period := uint(30)
	o := &otpauth.Option{}
	err := o.SetPeriod(period)
	if err != nil {
		t.Fatalf("SetPeriod(%d)=%#v; want nil, receiver %#v", period, err, o)
	}
	if got := o.Period(); got != want {
		t.Errorf("period: got %d, want %d, receiver %#v", got, want, o)
	}
}

func TestOption_SetPeriod_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOtpAuthOptionIsNil

	period := uint(30)
	var o *otpauth.Option
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
	o := &otpauth.Option{}
	err := o.SetPeriod(period)
	if err == nil {
		t.Fatalf("SetPeriod(%d)=nil; want %v, receiver nil", period, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetPeriod(%d)=%#v; want %v, receiver nil", period, err, wantErr)
	}
}

func TestOption_SetSecretSize(t *testing.T) {
	want := uint(20)

	secretSize := uint(20)
	o := &otpauth.Option{}
	err := o.SetSecretSize(secretSize)
	if err != nil {
		t.Fatalf("SetSecretSize(%d)=%#v; want nil, receiver %#v", secretSize, err, o)
	}
	if got := o.SecretSize(); got != want {
		t.Errorf("secretSize: got %d, want %d, receiver %#v", got, want, o)
	}
}

func TestOption_SetSecretSize_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOtpAuthOptionIsNil

	secretSize := uint(20)
	var o *otpauth.Option
	err := o.SetSecretSize(secretSize)
	if err == nil {
		t.Fatalf("SetSecretSize(%d)=nil; want %v, receiver nil", secretSize, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetSecretSize(%d)=%#v; want %v, receiver nil", secretSize, err, wantErr)
	}
}

func TestOption_SetSecretSize_InvalidSecretSize(t *testing.T) {
	wantErr := errors.New("invalid secretSize. please pass greater than 0")

	secretSize := uint(0)
	o := &otpauth.Option{}
	err := o.SetSecretSize(secretSize)
	if err == nil {
		t.Fatalf("SetSecretSize(%d)=nil; want %v, receiver nil", secretSize, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetSecretSize(%d)=%#v; want %v, receiver nil", secretSize, err, wantErr)
	}
}

func TestOption_Secret(t *testing.T) {
	want := "JXVF3ZJE2U52WP3B77D77VQJ3J3VYDUZ"

	secret := "JXVF3ZJE2U52WP3B77D77VQJ3J3VYDUZ"
	o := &otpauth.Option{}
	_ = o.SetSecret(secret)
	got := o.Secret()
	if got != want {
		t.Errorf("Secret()=%s; want %s, receiver %#v", got, want, o)
	}
}

func TestOption_Secret_Empty(t *testing.T) {
	want := ""

	var o *otpauth.Option
	got := o.Secret()
	if got != want {
		t.Errorf("Secret()=%s; want %s, receiver %#v", got, want, o)
	}
}

func TestOption_SetSecret(t *testing.T) {
	want := "JXVF3ZJE2U52WP3B77D77VQJ3J3VYDUZ"

	secret := "JXVF3ZJE2U52WP3B77D77VQJ3J3VYDUZ"
	o := &otpauth.Option{}
	err := o.SetSecret(secret)
	if err != nil {
		t.Fatalf("SetSecret(%s)=%#v; want nil, receiver %#v", secret, err, o)
	}
	if got := o.Secret(); got != want {
		t.Errorf("secret: got %s, want %s, receiver %#v", got, want, o)
	}
}

func TestOption_SetSecret_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOtpAuthOptionIsNil

	secret := "JXVF3ZJE2U52WP3B77D77VQJ3J3VYDUZ"
	var o *otpauth.Option
	err := o.SetSecret(secret)
	if err == nil {
		t.Fatalf("SetSecret(%s)=nil; want %v, receiver nil", secret, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetSecret(%s)=%#v; want %v, receiver nil", secret, err, wantErr)
	}
}

func TestOption_SetDigits(t *testing.T) {
	want := otpauth.DigitsSix

	digits := otpauth.DigitsSix
	o := &otpauth.Option{}
	err := o.SetDigits(digits)
	if err != nil {
		t.Fatalf("SetDigits(%d)=%#v; want nil, receiver %#v", digits, err, o)
	}
	if got := o.Digits(); got != want {
		t.Errorf("digits: got %d, want %d, receiver %#v", got, want, o)
	}
}

func TestOption_SetDigits_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOtpAuthOptionIsNil

	digits := otpauth.DigitsSix
	var o *otpauth.Option
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
	o := &otpauth.Option{}
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
	o := &otpauth.Option{}
	err := o.SetAlgorithm(a)
	if err != nil {
		t.Fatalf("SetAlgorithm(%d)=%#v; want nil, receiver %#v", a, err, o)
	}
	if got := o.Algorithm(); got != want {
		t.Errorf("algorithm: got %d, want %d, receiver %#v", got, want, o)
	}
}

func TestOption_SetAlgorithm_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOtpAuthOptionIsNil

	a := otpauth.AlgorithmSHA1
	var o *otpauth.Option
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
	o := &otpauth.Option{}
	err := o.SetAlgorithm(a)
	if err == nil {
		t.Fatalf("SetAlgorithm(%d)=nil; want %v, receiver nil", a, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetAlgorithm(%d)=%#v; want %v, receiver nil", a, err, wantErr)
	}
}

func TestOption_SetIconURL(t *testing.T) {
	want := "TEST_ICON_URL"

	url := "TEST_ICON_URL"
	o := &otpauth.Option{}
	err := o.SetIconURL(url)
	if err != nil {
		t.Fatalf("SetIconURL(%s)=%#v; want nil, receiver %#v", url, err, o)
	}
	if got := o.IconURL(); got != want {
		t.Errorf("icon url: got %s, want %s, receiver %#v", got, want, o)
	}
}

func TestOption_SetIconURL_ErrOptionIsNil(t *testing.T) {
	wantErr := otpauth.ErrOtpAuthOptionIsNil

	url := "TEST_ICON_URL"
	var o *otpauth.Option
	err := o.SetIconURL(url)
	if err == nil {
		t.Fatalf("SetIconURL(%s)=nil; want %v, receiver nil", url, wantErr)
	}
	if err.Error() != wantErr.Error() {
		t.Errorf("SetIconURL(%s)=%#v; want %v, receiver nil", url, err, wantErr)
	}
}

func TestNewOption(t *testing.T) {
	want := otpauth.DefaultOption()

	got, err := otpauth.NewOption()
	if err != nil {
		t.Fatalf("NewOption()=_, %#v; want nil", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewOption()=%#v, _; want %v", got, want)
	}
}
