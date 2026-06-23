package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/aliffatulmf/buku-tamu-apbj/app"
	"github.com/aliffatulmf/buku-tamu-apbj/database"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/handler"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/service"
	"github.com/aliffatulmf/buku-tamu-apbj/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed assets/*
var a embed.FS

//go:embed templates/*
var f embed.FS

var (
	StatusMode = gin.DebugMode
	Server     *gin.Engine
	DB         *gorm.DB
)

var (
	AppName  = "Buku Tamu"
	Version  = "1.0"
	Protocol = "HTTPS"
	Port     = "6170"
)

var run string

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

	flag.StringVar(&run, "run", "tls", "run service using TLS or HTTP\nOption: [\"tls\", \"http\"]")
	flag.Parse()

	switch run {
	case "tls":
		gin.SetMode(gin.ReleaseMode)
		Server = gin.New()
	case "http":
		gin.SetMode(gin.DebugMode)
		Server = gin.Default()
	}

	switch StatusMode {
	case gin.ReleaseMode:
		tmpl := template.Must(
			template.New("").Funcs(template.FuncMap{
				"increment": func(x int) int {
					return x + 1
				},
			}).ParseFS(f, "templates/**/*.html"),
		)

		Server.SetHTMLTemplate(tmpl)
		Server.GET("/assets/*filepath", func(ctx *gin.Context) {
			staticServer := http.FileServer(http.FS(a))
			staticServer.ServeHTTP(ctx.Writer, ctx.Request)
		})

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

	default:
		fmt.Println("ERROR: status mode not available.")
		os.Exit(1)
	}

	Server.Static("/media", "media/")
	Server.Use(middleware.WebviewMiddleware())
	Server.SetTrustedProxies(nil)
}

func Handler(app *app.App) {
	var (
		PemdaRepository    = repository.NewPemdaRepository(DB)
		PenyediaRepository = repository.NewPenyediaRepository(DB)
		InstansiRepository = repository.NewAgencyRepository(DB)
		PokjaRepository    = repository.NewPokjaRepository(DB)
		TujuanRepository   = repository.NewTujuanRepository(DB)
		ImageStorage       = repository.NewLocalImageStorage()
		Exporter           = repository.NewExcelExporter()
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
		app.Server.GET("/", DashbordHandler.DashbordIndex)
		app.Server.GET("/export", DashbordHandler.DashboardExport)
	}

	// Instansi
	{
		instansi := app.Server.Group("/instansi")
		instansi.GET("/registrasi", InstansiHandler.InstansiIndex)
		instansi.GET("/terdaftar", InstansiHandler.InstansiFind)
		instansi.GET("/terdaftar/:id", InstansiHandler.InstansiDetail)
		instansi.POST("/registrasi", InstansiHandler.InstansiCreate)
		instansi.POST("/terdaftar/:id", InstansiHandler.InstansiUpdate)
	}

	// Penyedia
	{
		penyedia := app.Server.Group("/penyedia")
		penyedia.GET("/registrasi", PenyediaHandler.PenyediaIndex)
		penyedia.GET("/terdaftar", PenyediaHandler.PenyediaList)
		penyedia.GET("/terdaftar/:id", PenyediaHandler.PenyediaDetail)
		penyedia.POST("/registrasi", PenyediaHandler.PenyediaCreate)
		penyedia.PUT("/terdaftar/:id", PenyediaHandler.PenyediaUpdatePermission)
		penyedia.DELETE("/terdaftar/:id", PenyediaHandler.PenyediaDelete)
	}

	// Pokja
	{
		pokja := app.Server.Group("/pokja")
		pokja.GET("/registrasi", PokjaHandler.PokjaIndex)
		pokja.GET("/terdaftar", PokjaHandler.PokjaList)
		pokja.GET("/terdaftar/:id", PokjaHandler.PokjaDetail)
		pokja.POST("/terdaftar/:id", PokjaHandler.PokjaUpdate)
		pokja.POST("/registrasi", PokjaHandler.PokjaCreate)
		pokja.DELETE("/terdaftar/:id", PokjaHandler.PokjaDelete)
	}

	// Pemda
	{
		pemda := app.Server.Group("/pemda")
		pemda.GET("/registrasi", PemdaHandler.PemdaIndex)
		pemda.GET("/terdaftar", PemdaHandler.PemdaList)
		pemda.GET("/terdaftar/:id", PemdaHandler.PemdaDetail)
		pemda.POST("/registrasi", PemdaHandler.PemdaCreate)
		pemda.PUT("/terdaftar/:id", PemdaHandler.UpdatePermission)
		pemda.DELETE("/terdaftar/:id", PemdaHandler.DeleteByID)
	}

	{
		app.Server.GET("/credits", func(ctx *gin.Context) {
			ctx.HTML(200, "credits.html", gin.H{
				"info": gin.H{
					"appname":  AppName,
					"version":  Version,
					"protocol": Protocol,
					"port":     Port,
				},
			})
		})
		app.Server.GET("/user", func(ctx *gin.Context) {
			ctx.HTML(200, "user-dashboard.html", gin.H{})
		})
	}
}

func main() {
	app := app.New(Server)
	app.SetHandler(Handler)
	app.SetPort(Port)

	switch {
	case run == "tls":
		app.RunTLS()
	case run == "http":
		app.RunHTTP()
	default:
		fmt.Println("command not recognized")
		os.Exit(1)
	}
}
