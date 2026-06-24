package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
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

const (
	AppName = "Buku Tamu"
	Version = "2.4"
	Port    = "6170"

	defaultImagePath = "media/images"

	envKeyDebug = "BUKUTAMU_DEBUG"
)

//go:embed assets/*
var assetsFS embed.FS

//go:embed templates/*
var templatesFS embed.FS

func initGin() *gin.Engine {
	var r *gin.Engine
	if os.Getenv(envKeyDebug) == "1" {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	return r
}

func initDB() *gorm.DB {
	return database.NewConnection(&gorm.Config{Logger: logger.Default})
}

func initDirs() {
	dirs := []string{"media/images", "Documents/Pemda", "Documents/Penyedia"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModeDir); err != nil {
			fmt.Printf("FATAL: %s\n", err.Error())
			os.Exit(1)
		}
	}
}

func main() {
	initDirs()

	db := initDB()
	r := initGin()

	assetFS, err := fs.Sub(assetsFS, "assets")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("/assets", http.FS(assetFS))

	subTmplFS, err := fs.Sub(templatesFS, "templates")
	if err != nil {
		log.Fatal(err)
	}

	tmplFiles, err := fs.Glob(subTmplFS, "*/*.html")
	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"start1": func(x int) int { return x + 1 },
	}

	tmpl := template.Must(template.New("templates").Funcs(funcMap).ParseFS(subTmplFS, tmplFiles...))
	r.SetHTMLTemplate(tmpl)
	r.Static("/media", "media/")
	r.SetTrustedProxies(nil)

	if !isClientIdValid() {
		db.AutoMigrate(
			&entity.Destination{},
			&entity.Consultation{},
			&entity.Pokja{},
			&entity.Agency{},
			&entity.Pemda{},
			&entity.Provider{},
		)
		database.Seed(db)
	}

	wire(r, db)

	port := Port
	if !strings.Contains(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}
	if err := r.Run(port); err != nil {
		fmt.Println("FATAL:", err.Error())
	}
}

func wire(r *gin.Engine, db *gorm.DB) {
	PemdaRepository := repository.NewPemdaRepository(db)
	PenyediaRepository := repository.NewPenyediaRepository(db)
	InstansiRepository := repository.NewAgencyRepository(db)
	PokjaRepository := repository.NewPokjaRepository(db)
	TujuanRepository := repository.NewTujuanRepository(db)
	ImageStorage := io.NewImageStorage(defaultImagePath)
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
	r.GET("/", DashboardHandler.DashbordIndex)
	r.GET("/export", DashboardHandler.DashboardExport)

	// Instansi
	instansi := r.Group("/instansi")
	instansi.GET("/registrasi", InstansiHandler.InstansiIndex)
	instansi.GET("/terdaftar", InstansiHandler.InstansiFind)
	instansi.GET("/terdaftar/:id", InstansiHandler.InstansiDetail)
	instansi.POST("/registrasi", InstansiHandler.InstansiCreate)
	instansi.POST("/terdaftar/:id", InstansiHandler.InstansiUpdate)

	// Penyedia
	penyedia := r.Group("/penyedia")
	penyedia.GET("/registrasi", PenyediaHandler.PenyediaIndex)
	penyedia.GET("/terdaftar", PenyediaHandler.PenyediaList)
	penyedia.GET("/terdaftar/:id", PenyediaHandler.PenyediaDetail)
	penyedia.POST("/registrasi", PenyediaHandler.PenyediaCreate)
	penyedia.PUT("/terdaftar/:id", PenyediaHandler.PenyediaUpdatePermission)
	penyedia.DELETE("/terdaftar/:id", PenyediaHandler.PenyediaDelete)

	// Pokja
	pokja := r.Group("/pokja")
	pokja.GET("/registrasi", PokjaHandler.PokjaIndex)
	pokja.GET("/terdaftar", PokjaHandler.PokjaList)
	pokja.GET("/terdaftar/:id", PokjaHandler.PokjaDetail)
	pokja.POST("/terdaftar/:id", PokjaHandler.PokjaUpdate)
	pokja.POST("/registrasi", PokjaHandler.PokjaCreate)
	pokja.DELETE("/terdaftar/:id", PokjaHandler.PokjaDelete)

	// Pemda
	pemda := r.Group("/pemda")
	pemda.GET("/registrasi", PemdaHandler.PemdaIndex)
	pemda.GET("/terdaftar", PemdaHandler.PemdaList)
	pemda.GET("/terdaftar/:id", PemdaHandler.PemdaDetail)
	pemda.POST("/registrasi", PemdaHandler.PemdaCreate)
	pemda.PUT("/terdaftar/:id", PemdaHandler.UpdatePermission)
	pemda.DELETE("/terdaftar/:id", PemdaHandler.DeleteByID)

	// Other pages
	r.GET("/credits", func(ctx *gin.Context) {
		ctx.HTML(200, "credits.html", gin.H{
			"info": gin.H{"appname": AppName, "version": Version, "port": Port},
		})
	})
	r.GET("/user", func(ctx *gin.Context) {
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
