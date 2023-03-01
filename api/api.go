package api

import (
	"context"
	"encoding/json"
	"github.com/gusleein/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goTest/db"
	"goTest/helpers"
	"goTest/models"
	"net/http"
)

type server struct {
	ctx context.Context
}

func Start(ctx context.Context) {
	s := &server{ctx: ctx}
	http.HandleFunc("/UserInfo", s.userInfoHandler)
	http.HandleFunc("/Top100", s.getTop100Handler)
	http.HandleFunc("/Top100With2Fields", s.getTop100With2FieldsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}
func (s *server) userInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		s.putUserInfoHandler(w, r)
	} else if r.Method == http.MethodGet {
		s.getUserInfoHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Добавление в бд юзера
func (s *server) putUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	var uI models.UserInfo
	err := json.NewDecoder(r.Body).Decode(&uI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Info(err)
		return
	}
	uI.ID = primitive.NewObjectID()
	_, err = db.Collection.InsertOne(s.ctx, uI)
	if err != nil {
		log.Error(err)
	}
	w.WriteHeader(http.StatusCreated)
}

// Достаём юзера из бд
func (s *server) getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	if idString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(idString)
	var result bson.M
	err = db.Collection.FindOne(s.ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Error(err)
	}
	if result == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	userInfoJSON, err := json.Marshal(result)
	if err != nil {
		log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userInfoJSON)
	w.WriteHeader(http.StatusOK)
}

// Top 100
func (s *server) getTop100Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	var result string

	// создание пайплайнов агрегации для поиска топ-100 наиболее популярных значений каждого поля
	osName := helpers.PiplineTop100("os_name")
	browserName := helpers.PiplineTop100("browser_name")
	phoneBrand := helpers.PiplineTop100("phone_brand")
	screenRes := helpers.PiplineTop100("screen_res")
	result += "OS:\n"
	result += helpers.StringFromPipeline(s.ctx, osName)
	result += "Браузеры:\n"
	result += helpers.StringFromPipeline(s.ctx, browserName)
	result += "Марка телефона:\n"
	result += helpers.StringFromPipeline(s.ctx, phoneBrand)
	result += "Разрешение экрана:\n"
	result += helpers.StringFromPipeline(s.ctx, screenRes)
	w.Write([]byte(result))
}

func (s *server) getTop100With2FieldsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	osVerOsName := helpers.PiplineTop100WithTwoFields("os_name", "os_version")
	result := helpers.StringFromPipelineWith2Field(s.ctx, osVerOsName)
	w.Write(result)
}
