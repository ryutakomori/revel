package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestHealthIndex() {
	// request
	t.Get("/health")

	// checking data
	//checker := new(model.HealthResponse)
	checkerStr := `{"code": 1}`
	//json.Unmarshal(([]byte)(checkerStr), checker)

	// response data
	// response := new(model.HealthResponse)
	// json.Unmarshal(t.TestSuite.ResponseBody, response)

	// assert
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	t.AssertEqual(t.TestSuite.ResponseBody, checkerStr)
	//t.AssertEqual(checker, response)
}

func (t *AppTest) TestAuthRegister() {

}

func (t *AppTest) After() {
	println("Tear down")
}
