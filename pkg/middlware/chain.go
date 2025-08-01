package middlware

import "net/http"

type Middlware func(http.Handler) http.Handler

func Chain(middlware ...Middlware) Middlware {
	return func(next http.Handler) http.Handler {
		for i := len(middlware) - 1; i >= 0; i-- {
			next = middlware[i](next)
		}

		return next
	}
}
