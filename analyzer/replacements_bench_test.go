package analyzer_test

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkStringFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint("hello") //nolint:gosimple //https://staticcheck.io/docs/checks#S1039
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%s", "hello") //nolint:gosimple //https://staticcheck.io/docs/checks#S1025
		}
	})

	b.Run("just string", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = "hello"
		}
	})
}

func BenchmarkErrorFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(context.DeadlineExceeded)
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%s", context.DeadlineExceeded)
		}
	})

	b.Run("Error()", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = context.DeadlineExceeded.Error()
		}
	})
}

func BenchmarkFormattingError(b *testing.B) {
	b.Run("fmt.Errorf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Errorf("onlystring")
		}
	})

	b.Run("errors.New", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = errors.New("onlystring")
		}
	})
}

func BenchmarkBoolFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(true)
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%t", true)
		}
	})

	b.Run("strconv.FormatBool", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strconv.FormatBool(true)
		}
	})
}

func BenchmarkHexEncoding(b *testing.B) {
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%x", []byte{'a', 'b', 'c'})
		}
	})

	b.Run("hex.EncodeToString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = hex.EncodeToString([]byte{'a', 'b', 'c'})
		}
	})
}

func BenchmarkHexArrayEncoding(b *testing.B) {
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			val := [3]byte{'a', 'b', 'c'}
			_ = fmt.Sprintf("%x", val)
		}
	})

	b.Run("hex.EncodeToString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			val := [3]byte{'a', 'b', 'c'}
			_ = hex.EncodeToString(val[:])
		}
	})
}

func BenchmarkIntFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(math.MaxInt)
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%d", math.MaxInt)
		}
	})

	b.Run("strconv.Itoa", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strconv.Itoa(math.MaxInt)
		}
	})
}

func BenchmarkIntConversionFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		u := int32(0x12345678)
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(u)
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		u := int32(0x12345678)
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%d", u)
		}
	})

	b.Run("strconv.FormatInt", func(b *testing.B) {
		u := int32(0x12345678)
		for n := 0; n < b.N; n++ {
			_ = strconv.FormatInt(int64(u), 10)
		}
	})
}

func BenchmarkUintFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(uint64(math.MaxUint))
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%d", uint64(math.MaxUint))
		}
	})

	b.Run("strconv.FormatUint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strconv.FormatUint(math.MaxUint, 10)
		}
	})
}

func BenchmarkUintHexFormatting(b *testing.B) {
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%x", uint64(math.MaxUint))
		}
	})

	b.Run("strconv.FormatUint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strconv.FormatUint(math.MaxUint, 16)
		}
	})
}

func BenchmarkStringAdditionFormatting(b *testing.B) {
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("Hello %s", "world")
		}
	})

	b.Run("string concatenation", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = "Hello " + "world"
		}
	})
}

func BenchmarkStringConcatLoop(b *testing.B) {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			s := ""
			for w := range words {
				s += words[w]
			}
			_ = s
		}
		b.ReportAllocs()
	})

	b.Run("strings Builder", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var s string
			var sb strings.Builder
			for w := range words {
				sb.WriteString(words[w])
			}
			s = sb.String()
			_ = s
		}
		b.ReportAllocs()
	})
}

func BenchmarkStringConcatLoopBig(b *testing.B) {
	words := make([]string, 0x10000)
	for i := 0; i < 0x10000; i++ {
		words[i] = strconv.Itoa(i)
	}
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			s := ""
			for w := range words {
				s += words[w]
			}
			_ = s
		}
		b.ReportAllocs()
	})

	b.Run("strings Builder", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var s string
			var sb strings.Builder
			for w := range words {
				sb.WriteString(words[w])
			}
			s = sb.String()
			_ = s
		}
		b.ReportAllocs()
	})
}
