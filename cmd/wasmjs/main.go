package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"syscall/js"

	"github.com/aaronland/go-tags"
)

func DeriveHashTagsFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		body := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			tags_derived, err := tags.DeriveHashTagsFromString(body)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to parse body, %v\n", err))
				return nil
			}

			tags_str := make([]string, len(tags_derived))

			for idx, t := range tags_derived {
				tags_str[idx] = t.Raw()
			}

			tags_enc, err := json.Marshal(tags_str)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to marshal results, %v\n", err))
				return nil
			}

			resolve.Invoke(string(tags_enc))
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {

	derive_func := DeriveHashTagsFunc()
	defer derive_func.Release()

	js.Global().Set("derive_hashtags_from_string", derive_func)

	c := make(chan struct{}, 0)

	slog.Info("Tag functions registered")
	<-c
}
