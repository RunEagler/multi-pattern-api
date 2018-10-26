package controller

import (
	"testing"
	"GolandProject/multi-pattern-api/common"
	"net/http"
	"GolandProject/multi-pattern-api/app"
	"GolandProject/multi-pattern-api/app/test"
	"github.com/magiconair/properties/assert"
)

func TestSamples(t *testing.T) {

	type testCase struct {
		testPattern common.TestPattern
		response    app.User
	}

	testCases := []testCase{
		{
			testPattern: common.TestPattern{
				Title:      "正常系",
				StatusCode: http.StatusOK,
			},
			response: app.User{
				Name: "test_user",
				Age:  34,
			},
		},
	}
	ctrl := NewSamplesController(common.Service)

	for _, testCase := range testCases {

		switch testCase.testPattern.StatusCode {
		case http.StatusOK:
			_, actualResponse := test.PingSamplesOK(t, nil, common.Service, ctrl)
			assert.Equal(t, testCase.response.Name, actualResponse.Name)
			assert.Equal(t, testCase.response.Age, actualResponse.Age)
		case http.StatusInternalServerError:
			test.PingSamplesInternalServerError(t, nil, common.Service, ctrl)
		default:
			t.Error("not seting status code")
		}
	}

}
