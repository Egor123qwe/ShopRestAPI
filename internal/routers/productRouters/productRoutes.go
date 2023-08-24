package productRouters

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/models/products"
	"ShopRestAPI/internal/routers/helperRoters"
	"ShopRestAPI/internal/server"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ProductRoutes struct {
	store Storage.Store
}

func ConfigureProductRoutes(s *server.Server) {
	r := ProductRoutes{store: s.Store}
	r.ConfigurePropertiesRoutes(s)
	s.Router.HandleFunc("/product", r.CreateProductRouter(s))
	s.Router.HandleFunc("/products", r.CreateProductFilterRouter())
}

func (r *ProductRoutes) CreateProductRouter(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		param := req.URL.Query().Get("id")
		id, err := strconv.ParseInt(param, 10, 32)
		if err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
		}
		if req.Method == "GET" {
			data, err := r.store.Product().Get(int(id))
			if err != nil {
				helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
			}
			info, err := json.Marshal(data)
			fmt.Fprintf(w, "%s\n", info)
		} else if req.Method == "DELETE" {
			if err := s.AuthUser(w, req); err != nil {
				helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				return
			}
			r.store.Product().Delete(int(id))
		} else if req.Method == "PUT" {
			if err := s.AuthUser(w, req); err != nil {
				helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				return
			}
			var product = &products.Product{}
			if err := json.NewDecoder(req.Body).Decode(product); err != nil {
				helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
			}
			if product.Id == 0 {
				err := r.store.Product().Create(product)
				if err != nil {
					helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				}
			} else {
				if err := r.store.Product().Edit(product); err != nil {
					helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
				}
			}
		}
	}
}

func (r *ProductRoutes) CreateProductFilterRouter() http.HandlerFunc {
	type response struct {
		Data []products.Product `json:"data"`
	}

	return func(w http.ResponseWriter, req *http.Request) {

		var filter = &products.Filter{}
		if err := json.NewDecoder(req.Body).Decode(filter); err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
		}

		products, err := r.store.Product().ProductsSearch(filter)
		if err != nil {
			helperRoters.ErrorHelper(w, req, http.StatusBadRequest, err)
		}
		data := response{
			Data: products,
		}
		info, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s\n", info)
	}
}
