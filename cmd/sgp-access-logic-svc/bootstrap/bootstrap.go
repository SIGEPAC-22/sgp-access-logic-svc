package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/dimiro1/health"
	kitlog "github.com/go-kit/log"
	_ "github.com/go-sql-driver/mysql"
	goconfig "github.com/iglin/go-config"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sgp-access-logic-svc/internal/accessAuthLogin/platform/handler"
	"sgp-access-logic-svc/internal/accessAuthLogin/platform/storage/mysql"
	"sgp-access-logic-svc/internal/accessAuthLogin/service"
	handler2 "sgp-access-logic-svc/internal/createPersonalInfo/platform/handler"
	mysql2 "sgp-access-logic-svc/internal/createPersonalInfo/platform/storage/mysql"
	service2 "sgp-access-logic-svc/internal/createPersonalInfo/service"
	handler4 "sgp-access-logic-svc/internal/deletePersonalInfo/platform/handler"
	mysql4 "sgp-access-logic-svc/internal/deletePersonalInfo/platform/storage/mysql"
	service4 "sgp-access-logic-svc/internal/deletePersonalInfo/service"
	handler3 "sgp-access-logic-svc/internal/updatePersonalInfo/platform/handler"
	mysql3 "sgp-access-logic-svc/internal/updatePersonalInfo/platform/storage/mysql"
	service3 "sgp-access-logic-svc/internal/updatePersonalInfo/service"
	"syscall"
)

func Run() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	port := config.GetString("server.port")

	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	mux := http.NewServeMux()
	errs := make(chan error, 2)
	////////////////////////////////////////////////////////////////////////
	////////////////////////CORS///////////////////////////////////////////
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handlerCORS := cors.Handler(mux)
	////////////////////////CORS///////////////////////////////////////////

	db, err := sql.Open("mysql", getStrConnection())
	if err != nil {
		log.Fatalf("unable to open database connection %s", err.Error())
	}

	/////////////////////GET DATA AUTH LOGIN/////////////////////
	repoGetDataAuthLogin := mysql.NewGetDataAuthLoginRepository(db, kitlogger)
	serviceGetDataAuthLogin := service.NewGetDataAuthLoginSvc(repoGetDataAuthLogin, kitlogger)
	endpointGetDataAuthLogin := handler.MakeGetDataAuthLoginEndpoint(serviceGetDataAuthLogin)
	endpointGetDataAuthLogin = handler.GetDataAuthLoginTransportMiddleware(kitlogger)(endpointGetDataAuthLogin)
	transportGetDataAuthLogin := handler.NewGetDataAuthLoginHandler(config.GetString("paths.getDataAuthLogin"), endpointGetDataAuthLogin)
	/////////////////////GET DATA AUTH LOGIN/////////////////////

	/////////////////////CREATE PERSONAL/////////////////////
	repoCreatePersonalInfo := mysql2.NewCreatePersonalInfoRepository(db, kitlogger)
	serviceCreatePersonalInfo := service2.NewCreatePersonalInfoSvc(repoCreatePersonalInfo, kitlogger)
	endpointCreatePersonalInfo := handler2.MakeCreatePersonalInfoEndpoint(serviceCreatePersonalInfo)
	endpointCreatePersonalInfo = handler2.CreatePersonalInfoTransportMiddleware(kitlogger)(endpointCreatePersonalInfo)
	transportCreatePersonalInfo := handler2.NewCreatePersonalInfoHandler(config.GetString("paths.createPersonalInfo"), endpointCreatePersonalInfo)
	/////////////////////GET DATA AUTH LOGIN/////////////////////

	/////////////////////UPDATE PERSONAL/////////////////////
	repoUpdatePersonalInfo := mysql3.NewUpdatePersonalInfoRepository(db, kitlogger)
	serviceUpdatePersonalInfo := service3.NewUpdateInfoPatientService(repoUpdatePersonalInfo, kitlogger)
	endpointUpdatePersonalInfo := handler3.MakeUpdatePersonalInfoEndpoint(serviceUpdatePersonalInfo)
	endpointUpdatePersonalInfo = handler3.UpdateInfoPatientTransportMiddleware(kitlogger)(endpointUpdatePersonalInfo)
	transportUpdatePersonalInfo := handler3.NewUpdatePersonalInfoHandler(config.GetString("paths.updatePersonalInfo"), endpointUpdatePersonalInfo)
	/////////////////////UPDATE PERSONAL/////////////////////

	/////////////////////DELETE PERSONAL/////////////////////
	repoDeletePersonalInfo := mysql4.NewDeletePersonalInfoRepository(db, kitlogger)
	serviceDeletePersonalInfo := service4.NewDeletePersonalInfoService(repoDeletePersonalInfo, kitlogger)
	endpointDeletePersonalInfo := handler4.MakeDeletePersonalInfoEndpoint(serviceDeletePersonalInfo)
	endpointDeletePersonalInfo = handler4.DeletePersonalInfoTransportMiddleware(kitlogger)(endpointDeletePersonalInfo)
	transportDeletePersonalInfo := handler4.NewDeletePersonalInfoHandler(config.GetString("paths.deletePersonalInfo"), endpointDeletePersonalInfo)
	/////////////////////DELETE PERSONAL/////////////////////

	mux.Handle(config.GetString("paths.getDataAuthLogin"), transportGetDataAuthLogin)
	mux.Handle(config.GetString("paths.createPersonalInfo"), transportCreatePersonalInfo)
	mux.Handle(config.GetString("paths.updatePersonalInfo"), transportUpdatePersonalInfo)
	mux.Handle(config.GetString("paths.deletePersonalInfo"), transportDeletePersonalInfo)
	mux.Handle("/health", health.NewHandler())

	go func() {
		kitlogger.Log("listening", "transport", "http", "address", port)
		errs <- http.ListenAndServe(":"+port, handlerCORS)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		db.Close()
	}()
	kitlogger.Log("terminated", <-errs)
}

func getStrConnection() string {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	pass := config.GetString("datasource.pass")
	dbname := config.GetString("datasource.dbname")
	strconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, dbname)
	return strconn
}
