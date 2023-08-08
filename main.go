package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	docs "sireg/rest-api-kegiatan/docs"
	"sireg/rest-api-kegiatan/internal/kategori"
	"sireg/rest-api-kegiatan/internal/kegiatan"
	"sireg/rest-api-kegiatan/internal/repository"
	"sireg/rest-api-kegiatan/pkg/dbcontext"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 		SIREG Kegiatan Kategori
// @version 	0.1
// @description	Endpoint SIREG Kegiatan Kategori

// @contact.name 	anhilmy
// @contact.email 	hilmyahmadnaufal@gmail.com

// @host 		localhost:8080
// @BasePath 	/api/1

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		os.Exit(-1)
	}

	fmt.Println(os.Getenv("DATABASE_URL"))
	db, err := dbx.MustOpen("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Cannot connect to database", err.Error())
		os.Exit(-1)
	}

	db.LogFunc = log.Printf

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("Error when closing the database", err.Error())
		}
	}()

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router = buildHandler(dbcontext.New(db), router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error when running", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func buildHandler(db *dbcontext.DB, router *gin.Engine) *gin.Engine {

	kategoriRepo := repository.NewKategoriRepo(db)
	kegiatanRepo := repository.NewKegiatanRepo(db)

	rv1 := router.Group("/api/v1")
	kategori.RegisterHandler(rv1.Group("/kategori"),
		kategori.NewService(kategoriRepo),
	)

	kegiatan.RegisterHandler(rv1.Group("/kegiatan"),
		kegiatan.NewService(kegiatanRepo),
	)

	return router
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do piing
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {Object} helloword
// @Router / [get]
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}
