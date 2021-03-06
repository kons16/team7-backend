package web

import (
	"github.com/kons16/team7-backend/usecase"
	"github.com/kons16/team7-backend/web/handler"
	"net/http"
)

func NewServer(userUC *usecase.UserUseCase, sessionUC *usecase.SessionUseCase, clothUC *usecase.ClothUseCase, cordiUC *usecase.CordinateUseCase) *http.Server {
	var s http.Server
	s.Addr = ":8000"

	userHandler := handler.NewUserHandler(userUC)
	sessionHandler := handler.NewSessionHandler(sessionUC)
	clothHandler := handler.NewClothHandler(clothUC, sessionUC)
	cordinateHandler := handler.NewCordinateHandler(cordiUC, sessionUC)

	// GET サーバーが立ち上がっているか確認
	http.HandleFunc("/api/v1/hello", handler.Hello)
	// POST アカウントの新規登録し, sessionID を返す
	http.HandleFunc("/api/v1/user", userHandler.CreateUser)
	// POST user_sessionテーブル に sessionID に紐づく userID のカラムを追加し, sessionID を返す
	http.HandleFunc("/api/v1/login", userHandler.Login)
	// GET クライアントから送られてきた sessionID が切れてないか確認
	http.HandleFunc("/api/v1/is_login", sessionHandler.FindUserBySession)
	// GET user_sessionテーブル から sessionID のカラムを削除する
	http.HandleFunc("/api/v1/logout", sessionHandler.Logout)

	// POST 服情報の追加
	http.HandleFunc("/api/v1/cloth", clothHandler.CreateCloth)
	// GET 服情報の取得
	http.HandleFunc("/api/v1/get_cloth", clothHandler.GetAll)
	// POST 服の購入。 購入した服の ID を ユーザーと紐付ける
	http.HandleFunc("/api/v1/buy", clothHandler.BuyCloth)
	// GET 購入した服の情報を持ってくる
	http.HandleFunc("/api/v1/my_cloth", clothHandler.GetBuyCloth)

	// POST コーディネート情報の追加
	http.HandleFunc("/api/v1/cordinate", cordinateHandler.CreateCordinate)
	// GET 自分のコーディネートの取得
	http.HandleFunc("/api/v1/get_cordinate", cordinateHandler.Get)

	return &s
}
