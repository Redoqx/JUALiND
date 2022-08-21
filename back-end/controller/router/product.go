package router

import (
	"JUALiND/helper"
	"JUALiND/models"
	"JUALiND/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateProduct(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		price, err_p := strconv.Atoi(r.FormValue("price"))
		CurrentQuantity, err_c_q := strconv.Atoi(r.FormValue("cur_quantity"))
		Quantity, err_q := strconv.Atoi(r.FormValue("quantity"))
		Desc := r.FormValue("desc")
		imageFile, imageHeader, err := r.FormFile("image")

		if len(name) < 1 ||
			err != nil ||
			err_c_q != nil ||
			err_p != nil ||
			err_q != nil ||
			price <= 0 ||
			CurrentQuantity <= 0 ||
			Quantity <= 0 {
			helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "Some of the field has a wrong value", http.StatusBadRequest)
			return
		}

		if CurrentQuantity > Quantity {
			helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "cur_quantity cannot be larger than quantity", http.StatusBadRequest)
			return
		}

		UserData := r.Context().Value("user_detail").(models.Users)

		var p models.Product
		if err != nil {
			// Langsung save user tanpa gambar
			p.Name = name
			p.Price = uint(price)
			p.CurrentQuantity = uint(CurrentQuantity)
			p.Quantity = uint(Quantity)
			p.Description = Desc
			p.ImageLoc = helper.StringToNullString("")
			p.OwnerID = UserData.ID
			err = productRepo.CreateProduct(p)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			}
			helper.SuccessResponseJSON(w, "success", nil)
		} else {
			if !helper.ImageIsJpgOrPng(imageHeader) {
				helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "image must be in png or jpg format", http.StatusBadRequest)
				return
			}
			// save user serta gambarnya
			p.Name = name
			p.Price = uint(price)
			p.Description = Desc
			fileLocation := helper.UploadImage(imageFile, imageHeader)
			if fileLocation == "" {
				helper.ErrorResponseJSON(w, fmt.Errorf("error : error while saving image"), "Internal Server Error", http.StatusInternalServerError)
			}
			p.ImageLoc = helper.StringToNullString(r.Host + "/" + fileLocation)
			p.CurrentQuantity = uint(CurrentQuantity)
			p.Quantity = uint(Quantity)
			p.OwnerID = UserData.ID
			err = productRepo.CreateProduct(p)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
				helper.RemoveFile(r, fileLocation)
				return
			}
			helper.SuccessResponseJSON(w, "success", nil)
		}
	})
}

func UpdateProduct(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err_id := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		price, err_p := strconv.Atoi(r.FormValue("price"))
		CurrentQuantity, err_c_q := strconv.Atoi(r.FormValue("cur_quantity"))
		Quantity, err_q := strconv.Atoi(r.FormValue("quantity"))
		Desc := r.FormValue("desc")
		imageFile, imageHeader, err := r.FormFile("image")

		if len(name) < 1 ||
			err_c_q != nil ||
			err_p != nil ||
			err_q != nil ||
			err_id != nil {
			helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "Some of the field has a wrong value", http.StatusBadRequest)
			return
		}

		if CurrentQuantity > Quantity {
			helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "cur_quantity cannot be larger than quantity", http.StatusBadRequest)
			return
		}
		currentProduct, _ := productRepo.GetProduct(id)

		var p models.Product
		if err != nil {
			// Langsung save user tanpa gambar
			p.ID = uint(id)
			p.Name = name
			p.Price = uint(price)
			p.CurrentQuantity = uint(CurrentQuantity)
			p.Quantity = uint(Quantity)
			p.Description = Desc
			p.ImageLoc = helper.StringToNullString(currentProduct.ImageLoc.String)
			err = productRepo.UpdateProduct(p)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			helper.SuccessResponseJSON(w, "success", nil)
		} else {
			if !helper.ImageIsJpgOrPng(imageHeader) {
				helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "image must be in png or jpg format", http.StatusBadRequest)
				return
			}
			// save user serta gambarnya
			p.ID = uint(id)
			p.Name = name
			p.Price = uint(price)
			p.Description = Desc
			fileLocation := helper.UploadImage(imageFile, imageHeader)
			if fileLocation == "" {
				helper.ErrorResponseJSON(w, fmt.Errorf("error : error while saving image"), "Internal Server Error", http.StatusInternalServerError)
			}
			p.ImageLoc = helper.StringToNullString(r.Host + "/" + fileLocation)
			p.CurrentQuantity = uint(CurrentQuantity)
			p.Quantity = uint(Quantity)

			if currentProduct.ImageLoc.Valid {
				helper.RemoveFile(r, currentProduct.ImageLoc.String)
			}
			log.Println(r.Host)
			err = productRepo.UpdateProduct(p)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
				helper.RemoveFile(r, fileLocation)
				return
			}
			helper.SuccessResponseJSON(w, "success", nil)
		}
	})
}
func GetAllProduct(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := productRepo.GetProducts()

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", products)
	})
}

func GetProductID(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		product, err := productRepo.GetProduct(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}

		helper.SuccessResponseJSON(w, "Success", product)
	})
}

func DeleteProductByID(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		product, err := productRepo.GetProduct(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		if product.ImageLoc.Valid {
			log.Println("detele : ", product.ImageLoc.String)
			if err := helper.RemoveFile(r, product.ImageLoc.String); err != nil {
				log.Println(err.Error())
			}

		}

		err = productRepo.DeleteProduct(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", nil)

	})
}

func GetProductByUser(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		UserData := r.Context().Value("user_detail").(models.Users)

		product, err := productRepo.GetProductByUser(int(UserData.ID))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}

		helper.SuccessResponseJSON(w, "Success", product)

	})
}

func GetProductByName(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name := strings.Trim(r.URL.Query().Get("name"), " ")

		if name == "" {
			GetAllProduct(productRepo).ServeHTTP(w, r)
			return
		}

		product, err := productRepo.GetProductByName(name)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}

		helper.SuccessResponseJSON(w, "Success", product)
	})
}
