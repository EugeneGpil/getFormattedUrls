package app

import "testing"

func Test_should_return_two_slashes_if_empty_string_provided_as_a_parameter(t *testing.T) {
	url1, url2 := Run("")

	if url1 != "/" || url2 != "/" {
		t.Fatalf(`url1, url2 = %q, %q, want match for "/", "/"`, url1, url2)
	}
}
