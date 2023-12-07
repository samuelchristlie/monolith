package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"

	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/pbkdf2"

	"crypto/sha256"

	"encoding/hex"
	"math/big"

	"strconv"
	// "fmt"

)

func generate_salt(site string, user string, counter int) string{
	return site + user + strconv.FormatInt(int64(counter), 16)
}

func generate_hash_scrypt(master string, salt string) string{
	hash, _ := scrypt.Key([]byte(master), []byte(salt), 16384, 8, 5, 32)

	return hex.EncodeToString(hash)
}

func generate_hash_pbkdf2(master string, salt string) string{
	hash := pbkdf2.Key([]byte(master), []byte(salt), 600000, 32, sha256.New)

	return hex.EncodeToString(hash)
}

func get_characters(digits bool, symbols bool, uppercase bool, lowercase bool) string{
	characters := ""

	if digits {
		characters += "0123456789"
	}

	if symbols {
		characters += "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	}

	if uppercase {
		characters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if lowercase {
		characters += "abcdefghijklmnopqrstuvwxyz"
	}

	return characters
}

func get_max_length(length int, digits bool, symbols bool, uppercase bool, lowercase bool) int {
	if digits {
		length--
	}

	if symbols {
		length--
	}

	if uppercase{
		length--
	}

	if lowercase{
		length--
	}

	return length
}


func consume(entropy big.Int, characters string, max_len int) (string, big.Int){
	generated := ""
	
	quotient := entropy
	remainder := new(big.Int)


	for i := 0; i < max_len; i++{
		dividend := quotient
		divisor := big.NewInt(int64(len(characters)))

		quotient.DivMod(&dividend, divisor, remainder)

		generated += string(characters[int(remainder.Int64())])
	}

	return generated, quotient

}

func insert(generated string, entropy big.Int, characters string) string{
	quotient := entropy
	remainder := new(big.Int)

	for _, x := range characters {
		dividend := quotient
		divisor := big.NewInt(int64(len(generated)))

		quotient.DivMod(&dividend, divisor, remainder)

		generated = generated[:int(remainder.Int64())] + string(x) + generated[int(remainder.Int64()):]

	}

	return generated
}

func generate_password(site string, user string, master string, counter int, length int, digits bool, symbols bool, uppercase bool, lowercase bool) string{
	salt := generate_salt(site, user, counter)
	
	entropy := new(big.Int)
	entropy.SetString(generate_hash_scrypt(master, salt), 16)


	

	characters := get_characters(digits, symbols, uppercase, lowercase)

	max_len := get_max_length(length, digits, symbols, uppercase, lowercase)

	password, password_entropy := consume(*entropy, characters, max_len)

	


	chars_to_add := ""
	var value string

	if digits {
		value, password_entropy = consume(password_entropy, "0123456789", 1)
		chars_to_add += value
	}
	if symbols {
		value, password_entropy = consume(password_entropy, "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", 1)
		chars_to_add += value
	}
	if uppercase {
		value, password_entropy = consume(password_entropy, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1)
		chars_to_add += value
	}
	if lowercase {
		value, password_entropy = consume(password_entropy, "abcdefghijklmnopqrstuvwxyz", 1)
		chars_to_add += value
	}
	
	return insert(password, password_entropy, chars_to_add)
	

}




func main(){
	// fmt.Println(generate_password("facebook.com", "user", "pass", 8, 16, true, true, true, true))

	a := app.New()
	w := a.NewWindow("Monolith")

	

	siteInput := widget.NewEntry()
	siteInput.SetPlaceHolder("www.website.com")

	userInput := widget.NewEntry()
	userInput.SetPlaceHolder("Username")

	masterInput := widget.NewPasswordEntry()
	masterInput.SetPlaceHolder("Password")


	counterInput := widget.NewLabel("1")
	counterSubtractButton := widget.NewButton("-", func(){
		counter, _ := strconv.Atoi(counterInput.Text)

		if counter > 1{
			counter--
		}
		
		counterInput.SetText(strconv.Itoa(counter))
	})
	counterAddButton := widget.NewButton("+", func(){
		counter, _ := strconv.Atoi(counterInput.Text)
		counter++

		counterInput.SetText(strconv.Itoa(counter))
	})


	lengthInput := widget.NewLabel("16")
	lengthSubtractButton := widget.NewButton("-", func(){
		length, _ := strconv.Atoi(lengthInput.Text)

		if length > 6{
			length--
		}
		
		lengthInput.SetText(strconv.Itoa(length))
	})
	lengthAddButton := widget.NewButton("+", func(){
		length, _ := strconv.Atoi(lengthInput.Text)
		length++

		lengthInput.SetText(strconv.Itoa(length))
	})

	digitsInput := widget.NewCheck("Digits", func(bool){})
	digitsInput.SetChecked(true)

	symbolsInput := widget.NewCheck("Symbols", func(bool){})
	symbolsInput.SetChecked(true)

	uppercaseInput := widget.NewCheck("Uppercase", func(bool){})
	uppercaseInput.SetChecked(true)

	lowercaseInput := widget.NewCheck("Lowercase", func(bool){})
	lowercaseInput.SetChecked(true)

	generatedLabel := widget.NewLabel("Generated Password")



	// label := widget.NewLabel("Hello World")
	// container.New(layout.NewCenterLayout(), label),
			
	// widget.NewButton("Hi", func(){
	// 	label.SetText("Welcome :)")
	// }),

	w.SetContent(
		container.NewVBox(
			container.New(layout.NewCenterLayout(), widget.NewLabel("Website")),
			siteInput,

			container.New(layout.NewCenterLayout(), widget.NewLabel("Username")),
			userInput,

			container.New(layout.NewCenterLayout(), widget.NewLabel("Master Password")),
			masterInput,

			container.NewHBox(
				layout.NewSpacer(),
				container.NewVBox(
					container.New(layout.NewCenterLayout(), widget.NewLabel("Counter")),
					container.NewHBox(counterSubtractButton, counterInput, counterAddButton,),
				),
				layout.NewSpacer(),
				container.NewVBox(
					container.New(layout.NewCenterLayout(), widget.NewLabel("Length")),
					container.NewHBox(lengthSubtractButton, lengthInput, lengthAddButton),
				),
				layout.NewSpacer(),
			),
			container.New(layout.NewGridLayout(2), digitsInput, symbolsInput, uppercaseInput, lowercaseInput),

			container.New(layout.NewCenterLayout(), generatedLabel),
			widget.NewButton("Generate and Copy", func(){
				counter, _ := strconv.Atoi(counterInput.Text)
				length, _ := strconv.Atoi(lengthInput.Text)


				generated_password := generate_password(siteInput.Text, userInput.Text, masterInput.Text, counter, length, digitsInput.Checked, symbolsInput.Checked, uppercaseInput.Checked, lowercaseInput.Checked)
				generatedLabel.SetText(generated_password)
				w.Clipboard().SetContent(generated_password)
			}),
		))

	w.Resize(fyne.NewSize(float32(300), float32(w.Canvas().Size().Height)))
	w.SetFixedSize(true)

	w.ShowAndRun()
}