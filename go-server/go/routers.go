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
		"/mzalt/swagger-petstore/1.0.0/",
		Index,
	},

	Route{
		"AddPet",
		"POST",
		"/mzalt/swagger-petstore/1.0.0/pet",
		AddPet,
	},

	Route{
		"DeletePet",
		"DELETE",
		"/mzalt/swagger-petstore/1.0.0/pet/{petId}",
		DeletePet,
	},

	Route{
		"FindPetsByStatus",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/pet/findByStatus",
		FindPetsByStatus,
	},

	Route{
		"FindPetsByTags",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/pet/findByTags",
		FindPetsByTags,
	},

	Route{
		"GetPetById",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/pet/{petId}",
		GetPetById,
	},

	Route{
		"UpdatePet",
		"PUT",
		"/mzalt/swagger-petstore/1.0.0/pet",
		UpdatePet,
	},

	Route{
		"UpdatePetWithForm",
		"POST",
		"/mzalt/swagger-petstore/1.0.0/pet/{petId}",
		UpdatePetWithForm,
	},

	Route{
		"UploadFile",
		"POST",
		"/mzalt/swagger-petstore/1.0.0/pet/{petId}/uploadImage",
		UploadFile,
	},

	Route{
		"DeleteOrder",
		"DELETE",
		"/mzalt/swagger-petstore/1.0.0/store/order/{orderId}",
		DeleteOrder,
	},

	Route{
		"GetInventory",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/store/inventory",
		GetInventory,
	},

	Route{
		"GetOrderById",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/store/order/{orderId}",
		GetOrderById,
	},

	Route{
		"PlaceOrder",
		"POST",
		"/mzalt/swagger-petstore/1.0.0/store/order",
		PlaceOrder,
	},

	Route{
		"CreateUser",
		"POST",
		"/mzalt/swagger-petstore/1.0.0/user",
		CreateUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		"POST",
		"/mzalt/swagger-petstore/1.0.0/user/createWithArray",
		CreateUsersWithArrayInput,
	},

	Route{
		"CreateUsersWithListInput",
		"POST",
		"/mzalt/swagger-petstore/1.0.0/user/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/mzalt/swagger-petstore/1.0.0/user/{username}",
		DeleteUser,
	},

	Route{
		"GetUserByName",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/user/{username}",
		GetUserByName,
	},

	Route{
		"LoginUser",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		"GET",
		"/mzalt/swagger-petstore/1.0.0/user/logout",
		LogoutUser,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/mzalt/swagger-petstore/1.0.0/user/{username}",
		UpdateUser,
	},

}