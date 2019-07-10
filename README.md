# Go package for accessing https://holidayapi.com/

``` go
client := holidayapi.New("my-key")
holidays, err := client.Holidays(holidayapi.Params{
	Country: "US",
	Year:    2018,
	Public:  true,
})
if err != nil {
	log.Fatal(err)
}
for _, h := range holidays {
	fmt.Println(h)
}
```

