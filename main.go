package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		println("please supply a JWT token")
		os.Exit(1)
	}
	rawToken := os.Args[1]

	parts := strings.Split(rawToken, ".")
	if len(parts) != 3 {
		println("unexpected token segment count")
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	partNames := []string{"header", "body", "signature"}
	for i, name := range partNames {
		contents, err := base64.RawURLEncoding.DecodeString(
			parts[i],
		)
		if err != nil {
			println("failed to decode", name+":", err.Error())
		}

		println(name + ":")
		if name == "signature" {
			println("  base64:", parts[i])
			println("  hex:", hex.EncodeToString(contents))
			println()
			continue
		}

		data := make(map[string]interface{})
		if err := json.Unmarshal(contents, &data); err != nil {
			println("failed to unmarshal", name)
			os.Exit(1)
		}

		if err := enc.Encode(data); err != nil {
			println("failed to marshal", name)
			os.Exit(1)
		}
		println()

		if name == "body" {
			expTs, ok := data["exp"].(float64)
			if ok {
				expTime := time.Unix(int64(expTs), 0)
				println("expires: ", expTime.String())
				println()
			}
		}

	}
}
