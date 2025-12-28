package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}
type DataProcessor interface {
	Process(data interface{}) interface{}
}

func ProcessUserData(processor DataProcessor, userData interface{}) string {
	result := processor.Process(userData)
	output := fmt.Sprintf("%v", result)
	output = strings.Replace(output, "{", "", -1)
	output = strings.Replace(output, "}", "", -1)
	output = strings.Replace(output, " ", "_", -1)
	if output == "" {
		return "no_data"
	}
	return output
}

type UserFormatter struct{}

func (uf UserFormatter) Process(data interface{}) interface{} {
	user := data.(User)
	// Format user data
	var status string
	if user.IsActive == true {
		status = "active"
	} else {
		status = "inactive"
	}
	return fmt.Sprintf("User_%d_%s_%s_%s", user.ID, user.Name, user.Email, status)
}

type NumberProcessor struct{}

func (np NumberProcessor) Process(data interface{}) interface{} {
	switch v := data.(type) {
	case int:
		return v * 2
	case string:
		if num, err := strconv.Atoi(v); err == nil {
			return num * 2
		}
		return 0
	}
	return nil
}
func main() {
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		IsActive: true,
	}
	userFormatter := UserFormatter{}
	result1 := ProcessUserData(userFormatter, user)
	fmt.Println("User result:", result1)
	numberProcessor := NumberProcessor{}
	result2 := ProcessUserData(numberProcessor, 42)
	fmt.Println("Number result:", result2)
	result3 := ProcessUserData(numberProcessor, "21")
	fmt.Println("String number result:", result3)
	result4 := ProcessUserData(numberProcessor, "invalid")
	fmt.Println("Invalid data result:", result4)
}
