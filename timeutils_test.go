package timeutils

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	var (
		date     = time.Date(2022, time.June, 2, 0, 0, 0, 0, time.UTC)
		expected = "2022-06-02"
		got      = date.Format(DateFormat)
	)

	if expected != got {
		t.Errorf("should returns only date: expected=%v, got=%v\n", expected, got)
	}
}

func TestTimeFormat(t *testing.T) {
	var (
		date     = time.Date(2022, time.June, 2, 10, 30, 55, 0, time.UTC)
		expected = "10:30:55"
		got      = date.Format(TimeFormat)
	)

	if expected != got {
		t.Errorf("should returns only date: expected=%v, got=%v\n", expected, got)
	}
}

func TestDateTimeFormat(t *testing.T) {
	var (
		date     = time.Date(2022, time.June, 2, 10, 30, 55, 0, time.UTC)
		expected = "2022-06-02 10:30:55"
		got      = date.Format(DateTimeFormat)
	)

	if expected != got {
		t.Errorf("should returns only date: expected=%v, got=%v\n", expected, got)
	}
}

func TestFullDateTimeFormat(t *testing.T) {
	var (
		date     = time.Date(2022, time.June, 2, 10, 30, 55, 0, time.UTC)
		expected = "2022-06-02 10:30:55+300"
		got      = date.Format(FullDateTimeFormat)
	)

	if expected != got {
		t.Errorf("should returns only date: expected=%v, got=%v\n", expected, got)
	}
}
