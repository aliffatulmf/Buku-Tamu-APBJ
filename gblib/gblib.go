package main

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	gblib "github.com/aliffatulmf/buku-tamu-apbj/gblib/pkg"
	"flag"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	tables := []interface{}{&entity.Agency{}, &entity.Destination{}, &entity.Consultation{}, &entity.Pokja{}, &entity.Pemda{}, &entity.Provider{}}
	models := []interface{}{gblib.InstansiModel, gblib.KonsultasiModel, gblib.TujuanModel, gblib.PokjaModel}
	db, err := gorm.Open(sqlite.Open("good.db"), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		fmt.Println("fatal:", err.Error())
		os.Exit(1)
	}

	var cmd, cert string
	flag.StringVar(&cmd, "cmd", "", "[\"migrate\", \"drop\"]")
	flag.StringVar(&cert, "cert", "", "[\"renew\"]")
	flag.Parse()

	switch cmd {
	case "drop":
		fmt.Println("dropping tables.")
		if err := db.Migrator().DropTable(tables...); err != nil {
			fmt.Println("fatal:", err.Error())
			os.Exit(1)
		}
	case "migrate":
		fmt.Println("migrating tables.")
		if err := db.AutoMigrate(tables...); err != nil {
			fmt.Println("fatal:", err.Error())
			os.Exit(1)
		}

		fmt.Println("starting to insert data.")
		var wg sync.WaitGroup

		fmt.Print("running ")
		for i, m := range models {
			wg.Add(1)

			go func(idx int, model interface{}) {
				defer wg.Done()
				fmt.Print(".")
				db.Create(model)
			}(i, m)
		}
		wg.Wait()
		fmt.Println("\nsuccess inserted data into database.")
	}

	switch cert {
	case "renew":
		fmt.Println("renew certificate in current path.")
		gblib.GenerateCert()
	}

	fmt.Println("done.")
}
