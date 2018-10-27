package controller

import (
	"testing"
	"GolandProject/multi-pattern-api/test"
	"net/http"
	"GolandProject/multi-pattern-api/app"
	goaTest "GolandProject/multi-pattern-api/app/test"
	"github.com/magiconair/properties/assert"
)

func TestSamples(t *testing.T) {

	type testCase struct {
		testPattern test.TestPattern
		response    app.User
	}

	testCases := []testCase{
		{
			testPattern: test.TestPattern{
				Title:      "正常系",
				StatusCode: http.StatusOK,
			},
			response: app.User{
				Name: "test_user",
				Age:  34,
			},
		},
	}
	ctrl := NewSamplesController(test.Service)

	for _, testCase := range testCases {

		switch testCase.testPattern.StatusCode {
		case http.StatusOK:
			_, actualResponse := goaTest.PingSamplesOK(t, nil, test.Service, ctrl)
			assert.Equal(t, testCase.response.Name, actualResponse.Name)
			assert.Equal(t, testCase.response.Age, actualResponse.Age)
		case http.StatusInternalServerError:
			goaTest.PingSamplesInternalServerError(t, nil, test.Service, ctrl)
		default:
			t.Error("not seting status code")
		}
	}

}
