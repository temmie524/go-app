package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ClothParams struct {
	Id            string    `json:"id"`
	BrandName     string    `json:"cloth_name"`
	ClothCategory string    `json:"cloth_category"`
	StoreName     string    `json:"store_name"`
	Price         int       `json:"price"`
	Memo          string    `json:"memo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

// itemsがデータベースの役割を果たす。ここを改変する。
var clothes []*ClothParams

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the go app server")
	fmt.Println("Root endpoint is hooked!")
}

func getAllClothes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(clothes)
}

func getSingleCloth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, close := range clothes {
		if close.Id == key {
			json.NewEncoder(w).Encode(close)
		}
	}

}

func createCloth(w http.ResponseWriter, r *http.Request) {
	reqbody, _ := io.ReadAll(r.Body)
	var cloth ClothParams

	if err := json.Unmarshal(reqbody, &cloth); err != nil {
		log.Fatal(err)
	}

	//DN化するときは以下を変更
	clothes = append(clothes, &cloth)
	json.NewEncoder(w).Encode(cloth)

}

func deleteCloth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, cloth := range clothes {
		if cloth.Id == id {
			clothes = append(clothes[:i], clothes[i+1:]...)
		}
	}
}

func updateClothes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := io.ReadAll(r.Body)
	var updateCloth ClothParams
	if err := json.Unmarshal(reqBody, &updateCloth); err != nil {
		log.Fatal(err)
	}

	for i, cloth := range clothes {
		if cloth.Id == id {
			clothes[i] = &ClothParams{
				Id:            cloth.Id,
				BrandName:     updateCloth.BrandName,
				ClothCategory: updateCloth.ClothCategory,
				StoreName:     updateCloth.StoreName,
				Price:         updateCloth.Price,
				Memo:          updateCloth.Memo,
				CreatedAt:     cloth.CreatedAt,
				UpdatedAt:     updateCloth.UpdatedAt,
				DeletedAt:     cloth.DeletedAt,
			}
		}
	}
}

func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", rootPage)
	router.HandleFunc("/clothes", getAllClothes).Methods("GET")
	router.HandleFunc("/clothes/{id}", getSingleCloth).Methods("GET")
	router.HandleFunc("/cloth", createCloth).Methods("POST")
	router.HandleFunc("/cloth/{id}", deleteCloth).Methods("DELETE")
	router.HandleFunc("/cloth/{id}", updateClothes).Methods("PUT")

	return http.ListenAndServe(fmt.Sprintf("%d", 8080), router)

}

func init() {
	clothes = []*ClothParams{
		&ClothParams{
			Id:            "1",
			BrandName:     "YohjiYamamoto",
			ClothCategory: "Pants",
			StoreName:     "Nubian",
			Price:         30000,
			Memo:          "とても好き",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			DeletedAt:     time.Now(),
		},
		&ClothParams{
			Id:            "2",
			BrandName:     "IsseyMiyake",
			ClothCategory: "t-shirt",
			StoreName:     "arknets",
			Price:         25000,
			Memo:          "good! My favorite shirt!!",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			DeletedAt:     time.Now(),
		},
		&ClothParams{
			Id:            "3",
			BrandName:     "SoshiOtsuki",
			ClothCategory: "Jacket",
			StoreName:     "Present by friend",
			Price:         82000,
			Memo:          "expensive",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			DeletedAt:     time.Now(),
		},
	}
}
