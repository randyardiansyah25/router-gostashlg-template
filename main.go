package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"router-gostashlg-template/delivery/consumer"
	"router-gostashlg-template/delivery/http/router"
	"router-gostashlg-template/entities/app"
	"router-gostashlg-template/entities/common"
	"router-gostashlg-template/entities/common/logger"
	"router-gostashlg-template/repository/built_in/broker"
	"router-gostashlg-template/repository/built_in/databasefactory"
	"router-gostashlg-template/repository/built_in/keyvaluefactory"

	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/randyardiansyah25/gostashlg"
	"github.com/randyardiansyah25/libpkg/util/env"
)

func main() {
	broker.IsUse = env.GetBool("rabbit.use", false)
	if broker.IsUse {
		broker.ConnectToRabbit()
		go broker.BrokerClosedChannelObserver()

		// For Consumer use case
		go consumer.Start()
	}
	router.Start()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger.Logger, _ = gostashlg.UseDefine(gostashlg.NewTemplate().
		Add(gostashlg.LOG, "{{.Data}}").
		Add(gostashlg.ERROR, "{{.Data}}").
		Add(gostashlg.WARN, "{{.Data}}").
		Add(gostashlg.INFO, "{{.Data}}"),
	)

	fmt.Println(`
 ____             _              _____                    _       _       
|  _ \ ___  _   _| |_ ___ _ __  |_   _|__ _ __ ___  _ __ | | __ _| |_ ___ 
| |_) / _ \| | | | __/ _ \ '__|   | |/ _ \ '_ ' _ \| '_ \| |/ _' | __/ _ \
|  _ < (_) | |_| | ||  __/ |      | |  __/ | | | | | |_) | | (_| | ||  __/
|_| \_\___/ \__,_|\__\___|_|      |_|\___|_| |_| |_| .__/|_|\__,_|\__\___|
                                                   |_|                                                                                                           
	`)
	logger.PrintInfof("%s : Ver.%s", app.Identifier, app.Version)
	LoadConfiguration(false)
	//* logger menggunakan gostashlg, tidak perlu rotation lagi karena sudah berbeda file jika ganti hari
	// go StartLoggerRotation()

	if os.Getenv("app.database_driver") != "" {
		PrepareRepo()
	}

	go ReloadObserver()
}

func LoadConfiguration(isReload bool) {
	var er error

	if isReload {
		logger.PrintLog("Reloading configuration file...")
		er = godotenv.Overload(".env")
	} else {
		logger.PrintLog("Loading configuration file...")
		er = godotenv.Load(".env")
	}

	if er != nil {
		logger.PrintError("Configuration file not found...")
		os.Exit(1)
	}

}

func PrepareRepo() {
	var er error
	if os.Getenv("app.database_driver") != "" {
		databasefactory.AppDb, er = databasefactory.GetDatabase()
		if er != nil {
			logger.PrintError("Fatal Error : ", er.Error())
			os.Exit(99)
		}

		logger.PrintLog("Connecting to database...")
		if er = databasefactory.AppDb.Connect(); er != nil {
			logger.PrintError("Connection to database failed : ", er.Error())
			os.Exit(99)
		}

		if er = databasefactory.AppDb.Ping(); er != nil {
			logger.PrintError("Cannot ping database : ", er.Error())
			os.Exit(99)
		}

		logger.PrintLog("Database Connected")
	}

	if os.Getenv("app.keyvalue_driver") != "" {
		keyvaluefactory.AppStore, er = keyvaluefactory.GetStore()
		if er != nil {
			logger.PrintError("Fatal Error : ", er.Error())
			os.Exit(99)
		}
		logger.PrintLog("Connecting to keyvalue store...")
		if er = keyvaluefactory.AppStore.Open(); er != nil {
			logger.PrintError("failed to open keyvalue store : ", er.Error())
			os.Exit(99)
		}

		if er = keyvaluefactory.AppStore.Echo(); er != nil {
			logger.PrintLog("failed to echo to keyvalue store : ", er.Error())
			os.Exit(99)
		}

		logger.PrintLog("Key value storage ready...")

	}
}

func ReloadObserver() {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGHUP)

	func() {
		for {
			<-sign
			logger.PrintInfo("Received SIGHUP, reloading configuration...")
			LoadConfiguration(true)
		}
	}()
}

func StartLoggerRotation() {
	at := env.GetString("app.log_rotation_at")
	if at != "" {
		gocron.Every(1).Day().At(at).Do(func() {
			logger.PrintLog("Stopping logging temporarily for backup..")
			logger.Logger.Reset()
			suffixFile := time.Now().Format(app.DATE_FORMAT_DATEONLY)
			logPath := fmt.Sprintf("%s%c%s", app.LogDir, os.PathSeparator, app.LogFileName)
			backupFile := fmt.Sprintf("%s_%s", logPath, suffixFile)

			srcLog := fmt.Sprintf("%s.log", logPath)
			destLog := fmt.Sprintf("%s.log", backupFile)
			if env.GetBool("app.log", false) {
				if er := common.CopyAndDelete(srcLog, destLog); er != nil {
					logger.PrintWarn(er)
				}
			}
			erSrcLog := fmt.Sprintf("%s.err", logPath)
			erDestLog := fmt.Sprintf("%s.err", backupFile)
			if er := common.CopyAndDelete(erSrcLog, erDestLog); er != nil {
				logger.PrintWarn(er)
			}

			logger.PrintInfof("The log has been truncated. Check the backup files: %s, %s ", destLog, erDestLog)
		})
		<-gocron.Start()
	}
}
