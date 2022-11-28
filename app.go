package main

import (
	"fmt"
	"regexp"
	"strings"
)

const registraionTrials uint = 3
const newRegistrationLimit uint = 10

var registrationName = "counchSpace"
var remainingRegistrationTrials = registraionTrials
var remainingNewRegistraitions = newRegistrationLimit
var registrations = make([]UserData, 0)

// new type of user
type UserData struct {
	firstName    string
	lastName     string
	email        string
	mobileNumber string
	username     string
	password     string
}

// main function
func main() {
	welcomeRegister()
	for {
		firstName, lastName, email, mobileNumber, username, password := getRegisterInput()
		isValidEmail, isValidMobileNumber, isValidusername, isValidPassowrd := validateRegisterInput(email, mobileNumber, username, password)
		if isValidEmail && isValidMobileNumber && isValidusername && isValidPassowrd {
			registering(firstName, lastName, email, mobileNumber, username, password)
			remainingRegistrationTrials = registraionTrials
			usernames := getusernames()
			fmt.Printf(" The user names of registrions are : %v\n", usernames)
			if remainingNewRegistraitions == 0 {
				fmt.Println(" We are so sorry. You can not register new accounts")
				break
			}
		} else {
			remainingRegistrationTrials = remainingRegistrationTrials - 1

		}
		if remainingRegistrationTrials == 0 {
			fmt.Println(" We are so sorry. You run out of registration trial and can not use the application anymore")
			break

		}
	}
}

func welcomeRegister() {
	fmt.Printf("Welcome to %v's regitration page", registrationName)
	fmt.Printf("There are %v possible registrations in total and  %v remaining registrations", newRegistrationLimit, remainingNewRegistraitions)
	fmt.Println("Let's start your registration: ")

}

// in put the information of new register
func getRegisterInput() (string, string, string, string, string, string) {
	var firstName string
	var lastName string
	var email string
	var mobileNumber string
	var username string
	var password string
	fmt.Println("Ener your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Ener your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Ener your email:")
	fmt.Scan(&email)

	fmt.Println("Ener your mobile number:")
	fmt.Scan(&mobileNumber)

	fmt.Println("Ener your username:")
	fmt.Scan(&username)

	fmt.Println("Ener your password:")
	fmt.Scan(&password)
	return firstName, lastName, email, mobileNumber, username, password
}

// validate if those input information valid or not
func validateRegisterInput(email string, mobileNumber string, username string, password string) (bool, bool, bool, bool) {
	isValidEmail := strings.Contains(email, "@")
	isValidMobileNumber := len(mobileNumber) >= 9 && len(mobileNumber) <= 11
	isValidusername := len(username) >= 8

	isValidPasswordLength := len(password) >= 8
	isValidPasswordUpperCase := regexp.MustCompile("[A-Z]").FindString(password) != ""
	isValidPasswordLowerCase := regexp.MustCompile("[a-z]").FindString(password) != ""
	isValidPasswordDigit := regexp.MustCompile("[0-9]").FindString(password) != ""
	isValidPassword := isValidPasswordLength && isValidPasswordUpperCase && isValidPasswordLowerCase && isValidPasswordDigit

	if !isValidEmail {
		fmt.Println("Your email must contain @")

	}
	if !isValidMobileNumber {
		fmt.Println("Your mobile number length must range from 9 to 11")
	}

	if !isValidusername {
		fmt.Println("Your username length must be at least 8")
	}

	if !isValidPasswordLength {
		fmt.Println("Your password length must be at least 8")
	}
	if !isValidPasswordUpperCase {
		fmt.Println("Your password must contain at least one upper case letter")
	}
	if !isValidPasswordLowerCase {
		fmt.Println("Your password must contain at least one lower case letter")
	}
	if !isValidPasswordDigit {
		fmt.Println("Your password must contain at least one digit")
	}

	return isValidEmail, isValidMobileNumber, isValidusername, isValidPassword
}

// If those input information are valid, we would store them into registration data
func registering(firstName string, lastName string, email string, mobileNumber string, username string, password string) {
	remainingNewRegistraitions = remainingNewRegistraitions - 1

	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		mobileNumber: mobileNumber,
		username:     username,
		password:     password,
	}
	registrations = append(registrations, userData)

	fmt.Printf("Welcome %v %v to %v, your username is %v\n ", firstName, lastName, registrationName, userData.username)
	fmt.Printf("%v new registration(s) left for %v\n", remainingNewRegistraitions, registrationName)

}

// retrieve all the registered usernames so far
func getusernames() []string {
	usernames := []string{}
	for _, registration := range registrations {
		usernames = append(usernames, registration.username)
	}
	return usernames
}
