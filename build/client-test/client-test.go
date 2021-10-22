package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func testEditSoftware(id int, d, token string) {
	reqBody, err := json.Marshal(map[string]string{
		"id":           fmt.Sprintf("%v", id),
		"name":         "pgadmin",
		"description":  d,
		"year":         "2012",
		"release_date": "2012-01-01",
	})
	if err != nil {
		print(err)
	}

	client := http.Client{}
	req, err := http.NewRequest("POST",
		"http://localhost:4000/v1/admin/editSoftware", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header = http.Header{
		"Host":          []string{"www.domain.com"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{fmt.Sprintf("bearer %v", token)},
	}

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

func testDeleteSoftware(id int, token string) {
	client := http.Client{}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("http://localhost:4000/v1/admin/deletesoftware/%v", id), nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header = http.Header{
		"Host":          []string{"www.domain.com"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{fmt.Sprintf("bearer %v", token)},
	}

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

type Items struct {
	Software []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"software"`
}

func GetPGAdminID() int {
	resp, err := http.Get("http://localhost:4000/v1/software")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	var s Items
	json.Unmarshal(body, &s)

	for _, item := range s.Software {
		if item.Name == "pgadmin" {
			return item.ID
		}
	}
	return 0
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

func LoginAdminUser() string {
	reqBody, err := json.Marshal(map[string]string{
		"email":    "admin@domain.com",
		"password": "password",
	})
	if err != nil {
		print(err)
	}

	resp, err := http.Post("http://localhost:4000/v1/login",
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	var jwt JWT
	json.Unmarshal(body, &jwt)

	return jwt.Token

}

func SignUpNewUser(token string) {
	reqBody, err := json.Marshal(map[string]string{
		"email":    "user@domain.com",
		"password": "12345",
	})
	if err != nil {
		print(err)
	}

	client := http.Client{}
	req, err := http.NewRequest("POST",
		"http://localhost:4000/v1/admin/signup", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header = http.Header{
		"Host":          []string{"www.domain.com"},
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{fmt.Sprintf("bearer %v", token)},
	}

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

func main() {
	//fmt.Println(LoginAdminUser())
	//SignUpNewUser("bearer bad-token-example")
	//SignUpNewUser(LoginAdminUser())

	id := GetPGAdminID()
	for id == 0 {
		testEditSoftware(0, "SQL management tool", LoginAdminUser())
		id = GetPGAdminID()
	}
	fmt.Printf("pgadmin id is: %v\n", id)
	testEditSoftware(id, "PostgreSQL management tool", LoginAdminUser())
	testEditSoftware(id, "wrong token", "bearer bad-token-example")

	/*
		if GetPGAdminID() != 0 {
			testDeleteSoftware(GetPGAdminID(), LoginAdminUser())
		}
	*/
}
