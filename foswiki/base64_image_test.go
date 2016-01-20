package foswiki

import "testing"

var TestGifBase64 = []byte("R0lGODlhLAAPAOZbAMsAMuaAmMwCNMwAMf///dk/ZdhAZv/+//ne5f7+/tYwWM0BMssANeV+mfLAzMwCMs4AM9AQQdxQc/7+/O+vv+yfs9MfTNhAZ/XP2s0BNvG/yt9fgOiPpdEgTNAhTPvv8eJvjO+xwOFui84AMeaAlvK+yuZ+meuesPXQ2NcwW+ugtOuhss8RQeN/l+mQpvjf5dEgTvjd5Ng+ZNAQQ/LAy/TQ2tYxWeR+luRvitUwWtc/Zfrf5u6wvddBZtEiTdhCZ91gfssBMcsAMMoAMOR9mPTO2dtPcvnf6ORvjNlBZ+Bhf//9/tk/Y8oAMvvd5e2eseSAmtASQtUuWd1RdNIhTeSAmMwANc0BNMsBM////8wAMwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACH5BAEAAFsALAAAAAAsAA8AAAf/gA9XVwtaWliDhlhbW4aGVwKOWI6UipOOV44LC5eUVo6MjlYDWgCUnZKVWiNYppqKqoysqqiVk01XrloMqo8uCpSMur3ElJGex5MeSwVDh1rClwJHBqWHn88DtSgGApdYx4YcWQVaV4iMlwohWSAGDBE4AUpRiDYkAQUDKTwHSNWVFmTYkIXcMwahDDlAkKUIjRlOEDR4UUOLBAIUKBxo4GBHFgwarAhxxGJKiYIJngQI0KJKOlcBCCTBAiILsItAoGQhoqXBBiwxDWAzx6ABgQkJChIoeIBAqEwAAmTpccVBFgIEEhy4YYHhhxDApF7AQgqWORlZJggYIOBcurIBkQ5Uq0AggpYIPzooMEDFRBYEQLNcuGZNCykdBVElPCQ1gIICBCp0qEDuRAIJHj4g0CLigIoclMpqKVCw0mItPjAk4MCZYQwjACxQSKtBihYYDg6sqARBixXSWUw38mQI25Xe5nz7zmTlSgZKV0QzSSqcEoRMqlwNnQQACwTRWtqCA37sygNGDxSBL0YsiKEHgQAAOw==")
var TestGifDimensions = []int{820, 44, 15}

func TestNewBase64Image(t *testing.T) {
	filename := "test.gif"
	image, _ := NewBase64Image(filename, TestGifBase64)

	switch {
	case image.FileName() != filename:
		t.Errorf("image.FileName() = %v, expected: %v", image.FileName(), filename)
	case image.Size() != TestGifDimensions[0]:
		t.Errorf("image.Size() = %v, expected: %v", image.Size(), TestGifDimensions[0])
	case image.Width() != TestGifDimensions[1]:
		t.Errorf("image.Width() = %v, expected: %v", image.Width(), TestGifDimensions[1])
	case image.Height() != TestGifDimensions[2]:
		t.Errorf("image.Width() = %v, expected: %v", image.Height(), TestGifDimensions[2])
	}
}
