package productRouters

import (
	"ShopRestAPI/internal/models/products"
	"ShopRestAPI/internal/routers/helperRoters"
	"ShopRestAPI/internal/server"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (r *ProductRoutes) ConfigurePropertiesRoutes(s *server.Server) {
	s.Router.HandleFunc("/props/color", r.CreateColorRouter("color"))
	s.Router.HandleFunc("/props/country", r.CreateColorRouter("country"))
	s.Router.HandleFunc("/props/sizes", r.CreateColorRouter("sizes"))
	s.Router.HandleFunc("/props/styles", r.CreateColorRouter("styles"))
	s.Router.HandleFunc("/props/types", r.CreateColorRouter("types"))
	s.Router.HandleFunc("/props/season", r.CreateColorRouter("season"))
	s.Router.HandleFunc("/property", r.CreatePropertyRouter(s))
}

func (r *ProductRoutes) CreateColorRouter(table string) http.HandlerFunc {
	type response struct {
		Data []string `json:"data"`
	}
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := r.store.Product().GetPropertyList(table)
		if err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
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

func (r *ProductRoutes) CreatePropertyRouter(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "DELETE" {
			if err := s.AuthUser(w, req); err != nil {
				helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				return
			}
			param := req.URL.Query().Get("id")
			id, err := strconv.ParseInt(param, 10, 32)
			if err != nil {
				helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
			}
			r.store.Product().DeleteInstance(int(id))
		} else if req.Method == "PUT" {
			if err := s.AuthUser(w, req); err != nil {
				helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				return
			}
			var property = &products.Instance{}
			if err := json.NewDecoder(req.Body).Decode(property); err != nil {
				log.Fatal(err)
			}
			if property.Id == 0 {
				if err := r.store.Product().CreateInstance(property); err != nil {
					helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				}
			} else {
				if err := r.store.Product().EditInstance(property); err != nil {
					helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				}
			}
		}
	}
}
