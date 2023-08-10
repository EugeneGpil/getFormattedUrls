package app

import (
	"testing"

	"github.com/EugeneGpil/getFormattedUrls/app"
)

func Test_should_return_two_slashes_if_empty_string_provided_as_a_parameter(t *testing.T) {
	url1, url2 := app.Run("")

	if url1 != "/" || url2 != "/" {
		t.Fatalf(`url1, url2 = %q, %q, want match for "/", "/"`, url1, url2)
	}
}

func Test_should_return_correct_urls(t *testing.T) {
	minimalUrl := "correct/url"
	correctUrlShort := "/" + minimalUrl
	correctUrlLong := correctUrlShort + "/"

	res1Short, res1Long := app.Run(minimalUrl)
	res2Short, res2Long := app.Run(correctUrlShort)
	res3Short, res3Long := app.Run(minimalUrl + "/")
	res4Short, res4Long := app.Run(correctUrlLong)

	shortResults := map[string]string{
		"res1Short": res1Short,
		"res2Short": res2Short,
		"res3Short": res3Short,
		"res4Short": res4Short,
	}

	longResults := map[string]string{
		"res1Long": res1Long,
		"res2Long": res2Long,
		"res3Long": res3Long,
		"res4Long": res4Long,
	}

	assertUrls(t, shortResults, correctUrlShort)
	assertUrls(t, longResults, correctUrlLong)
}

func assertUrls(t *testing.T, urls map[string]string, correctUrl string) {
	for key, value := range urls {
		if value != correctUrl {
			t.Fatalf(`%q = %q, want match for %q`, key, value, correctUrl)
		}
	}
}
