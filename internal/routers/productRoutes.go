package routers

import (
	"ShopRestAPI/internal/models/products"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (r *Routes) ConfigureProductRoutes(mux *http.ServeMux) {
	r.ConfigurePropertiesRoutes(mux)
	mux.HandleFunc("/product", r.CreateProductRouter())
	mux.HandleFunc("/products", r.CreateProductFilterRouter())
}

func (r *Routes) CreateProductRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		param := req.URL.Query().Get("id")
		id, err := strconv.ParseInt(param, 10, 32)
		if err != nil {
			errorHelper(w, req, http.StatusBadRequest, err)
		}
		if req.Method == "GET" {
			data, err := r.store.Product().Get(int(id))
			if err != nil {
				errorHelper(w, req, http.StatusBadRequest, err)
			}
			info, err := json.Marshal(data)
			fmt.Fprintf(w, "%s\n", info)
		} else if req.Method == "DELETE" {
			r.store.Product().Delete(int(id))
		} else if req.Method == "PUT" {
			var product = &products.Product{}
			if err := json.NewDecoder(req.Body).Decode(product); err != nil {
				errorHelper(w, req, http.StatusBadRequest, err)
			}
			if product.Id == 0 {
				err := r.store.Product().Create(product)
				if err != nil {
					errorHelper(w, req, http.StatusBadRequest, err)
				}
			} else {
				if err := r.store.Product().Edit(product); err != nil {
					errorHelper(w, req, http.StatusBadRequest, err)
				}
			}
		}
	}
}

func (r *Routes) CreateProductFilterRouter() http.HandlerFunc {
	type response struct {
		Data []products.Product `json:"data"`
	}

	return func(w http.ResponseWriter, req *http.Request) {

		var filter = &products.Filter{}
		if err := json.NewDecoder(req.Body).Decode(filter); err != nil {
			errorHelper(w, req, http.StatusBadRequest, err)
		}

		products, err := r.store.Product().ProductsSearch(filter)
		if err != nil {
			errorHelper(w, req, http.StatusBadRequest, err)
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
