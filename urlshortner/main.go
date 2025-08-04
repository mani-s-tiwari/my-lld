package main

import (
	"fmt"
	service "urlshort/service"
)

func main() {
	url := service.NewUrlService()

	newurl := url.MakeUrl("www.thefacebook.com")
	fmt.Printf("#%s is mapped to #%s \n", newurl.Original, newurl.ShortUrl)

	redirected := url.RedirectUrl(newurl.ShortUrl)
	fmt.Printf("Redirected url #%s \n", redirected)

	editedurl := url.EditUrl(newurl.ShortUrl, "www.facebook.com")
	fmt.Printf("#%s is mapped to #%s \n", editedurl.Original, editedurl.ShortUrl)

	if exist := url.DeleteUrl("www.facebook.com"); exist {
		fmt.Println("deleted successfully")
	} else {
		fmt.Println("does not exist")
	}
}
