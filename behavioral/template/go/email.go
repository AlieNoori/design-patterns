package main

import "fmt"

type Email struct {
	Otp
}

func (e *Email) getRandomOTP(len int) string {
	randomTOP := "1234"
	fmt.Printf("EMAIL: generating random top %s\n", randomTOP)
	return randomTOP
}

func (e *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving opt: %s to cache\n", otp)
}

func (e *Email) getMessage(otp string) string {
	return fmt.Sprintf("EMAIL OTP for login is %s", otp)
}

func (e *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
