package main

import(
   "fmt"
   "os"
   "io/ioutil"
   "net/http"
) 

func getPage(url string) (int, error){ 
   resp, err := http.Get(url)
   if err != nil {
   return 0, err
   }
  
   defer resp.Body.Close()
   
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
	return 0, err
   }

   return len(body), nil
}

func main() {

  	urls := []string{"https://staff.hennepintech.edu", 
"http://admin.hennepintech.edu",
"http://ardeshir.org",
"http://google.com", "http://www.yahoo.com","http://www.bing.com","http://bbc.co.uk"}

       for _, url := range urls {

          pageLength, err := getPage(url)
          if err != nil {
	   os.Exit(1)
          }

          fmt.Printf("%s is length %d\n", url, pageLength)

     }
}
