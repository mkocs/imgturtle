/*
 * IMGCAT
 * GO PROTOTYPE
 * 2015
 */

package http

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/fatih/color"
  "github.com/codegangsta/negroni"
)

/*
 * initialise http server
 */
func Init(){
  r := mux.NewRouter().StrictSlash(false)

  /*
   * main page
   */
  r.HandleFunc("/", mainPageHandler)

  /*
   * custom profile url
   */
  r.HandleFunc("/people/{name}", peoplePageHandler)

  /*
   * personal profile url
   */
   r.HandleFunc("/me", mePageHandler)

   /*
    *
    */
    r.Handle("/img/", http.FileServer(http.Dir("/public/img/")))

  /*
   * init negroni middleware
   */
  n := negroni.New(
    negroni.NewRecovery(),
    negroni.HandlerFunc(MiddleWare),
    //negroni.NewLogger(),
    negroni.NewStatic(http.Dir("public")),
  )
  n.UseHandler(r)
  color.Green("IMGCAT Server running on port 8000")

  http.ListenAndServe(":8000", n)
  //http.ListenAndServeTLS(port, certificate.pem, key.pem, nil) for https
}
