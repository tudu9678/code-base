package server

// var whiteList = []*cm.ApiUrl{
// 	{
// 		UrlStr: "/backoffice/v1/auth/google/login",
// 		Method: echo.POST,
// 	},
// }

// func (s Server) MapHandlers(e *echo.Echo) error {
// 	ce, err := rbacutil.New(s.db)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// versioning
// 	clientRouter := e.Group("/client")

// 	backofficeRouter := e.Group("/backoffice")
// 	am := cm.NewApiMiddleware(s.conf, s.cache, s.logger)
// 	amBO := cm.NewApiMiddlewareBO(s.conf)

// 	logService := activityLog.NewActivityLogService(s.db, s.conf)
// 	systemLogService := service.NewSystemLogService(s.db, s.conf)

// 	clientRouter.Use(cm.ActivityMiddleware(s.logger, s.db, logService, am))

// 	backofficeRouter.Use(cm.RBACMiddleware(ce, s.cache, s.conf, whiteList))
// 	// User
// 	userHttp.MapRoutesClient(e, clientRouter, s.db, s.conf, am, systemLogService)
// 	userHttp.MapRoutesBackoffice(e, backofficeRouter, s.db, s.logger, s.conf, s.cache, amBO, systemLogService)

// 	// NDID
// 	ndidHttp.MapRoutesClient(e, clientRouter, s.db, s.cache, s.conf, am, systemLogService, s.logger)
// 	appmanHttp.MapRoutesClient(e, clientRouter, s.db, s.conf, am, systemLogService, s.logger)

// 	// Suitability-test
// 	suitabilityhttp.MapRoutesClient(e, clientRouter, s.db, s.conf, am)

// 	//E-Kyc
// 	eKycHttp.MapRoutesClient(e, clientRouter, s.db, s.cache, s.logger, s.conf, am, systemLogService)
// 	eKycHttp.MapRoutesBO(e, backofficeRouter, s.db, s.cache, s.logger, s.conf, amBO, systemLogService)

// 	//BO
// 	userAdminHttp.MapRoutesClient(e, backofficeRouter, s.db, s.conf, s.cache, ce)

// 	// Fund Model
// 	fundModelHttp.MapRoutesClient(e, clientRouter, s.db, s.cache, s.conf, am)
// 	fundModelHttp.MapRoutesBackOffice(e, backofficeRouter, s.db, s.cache, s.conf, amBO)

// 	// Order
// 	orderModule := order.NewModule(s.logger, s.conf)
// 	// client
// 	orderModule.RegisterClientRoute(clientRouter.Group("/v1/orders"), s.db, s.cache, s.conf, am, systemLogService)
// 	// BO
// 	orderModule.RegisterBORoute(backofficeRouter.Group("/v1/orders"), s.db, s.cache, s.conf, systemLogService)
// 	// callback
// 	kkpGroup := e.Group("/kkp/payment/callback")
// 	kkpGroup.Use(cm.InboundMiddleware(s.logger, s.db, systemLogService))
// 	orderModule.RegisterCallbackRoute(kkpGroup, s.db, s.cache, s.conf, systemLogService)

// 	// Bank
// 	banks.NewModule(s.db, s.logger, s.conf, systemLogService).RegisterRoute(clientRouter, am)
// 	banks.NewModuleBO(s.db, s.logger, s.conf).RegisterRoute(backofficeRouter)

// 	// Role
// 	roles.NewModuleBO(s.db, s.logger, s.conf, s.cache).RegisterRoute(backofficeRouter)

// 	// User-role
// 	userRole.MapRoutesClientBO(e, backofficeRouter, s.db, s.cache, s.conf, amBO)

// 	// Env setting
// 	envSetting.MapRoutesClient(e, clientRouter, s.db, s.cache, s.conf, am)
// 	envSetting.MapRoutesBackOffice(e, backofficeRouter, s.db, s.cache, s.conf, amBO)

// 	// Task
// 	task.MapRoutesBackOffice(e, backofficeRouter, s.db, s.cache, s.conf, amBO)

// 	// Investment risk
// 	investmentRisk.MapRoutesClient(e, clientRouter, s.db, s.cache, s.conf, am)
// 	investmentRisk.MapRoutesBO(e, backofficeRouter, s.db, s.cache, s.conf, amBO)

// 	return nil
// }
