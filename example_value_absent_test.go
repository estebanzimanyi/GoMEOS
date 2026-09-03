package gomeos

// A temporal value holds nothing outside its own span, and MEOS says so by
// answering NULL. These examples pin what the accessors do with that answer:
// they report absence through a second result and leave the error free for a
// MEOS failure, and the nil handle it arrives as is a value the caller can test
// rather than something to dereference.

import (
	"fmt"
	"time"
)

func ExampleTBoolValueAtTimestamp_absent() {
	tb := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02, TRUE@2022-10-03}")
	outside, _ := time.Parse("2006-01-02", "2021-01-01")

	value, ok, err := TBoolValueAtTimestamp(tb, outside)
	fmt.Println(value, ok, err)
	// Output:
	// false false <nil>
}

func ExampleTFloatValueAtTimestamp_absent() {
	tf := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02, 3.4@2022-10-03}", &TFloatSeq{})
	outside, _ := time.Parse("2006-01-02", "2021-01-01")

	value, ok, err := TFloatValueAtTimestamp(tf, outside)
	fmt.Println(value, ok, err)
	// Output:
	// 0 false <nil>
}

func ExampleCreateTemporal_absent() {
	tb := NewTBoolSeq("{FALSE@2022-10-01, TRUE@2022-10-03}")
	outside, _ := time.Parse("2006-01-02", "2021-01-01")

	// Restricting to a moment the value does not cover leaves nothing, which
	// reaches Go as a nil Temporal rather than a handle onto nothing.
	fmt.Println(TemporalAtTimestamptz(tb, outside) == nil)
	// Output:
	// true
}
