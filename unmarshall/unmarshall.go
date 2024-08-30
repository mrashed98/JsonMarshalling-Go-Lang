package unmarshall

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	UserResponse struct {
		Results []User
	}
	User struct {
		Gender     string
		Name       UserName
		Location   UserLocation
		Email      string
		Login      UserLogin
		Dob        DateOfBirth
		Registered Registered
		Phone      string
		Cell       string
		Id         UserID
		Picture    UserPicture
		Nat        string
	}
	UserLogin struct {
		Uuid     string
		Username string
		Password string
		Salt     string
		Md5      string
		Sha1     string
		Sha256   string
	}
	DateOfBirth struct {
		Date string
		Age  int
	}
	Registered struct {
		Date string
		Age  int
	}
	UserID struct {
		Name  string
		Value string
	}
	UserPicture struct {
		Large     string
		Medium    string
		Thumbnail string
	}
	UserName struct {
		Title string
		First string
		Last  string
	}
	UserLocation struct {
		Street      UserStreet
		City        string
		State       string
		Country     string
		Postcode    any
		Coordinates LocationCoordinates
		Timezone    Timezone
	}
	UserStreet struct {
		Number int
		Name   string
	}
	LocationCoordinates struct {
		Latitude  string
		Longitude string
	}
	Timezone struct {
		Offset      string
		Description string
	}
)

func TryUnmarshall() {
	const url = "https://randomuser.me/api/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request Failed:", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var users UserResponse

	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error occurred while Unmarshal the response, ERROR:", err.Error())
		return
	}
	for _, user := range users.Results {
		fmt.Printf("User: %+v\n", user)
	}
}
