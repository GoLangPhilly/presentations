type Person struct {
	Name string
	Age int
}

type APIResponse struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func before() {
	return Person {
		Name: resp.Name,
		Age: resp.Age,
	}

func now() {
	return Person(resp)
}
