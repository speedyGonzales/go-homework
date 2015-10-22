package main

import (
	"testing"
)

func TestExtractingIPs(t *testing.T) {
	logContents := `2015-08-23 12:37:03 8.8.8.8 As far as we can tell this is a DNS
2015-08-23 12:37:04 8.8.4.4 Yet another DNS, how quaint!
2015-08-23 12:37:05 208.122.23.23 There is definitely some trend here
`

	expected := `8.8.8.8
8.8.4.4
208.122.23.23
`

	found := ExtractColumn(logContents, 1)

	if found != expected {
		t.Errorf("Expected\n---\n%s\n---\nbut found\n---\n%s\n---\n", expected, found)
	}
}

func TestExtractingDates(t *testing.T) {
	logContents := `2015-08-23 12:37:03 8.8.8.8 As far as we can tell this is a DNS
2015-08-23 12:37:04 8.8.4.4 Yet another DNS, how quaint!
2015-08-23 12:37:05 208.122.23.23 There is definitely some trend here
`

	expected := `2015-08-23 12:37:03
2015-08-23 12:37:04
2015-08-23 12:37:05
`

	found := ExtractColumn(logContents, 0)

	if found != expected {
		t.Errorf("Expected\n---\n%s\n---\nbut found\n---\n%s\n---\n", expected, found)
	}
}

func TestExtractingTexts(t *testing.T) {
	logContents := `2015-08-23 12:37:03 8.8.8.8 As far as we can tell this is a DNS
2015-08-23 12:37:04 8.8.4.4 Yet another DNS, how quaint!
2015-08-23 12:37:05 208.122.23.23 There is definitely some trend here
`

	expected := `As far as we can tell this is a DNS
Yet another DNS, how quaint!
There is definitely some trend here
`

	found := ExtractColumn(logContents, 2)

	if found != expected {
		t.Errorf("Expected\n---\n%s\n---\nbut found\n---\n%s\n---\n", expected, found)
	}
}
