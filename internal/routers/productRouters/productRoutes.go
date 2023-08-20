package productRouters

import (
	"ShopRestAPI/internal/Storage"
	"ShopRestAPI/internal/models/products"
	"ShopRestAPI/internal/routers/helperRoters"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ProductRoutes struct {
	store Storage.Store
}

func ConfigureProductRoutes(mux *http.ServeMux, store Storage.Store) {
	r := ProductRoutes{store: store}
	r.ConfigurePropertiesRoutes(mux)
	mux.HandleFunc("/product", r.CreateProductRouter())
	mux.HandleFunc("/products", r.CreateProductFilterRouter())
}

func (r *ProductRoutes) CreateProductRouter() http.HandlerFunc {
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
			r.store.Product().Delete(int(id))
		} else if req.Method == "PUT" {
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
