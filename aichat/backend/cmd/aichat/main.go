package main

import (
	"aichat/internal/app/auth"
	"aichat/internal/app/chat"
	"aichat/pkg/mysql"
	"aichat/pkg/sqlm"
	"log"
	"net/http"

	"aichat/web"
	"aichat/web/handlers"
)

func main() {
	db, err := mysql.NewConnection()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	// defer db.Close()

	sqlmService := sqlm.NewSQLM(db)
	chatRepo := chat.NewRepository()
	chatService := chat.NewService(chatRepo)
	authRepo := auth.NewRepository(db, &sqlmService)
	authService := auth.NewService(authRepo)
	chatHandler := handlers.NewChatHandler(chatService, authService)

	r := web.NewRouter(chatHandler)

	log.Println("Service in port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}
