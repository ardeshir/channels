package simplecms

import (
   "net/http"
   "time"
   "strings"
)

func HandleNew(w http.ResponseWriter, r *http.Request) {
    
    switch r.Method {
   
    case "GET":
        Tmpl.ExecuteTemplate(w, "new", nil)

    case "POST":
       title       := r.FormValue("title")
       content     := r.FormValue("content")
       contentType := r.FormValue("content-type")
       r.ParseForm()

       if contentType == "page" {
          Tmpl.ExecuteTemplate(w, "page", &Page{
               Title: title, 
               Content: content,
          })
          return
       }
       
       if contentType == "post" {
           Tmpl.ExecuteTemplate(w, "post", &Post{
              Title: title,
              Content: content,
            })
           return
       }

     default: 
        http.Error(w, "Method not supported: "+r.Method, http.StatusMethodNotAllowed)
    } // end of switch

} // end of new


//  Handler to deal with index
func ServeIndex(w http.ResponseWriter, r *http.Request) {
   
   p := &Page{
       Title: "Go SimpleCMS",
       Content: "Welcome to SimpleCMS Dashboard!",
       Posts: []*Post{
            &Post{
               Title:  "Hi, from Post 1",
               Content: "This is the content of Post 1",
               DatePublished: time.Now(),
            },
            &Post{
               Title:  "Hi, from Post 2",
               Content:  "This is the content of Post 2",
               DatePublished: time.Now().Add(-time.Hour),
               Comments: []*Comment{
                    &Comment{
                       Author:  "Ardeshir Sepahsalar",
                       Comment: "This is the first comment...",
                       DatePublished: time.Now().Add(-time.Hour / 2 ),
                   },
                },
             },
          },
      } 

     Tmpl.ExecuteTemplate(w, "page", p)
}

func ServePage(w http.ResponseWriter, r *http.Request) {
   path := strings.TrimLeft(r.URL.Path, "/page/")
   if path == "" {
    http.NotFound(w,r)
    return
  }

  p := &Page{
      Title: strings.ToTitle(path),
      Content: "Here is my page",
  }
  Tmpl.ExecuteTemplate(w, "page", p)


} // end of serve page


// ServePost servers a post
func ServePost(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimLeft(r.URL.Path, "/post/")

    if path == "" {
      http.NotFound(w,r)
      return
    }

   p := &Post{
       Title: strings.ToTitle(path),
       Content: "Here is my page",
       Comments: []*Comment{
            &Comment{
                Author: "Ardeshir Sepahsalar",
                Comment: "Looks super!",
                DatePublished: time.Now(),
           },
       },
    }
    Tmpl.ExecuteTemplate(w, "post", p)
} // end serve post
