package v2

import (
	"net/http"
	"toscactl/entity"
)

const (
	APIProjects = "Projects"
	APIEmployee = "/employee/1"
)

type Provider struct {
	config *entity.ApplicationConfig
}

func GetProjects() (*http.Response, error) {
	url := "http://10.183.56.17:443/AutomationObjectService/api/"
	req, err := http.NewRequest("GET", url+APIProjects, nil)
	req.Header.Set("X-Tricentis", "OK")
	client := &http.Client{}
	response, err := client.Do(req)
	return response, err
}
func GetEmployee() (*http.Response, error) {
	baseURL := "https://dummy.restapiexample.com/api/v1"
	req, err := http.NewRequest("GET", baseURL+APIEmployee, nil)
	client := &http.Client{}
	response, err := client.Do(req)
	return response, err
}
