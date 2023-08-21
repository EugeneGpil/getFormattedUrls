package app

import (
	"testing"

	"github.com/EugeneGpil/getFormattedUrls/app"
	"github.com/EugeneGpil/tester"
)

func Test_should_return_one_url_with_slash_if_empty_string_provided_as_a_parameter(t *testing.T) {
	tester.SetTester(t)
	urls := app.Run("")

	tester.AssertLen(urls, 1)
	tester.AssertSame(urls[0], "/")
}

func Test_should_return_correct_urls(t *testing.T) {
	tester.SetTester(t)

	minimalUrl := "correct/url"
	correctUrlShort := "/" + minimalUrl
	correctUrlLong := correctUrlShort + "/"

	inputUrls := []string{
		minimalUrl,
		correctUrlShort,
		minimalUrl + "/",
		correctUrlLong,
	}

	for _, inputUrl := range inputUrls {
		result := app.Run(inputUrl)

		tester.AssertLen(result, 2)
		tester.AssertSame(result[0], correctUrlShort)
		tester.AssertSame(result[1], correctUrlLong)
	}
}
