package productRouters

import (
	"ShopRestAPI/internal/models/products"
	"ShopRestAPI/internal/routers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (r *ProductRoutes) ConfigurePropertiesRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/props/color", r.CreateColorRouter("color"))
	mux.HandleFunc("/props/country", r.CreateColorRouter("country"))
	mux.HandleFunc("/props/sizes", r.CreateColorRouter("sizes"))
	mux.HandleFunc("/props/styles", r.CreateColorRouter("styles"))
	mux.HandleFunc("/props/types", r.CreateColorRouter("types"))
	mux.HandleFunc("/props/season", r.CreateColorRouter("season"))
	mux.HandleFunc("/property", r.CreatePropertyRouter())
}

func (r *ProductRoutes) CreateColorRouter(table string) http.HandlerFunc {
	type response struct {
		Data []string `json:"data"`
	}
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := r.store.Product().GetPropertyList(table)
		if err != nil {
			routers.ErrorHelper(w, req, http.StatusBadRequest, err)
		}
		res := &response{
			Data: data,
		}
		info, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s\n", info)
	}
}

func (r *ProductRoutes) CreatePropertyRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "DELETE" {
			param := req.URL.Query().Get("id")
			id, err := strconv.ParseInt(param, 10, 32)
			if err != nil {
				routers.ErrorHelper(w, req, http.StatusBadRequest, err)
			}
			r.store.Product().DeleteInstance(int(id))
		} else if req.Method == "PUT" {
			var property = &products.Instance{}
			if err := json.NewDecoder(req.Body).Decode(property); err != nil {
				log.Fatal(err)
			}
			if property.Id == 0 {
				if err := r.store.Product().CreateInstance(property); err != nil {
					routers.ErrorHelper(w, req, http.StatusBadRequest, err)
				}
			} else {
				if err := r.store.Product().EditInstance(property); err != nil {
					routers.ErrorHelper(w, req, http.StatusBadRequest, err)
				}
			}
		}
	}
}
