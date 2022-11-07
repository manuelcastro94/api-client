package f1_api_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type F1Client interface {
	Create(client *http.Client)
}
type Result map[string]interface{}

type ErgastF1Client struct {
	httpClient *http.Client
	query      string
}

func Create(client *http.Client) ErgastF1Client {
	c := ErgastF1Client{
		httpClient: client,
		query:      API_URL,
	}
	return c
}

func (c ErgastF1Client) Query(format string) Result {
	c.query = getFormat(c.query, format)
	req, err := http.NewRequest("GET", c.query, nil)
	if err != nil {
		return Result{}
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Errorf("error making request: %s", err.Error())
		return Result{}
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Errorf("error reading body from response: %s", err.Error())
		return Result{}
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Errorf("error unmarshaling body from response: %s", err.Error())
		return Result{}
	}
	return result
}

func (c ErgastF1Client) Seasons(season string) ErgastF1Client {
	c.query = c.query + "/seasons"
	c.query = addResource(c.query, season)
	return c
}

func (c ErgastF1Client) Drivers(driver string) ErgastF1Client {
	c.query = c.query + "/drivers"
	c.query = addResource(c.query, driver)
	return c
}

func addResource(q string, r string) string {
	if r != "" {
		return q + "/" + r
	}
	return q
}

func getFormat(q string, format string) string {
	switch format {
	case "json":
		q = q + ".json"
	case "xml":
		q = q + ".xml"
	default:
		q = q + ".xml"
	}
	return q
}

func (c ErgastF1Client) Constructors(constructor string) ErgastF1Client {
	c.query = c.query + "/constructors"
	c.query = addResource(c.query, constructor)
	return c
}

func (c ErgastF1Client) Grid(position string) ErgastF1Client {
	c.query = c.query + "/grid"
	c.query = addResource(c.query, position)
	return c
}

func (c ErgastF1Client) Results(position string) ErgastF1Client {
	c.query = c.query + "/results"
	c.query = addResource(c.query, position)
	return c
}

func (c ErgastF1Client) Current() ErgastF1Client {
	c.query = c.query + "/current"
	return c
}

func (c ErgastF1Client) DriverStandings(standing string) ErgastF1Client {
	c.query = c.query + "/driverStandings"
	c.query = addResource(c.query, standing)
	return c
}

func (c ErgastF1Client) ConstructorStandings(standing string) ErgastF1Client {
	c.query = c.query + "/constructorStandings"
	c.query = addResource(c.query, standing)
	return c
}
