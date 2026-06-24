package main

import (
	"errors"
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
	Version = "2.1"
	Port    = "6170"
)

func init() {
	cdir := [3]string{"media/img", "Documents/Pemda", "Documents/Penyedia"}
	for _, c := range cdir {
		if _, err := os.Stat(c); errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(c, os.ModeDir); err != nil {
				fmt.Printf("FATAL: %s\n", err.Error())
				os.Exit(1)
			}
		}
	}

	debug := os.Getenv("BUKUTAMU_DEBUG")
	if debug == "1" {
		StatusMode = gin.DebugMode
		gin.SetMode(gin.DebugMode)
		Server = gin.Default()
	} else {
		StatusMode = gin.ReleaseMode
		gin.SetMode(gin.ReleaseMode)
		Server = gin.New()
	}

	switch StatusMode {
	case gin.ReleaseMode:
		Server.SetFuncMap(template.FuncMap{
			"increment": func(x int) int {
				return x + 1
			},
		})
		Server.LoadHTMLGlob("templates/**/*.html")
		Server.Static("/assets", "assets/")

		DB = database.NewConnection(&gorm.Config{
			Logger: logger.Default,
		})

	case gin.DebugMode:
		Server.SetFuncMap(template.FuncMap{
			"increment": func(x int) int {
				return x + 1
			},
		})
		Server.LoadHTMLGlob("templates/**/*.html")
		Server.Static("/assets", "assets/")

		DB = database.NewConnection(&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})

		DB.AutoMigrate(
			&entity.Destination{},
			&entity.Consultation{},
			&entity.Pokja{},
			&entity.Agency{},
			&entity.Pemda{},
			&entity.Provider{},
		)

		database.Seed(DB)

	default:
		fmt.Println("ERROR: status mode not available.")
		os.Exit(1)
	}

	Server.Static("/media", "media/")
	// Server.Use(middleware.WebviewMiddleware())
	Server.SetTrustedProxies(nil)
}

func Handler() {
	var (
		PemdaRepository    = repository.NewPemdaRepository(DB)
		PenyediaRepository = repository.NewPenyediaRepository(DB)
		InstansiRepository = repository.NewAgencyRepository(DB)
		PokjaRepository    = repository.NewPokjaRepository(DB)
		TujuanRepository   = repository.NewTujuanRepository(DB)
		ImageStorage       = io.NewImageStorage("media/img")
		Exporter           = io.NewExcelExporter()
	)

	var (
		DashboardService = service.NewDashboardServices(
			PemdaRepository,
			PenyediaRepository,
			PokjaRepository,
			InstansiRepository,
			Exporter,
		)
		PenyediaService = service.NewProviderService(PenyediaRepository, ImageStorage)
		TujuanService   = service.NewTujuanService(TujuanRepository)
		InstansiService = service.NewAgencyService(InstansiRepository)
		PokjaService    = service.NewPokjaService(PokjaRepository)
		PemdaService    = service.NewPemdaService(PemdaRepository, InstansiRepository, ImageStorage)
	)

	var (
		DashbordHandler = handler.NewDashboardHandler(DashboardService)
		PenyediaHandler = handler.NewPenyediaHandler(PenyediaService, TujuanService)
		PokjaHandler    = handler.NewPokjaHandler(PokjaService)
		InstansiHandler = handler.NewAgencyHandler(InstansiService)
		PemdaHandler    = handler.NewPemdaHandler(PemdaService, TujuanService, InstansiService)
	)

	// Dashboard
	{
		Server.GET("/", DashbordHandler.DashbordIndex)
		Server.GET("/export", DashbordHandler.DashboardExport)
	}

	// Instansi
	{
		instansi := Server.Group("/instansi")
		instansi.GET("/registrasi", InstansiHandler.InstansiIndex)
		instansi.GET("/terdaftar", InstansiHandler.InstansiFind)
		instansi.GET("/terdaftar/:id", InstansiHandler.InstansiDetail)
		instansi.POST("/registrasi", InstansiHandler.InstansiCreate)
		instansi.POST("/terdaftar/:id", InstansiHandler.InstansiUpdate)
	}

	// Penyedia
	{
		penyedia := Server.Group("/penyedia")
		penyedia.GET("/registrasi", PenyediaHandler.PenyediaIndex)
		penyedia.GET("/terdaftar", PenyediaHandler.PenyediaList)
		penyedia.GET("/terdaftar/:id", PenyediaHandler.PenyediaDetail)
		penyedia.POST("/registrasi", PenyediaHandler.PenyediaCreate)
		penyedia.PUT("/terdaftar/:id", PenyediaHandler.PenyediaUpdatePermission)
		penyedia.DELETE("/terdaftar/:id", PenyediaHandler.PenyediaDelete)
	}

	// Pokja
	{
		pokja := Server.Group("/pokja")
		pokja.GET("/registrasi", PokjaHandler.PokjaIndex)
		pokja.GET("/terdaftar", PokjaHandler.PokjaList)
		pokja.GET("/terdaftar/:id", PokjaHandler.PokjaDetail)
		pokja.POST("/terdaftar/:id", PokjaHandler.PokjaUpdate)
		pokja.POST("/registrasi", PokjaHandler.PokjaCreate)
		pokja.DELETE("/terdaftar/:id", PokjaHandler.PokjaDelete)
	}

	// Pemda
	{
		pemda := Server.Group("/pemda")
		pemda.GET("/registrasi", PemdaHandler.PemdaIndex)
		pemda.GET("/terdaftar", PemdaHandler.PemdaList)
		pemda.GET("/terdaftar/:id", PemdaHandler.PemdaDetail)
		pemda.POST("/registrasi", PemdaHandler.PemdaCreate)
		pemda.PUT("/terdaftar/:id", PemdaHandler.UpdatePermission)
		pemda.DELETE("/terdaftar/:id", PemdaHandler.DeleteByID)
	}

	{
		Server.GET("/credits", func(ctx *gin.Context) {
			ctx.HTML(200, "credits.html", gin.H{
				"info": gin.H{
					"appname": AppName,
					"version": Version,
					"port":    Port,
				},
			})
		})
		Server.GET("/user", func(ctx *gin.Context) {
			ctx.HTML(200, "user-dashboard.html", gin.H{})
		})
	}
}

func main() {
	Handler()

	port := Port
	if !strings.Contains(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	if err := Server.Run(port); err != nil {
		fmt.Println("FATAL:", err.Error())
	}
}
