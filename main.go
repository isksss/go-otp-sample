package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/pquerna/otp/totp"
)

func main() {
	var s string
	fileName := "otp.key"
	if ExistKey(fileName) {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "Example.com",
			AccountName: "alice@example.com",
		})

		if err != nil {
			log.Fatal(err)
		}

		img, err := key.Image(200, 200)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create("QR.png")

		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := png.Encode(f, img); err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(fileName, []byte(key.Secret()), 0600)
		s = key.Secret()
		fmt.Println("save new otp")
	} else {
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		s = string(b)

	}

	var passcode string
	fmt.Print("input otp:")
	fmt.Scanf("%s", &passcode)
	valid := totp.Validate(passcode, s)

	if valid {
		fmt.Println("succeed")
	} else {
		fmt.Println("denied")
	}

}

func ExistKey(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return true
	} else {
		return false
	}
}
