package router

import (
	"JUALiND/helper"
	"JUALiND/models"
	"JUALiND/repository"
	"log"
	"net/http"
	"strconv"
)

type OrderStructCSV struct {
	BuyerName        string `csv:"buyer_name"`
	ProductName      string `csv:"product_name"`
	ProductPrice     int    `csv:"product_price"`
	Amount           int    `csv:"amount"`
	Date             string `csv:"date"`
	ConfirmationLink string `csv:"confirmation_link"`
}

func CreateOrder(orderRepo *repository.OrderRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var o models.Order
		var err error
		o.Amount, err = strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}
		o.BuyerID = int(r.Context().Value("user_detail").(models.Users).ID)
		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}
		o.ProductID, err = strconv.Atoi(r.FormValue("id_product"))

		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		err = orderRepo.CreateOrder(o)

		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success Creating User", http.StatusOK)

	})
}
func GetAllOrdersByUser(orderRepo *repository.OrderRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserData := r.Context().Value("user_detail").(models.Users)

		results, err := orderRepo.GetAllOrdersByUser(int(UserData.ID))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}

		helper.SuccessResponseJSON(w, "Success", results)
	})
}

func GenerateFileByUser(orderRepo *repository.OrderRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		UserData := r.Context().Value("user_detail").(models.Users)

		results, err := orderRepo.GetAllOrdersByUser(int(UserData.ID))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}
		var res []OrderStructCSV
		for _, v := range results {
			var o OrderStructCSV
			o.Amount = v.Amount
			o.BuyerName = v.BuyerName
			o.Date = v.Date
			o.ProductName = v.ProductName
			o.ProductPrice = v.ProductPrice
			if v.ConfirmationLink.Valid {
				o.ConfirmationLink = v.ConfirmationLink.String
			} else {
				o.ConfirmationLink = "Not Yet Confirm"
			}

			res = append(res, o)
		}

		fileName, err := helper.SaveToCSV(res)
		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", map[string]string{
			"filename": fileName,
		})
	})

}
