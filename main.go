package main

import (

 "net/http"

	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)


// BasicFeedPost will show you about
// how to post a feed to your timeline
func main() {


  http.HandleFunc("/", homeHandler)

  http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  // initalize facebook-go-sdk
  fb := facebook.New()

  // set your access token
  // NOTES: Please exchange with your access token
  fb.SetAccessToken(`EAAJl27dqpTwBADgnF5IjmZCHkZC4tDXYV5VEZARS7cT8lnpGQfmaPAnsVa3PiyIOBDIXomuElO9h3TNv2Bc2grOuJnZB11oEZBBQWadRI8EejvHGMhT12gkQfBH5swi3QZCjJZBF9uTRIBoh5WR26YhQznO7EQIPr8ZD`)

  // submit your feed
  data, err := fb.API(`/me/feed`).Messages(`Hola mundo desde Go Astrid 1202`).Post()
  if err != nil {
    panic(err)
  }

  fmt.Println(`
    ## SAMPLE - POST A FEED
  `)

  fmt.Println(data)

  // if you want to share a link you can use the `Link` method
  // e.g. fb.API(`/me/feed`).Messages(`coba yo`).Link(`http://madebyais.com`).Post()
}
