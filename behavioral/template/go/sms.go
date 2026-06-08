package main

import "fmt"

type Sms struct {
	Otp
}

func (s *Sms) getRandomOTP(len int) string {
	randomTOP := "1234"
	fmt.Printf("SMS: generating random top %s\n", randomTOP)
	return randomTOP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving opt: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return fmt.Sprintf("SMS OTP for login is %s", otp)
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
