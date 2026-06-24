package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/aliffatulmf/buku-tamu-apbj/database"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/handler"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/io"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/service"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	StatusMode string
	Server     *gin.Engine
	DB         *gorm.DB
)

const (
	AppName = "Buku Tamu"
	Version = "2.2"
	Port    = "6170"
)

func main() {
	cdirs := []string{"media/images", "Documents/Pemda", "Documents/Penyedia"}
	for _, dir := range cdirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModeDir); err != nil {
				fmt.Printf("FATAL: %s\n", err.Error())
				os.Exit(1)
			}
		}
	}

	if os.Getenv("BUKUTAMU_DEBUG") == "1" {
		StatusMode = gin.DebugMode
		gin.SetMode(gin.DebugMode)
		Server = gin.Default()
	} else {
		StatusMode = gin.ReleaseMode
		gin.SetMode(gin.ReleaseMode)
		Server = gin.New()
	}

	Server.SetFuncMap(template.FuncMap{
		"increment": func(x int) int { return x + 1 },
	})
	Server.LoadHTMLGlob("templates/**/*.html")
	Server.Static("/assets", "assets/")
	Server.Static("/media", "media/")
	Server.SetTrustedProxies(nil)

	dbConfig := gorm.Config{Logger: logger.Default}
	DB = database.NewConnection(&dbConfig)

	if !isClientIdValid() {
		DB.AutoMigrate(
			&entity.Destination{},
			&entity.Consultation{},
			&entity.Pokja{},
			&entity.Agency{},
			&entity.Pemda{},
			&entity.Provider{},
		)
		database.Seed(DB)
	}

	wire()

	port := Port
	if !strings.Contains(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}
	if err := Server.Run(port); err != nil {
		fmt.Println("FATAL:", err.Error())
	}
}

func wire() {
	PemdaRepository := repository.NewPemdaRepository(DB)
	PenyediaRepository := repository.NewPenyediaRepository(DB)
	InstansiRepository := repository.NewAgencyRepository(DB)
	PokjaRepository := repository.NewPokjaRepository(DB)
	TujuanRepository := repository.NewTujuanRepository(DB)
	ImageStorage := io.NewImageStorage("media/img")
	Exporter := io.NewExcelExporter()

	DashboardService := service.NewDashboardServices(PemdaRepository, PenyediaRepository, PokjaRepository, InstansiRepository, Exporter)
	PenyediaService := service.NewProviderService(PenyediaRepository, ImageStorage)
	TujuanService := service.NewTujuanService(TujuanRepository)
	InstansiService := service.NewAgencyService(InstansiRepository)
	PokjaService := service.NewPokjaService(PokjaRepository)
	PemdaService := service.NewPemdaService(PemdaRepository, InstansiRepository, ImageStorage)

	DashboardHandler := handler.NewDashboardHandler(DashboardService)
	PenyediaHandler := handler.NewPenyediaHandler(PenyediaService, TujuanService)
	PokjaHandler := handler.NewPokjaHandler(PokjaService)
	InstansiHandler := handler.NewAgencyHandler(InstansiService)
	PemdaHandler := handler.NewPemdaHandler(PemdaService, TujuanService, InstansiService)

	// Dashboard
	Server.GET("/", DashboardHandler.DashbordIndex)
	Server.GET("/export", DashboardHandler.DashboardExport)

	// Instansi
	instansi := Server.Group("/instansi")
	instansi.GET("/registrasi", InstansiHandler.InstansiIndex)
	instansi.GET("/terdaftar", InstansiHandler.InstansiFind)
	instansi.GET("/terdaftar/:id", InstansiHandler.InstansiDetail)
	instansi.POST("/registrasi", InstansiHandler.InstansiCreate)
	instansi.POST("/terdaftar/:id", InstansiHandler.InstansiUpdate)

	// Penyedia
	penyedia := Server.Group("/penyedia")
	penyedia.GET("/registrasi", PenyediaHandler.PenyediaIndex)
	penyedia.GET("/terdaftar", PenyediaHandler.PenyediaList)
	penyedia.GET("/terdaftar/:id", PenyediaHandler.PenyediaDetail)
	penyedia.POST("/registrasi", PenyediaHandler.PenyediaCreate)
	penyedia.PUT("/terdaftar/:id", PenyediaHandler.PenyediaUpdatePermission)
	penyedia.DELETE("/terdaftar/:id", PenyediaHandler.PenyediaDelete)

	// Pokja
	pokja := Server.Group("/pokja")
	pokja.GET("/registrasi", PokjaHandler.PokjaIndex)
	pokja.GET("/terdaftar", PokjaHandler.PokjaList)
	pokja.GET("/terdaftar/:id", PokjaHandler.PokjaDetail)
	pokja.POST("/terdaftar/:id", PokjaHandler.PokjaUpdate)
	pokja.POST("/registrasi", PokjaHandler.PokjaCreate)
	pokja.DELETE("/terdaftar/:id", PokjaHandler.PokjaDelete)

	// Pemda
	pemda := Server.Group("/pemda")
	pemda.GET("/registrasi", PemdaHandler.PemdaIndex)
	pemda.GET("/terdaftar", PemdaHandler.PemdaList)
	pemda.GET("/terdaftar/:id", PemdaHandler.PemdaDetail)
	pemda.POST("/registrasi", PemdaHandler.PemdaCreate)
	pemda.PUT("/terdaftar/:id", PemdaHandler.UpdatePermission)
	pemda.DELETE("/terdaftar/:id", PemdaHandler.DeleteByID)

	// Other pages
	Server.GET("/credits", func(ctx *gin.Context) {
		ctx.HTML(200, "credits.html", gin.H{
			"info": gin.H{"appname": AppName, "version": Version, "port": Port},
		})
	})
	Server.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(200, "user-dashboard.html", gin.H{})
	})
}

func isClientIdValid() bool {
	const clientIdFile = ".client"

	data, err := os.ReadFile(clientIdFile)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(clientIdFile, []byte(uuid.NewString()), 0644); err != nil {
				fmt.Printf("FATAL: cannot create %s: %s\n", clientIdFile, err.Error())
				os.Exit(1)
			}
			return true
		}
		fmt.Printf("FATAL: cannot read %s: %s\n", clientIdFile, err.Error())
		os.Exit(1)
	}

	if _, err := uuid.Parse(strings.TrimSpace(string(data))); err != nil {
		fmt.Printf("FATAL: invalid client ID in %s\n", clientIdFile)
		os.Exit(1)
	}

	return true
}
