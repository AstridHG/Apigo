package main

import (

 "net/http"

	"fmt"

	facebook "facebook-go-sdk.wiki"
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
  fb.SetAccessToken(`EAAJl27dqpTwBAPDM6nIbC10fRaZAQi68EDhidyuiXoNdvtaCVGya0e5tKbDVpULz4SdNmKroYXA3m3gDad6zBBd43Rhuc4dsMtnfYwimos831ZBAdnckWifYDLqPt9M0XSJEjDqQX9kxtTZAlXAr6Oofc4m2bos6SArYY6qgfGqvp4tQhu2RY56ZCctBr6SI8fasGA8ZAMwZDZD`)

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
