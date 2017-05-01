package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"//",
		Index,
	},

	Route{
		"GetAppartment",
		"GET",
		"//appartment",
		GetAppartment,
	},

	Route{
		"GetAppartmentsForBuilding",
		"GET",
		"//buildings/{id}/appartments",
		GetAppartmentsForBuilding,
	},

	Route{
		"GetBuilding",
		"GET",
		"//buildings",
		GetBuilding,
	},

	Route{
		"GetCompanies",
		"GET",
		"//companies",
		GetCompanies,
	},

	Route{
		"GetCompanyLocation",
		"GET",
		"//companies/{id_company}/location",
		GetCompanyLocation,
	},

	Route{
		"GetProfiles",
		"GET",
		"//profiles",
		GetProfiles,
	},

	Route{
		"GetTypes",
		"GET",
		"//User_types",
		GetTypes,
	},

	Route{
		"CreateUser",
		"POST",
		"//users",
		CreateUser,
	},

	Route{
		"DeleteUserByid",
		"DELETE",
		"//users/{id}",
		DeleteUserByid,
	},

	Route{
		"DeleteUserVehiculesByid",
		"DELETE",
		"//users/{id}/vehicules/{idVehicule}",
		DeleteUserVehiculesByid,
	},

	Route{
		"GetUserByid",
		"GET",
		"//users/{id}",
		GetUserByid,
	},

	Route{
		"GetUserLives",
		"GET",
		"//users/{id}/lives",
		GetUserLives,
	},

	Route{
		"GetUserVehiculesByid",
		"GET",
		"//users/{id}/vehicules",
		GetUserVehiculesByid,
	},

	Route{
		"GetUsers",
		"GET",
		"//users",
		GetUsers,
	},

	Route{
		"PostUserVehiculesByid",
		"POST",
		"//users/{id}/vehicules",
		PostUserVehiculesByid,
	},

	Route{
		"UpdateUser",
		"PUT",
		"//users/{id}",
		UpdateUser,
	},

	Route{
		"UpdateUserVehiculesByid",
		"PUT",
		"//users/{id}/vehicules/{idVehicule}",
		UpdateUserVehiculesByid,
	},

}