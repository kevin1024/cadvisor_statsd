package cadvisor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"log"
)

type containers struct {
	Name          string `json:"name"`
	Subcontainers []struct {
		Name string `json:"name"`
	} `json:"subcontainers"`
}

func (c *containers) Names() []string {
	output := make([]string, 0, len(c.Subcontainers))
	for _, subcontainer := range c.Subcontainers {
		output = append(output, subcontainer.Name)
	}
	return output
}

func decodeContainerJson(jsonString io.Reader) (*containers, error) {
	parsedJson := new(containers)
	if err := json.NewDecoder(jsonString).Decode(parsedJson); err != nil {
		return nil, err
	}
	return parsedJson, nil
}

func getContainerList(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	c, err := decodeContainerJson(resp.Body)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return c.Names(), err
}

func GetAllSubcontainers() ([]string, error) {
	containers, err := getContainerList("http://192.168.59.103:8080/api/v1.1/containers")
	if err != nil {
		fmt.Printf("THIS IS ANNOYING: %v", err)
	}

	output := make([]string, 0, len(containers))

	for _, name := range containers {
	        url := fmt.Sprintf("http://192.168.59.103:8080/api/v1.1/containers%v/", name)
		log.Printf("----> fetching %v <---\n", url)
		subcontainers, _ := getContainerList(url)
		log.Printf("subcontainers: %v\n", subcontainers)
		output = append(output, subcontainers...)
	}
	return output, err
}



//func GetStatistics(containerID string) (
