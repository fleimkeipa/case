package main

import (
	"log"
	"os"

	"github.com/fleimkeipa/case/controller"
	"github.com/fleimkeipa/case/pkg"
	"github.com/fleimkeipa/case/repositories"
	"github.com/fleimkeipa/case/uc"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	// Start the application
	serveApplication()
}

func serveApplication() {
	// init config
	loadConfig()

	// Create a new Echo instance
	e := echo.New()

	// Configure Echo settings
	configureEcho(e)

	// Configure CORS middleware
	configureCORS(e)

	// Configure the logger
	sugar := configureLogger(e)
	defer sugar.Sync() // Clean up logger at the end

	// Initialize PostgreSQL client
	dbClient := initDB()
	defer dbClient.Close() // Clean up db connections at the end

	httpClient := pkg.NewHTTPClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	productDBRepo := repositories.NewProductDBRepository(dbClient)
	productDBUC := uc.NewProductDBUC(productDBRepo)

	productAPIRepo := repositories.NewProductAPIRepository(httpClient)
	productAPIUC := uc.NewProductAPIUC(productAPIRepo, *productDBUC)

	productController := controller.NewProductController(productAPIUC)

	// Register routes
	e.GET("/products", productController.FindAll)
	e.GET("/products/:id", productController.FindOne)

	e.Logger.Fatal(e.Start(":8080"))
}

func loadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// Configures the Echo instance
func configureEcho(e *echo.Echo) {
	e.HideBanner = true
	e.HidePort = true

	// Add Recover middleware
	e.Use(middleware.Recover())
}

// Configures CORS settings
func configureCORS(e *echo.Echo) {
	corsConfig := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	})

	e.Use(corsConfig)
}

// Configures the logger and adds it as middleware
func configureLogger(e *echo.Echo) *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	e.Use(pkg.ZapLogger(logger))

	sugar := logger.Sugar()
	loggerHandler := controller.NewLogger(sugar)
	e.Use(loggerHandler.LoggerMiddleware)

	return sugar
}

// Initializes the PostgreSQL client
func initDB() *pg.DB {
	db := pkg.NewPSQLClient()
	if db == nil {
		log.Fatal("Failed to initialize PostgreSQL client")
	}

	log.Println("PostgreSQL client initialized successfully")
	return db
}
