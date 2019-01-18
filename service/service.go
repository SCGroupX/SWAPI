package service

import (
    "os"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        Directory:  "templates",
        Extensions: []string{".html"},
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    webRoot := os.Getenv("WEBROOT")
    if len(webRoot) == 0 {
        if root, err := os.Getwd(); err != nil {
            panic("Could not retrive working directory")
        } else {
            webRoot = root
            //fmt.Println(root)
        }
    }

    filmsSubRouter := mx.PathPrefix("/api/films").Subrouter()
    filmsSubRouter.HandleFunc("/{id:[0-9]+}", routerHandler(formatter)).Methods("GET")

    peopleSubRouter := mx.PathPrefix("/api/people").Subrouter()
    peopleSubRouter.HandleFunc("/{id:[0-9]+}", routerHandler(formatter)).Methods("GET")

    planetsSubRouter := mx.PathPrefix("/api/planets").Subrouter()
    planetsSubRouter.HandleFunc("/{id:[0-9]+}", routerHandler(formatter)).Methods("GET")

    speciesSubRouter := mx.PathPrefix("/api/species").Subrouter()
    speciesSubRouter.HandleFunc("/{id:[0-9]+}", routerHandler(formatter)).Methods("GET")

    starshipsSubRouter := mx.PathPrefix("/api/starships").Subrouter()
    starshipsSubRouter.HandleFunc("/{id:[0-9]+}", routerHandler(formatter)).Methods("GET")

    vehiclesSubRouter := mx.PathPrefix("/api/vehicles").Subrouter()
    vehiclesSubRouter.HandleFunc("/{id:[0-9]+}", routerHandler(formatter)).Methods("GET")
    // mx.HandleFunc("/api/people/", peopleHandler(formatter)).Methods("GET")
    // mx.HandleFunc("/api/planets/", planetsHandler(formatter)).Methods("GET")
    // mx.HandleFunc("/api/species/", speciesHandler(formatter)).Methods("GET")
    // mx.HandleFunc("/api/starships/", starshipsHandler(formatter)).Methods("GET")
    // mx.HandleFunc("/api/vehicles/", vehiclesHandler(formatter)).Methods("GET")

    //mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
    mx.HandleFunc("/", homeHandler(formatter))

}