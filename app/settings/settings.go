package settings

import (
	"os"

	"github.com/chocomilk/app/model"
	general "github.com/chocomilk/app/pkg/general/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Service interface {
	InitializeRouter(app fiber.Router)
	InitializeValidator()
	MakeMigrations()
}

var Services = map[string]Service{
	"general": general.ChocomilkGeneral{},
}

var FiberConfig = fiber.Config{
	AppName:   "Chocomilk",
	BodyLimit: 50 * 1024 * 1024, // 50 MB

	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		/*
			Generates an error log
		*/

		// headers := ctx.GetReqHeaders()
		// logString := fmt.Sprintf(
		// 	`ERROR, Time: %s, URL: %s, Method: %s, Headers: %s, Body: %s, Error: %s`,
		// 	time.Now().Format("2006-01-02 15:04:05"), ctx.BaseURL()+ctx.OriginalURL(), ctx.Method(), headers, string(ctx.Body()), err.Error(),
		// )
		// fmt.Println(logString)

		// curlCommand := GetCurlCommand(ctx)
		// fmt.Println("Curl: ", curlCommand)

		return model.ErrorResponse(ctx, code, err.Error())
	},
}

/*
	Logger middleware config
*/
var LoggerConfig = logger.Config{
	Format: "Ip and Port: ${ip}:${port}, " +
		"Date: [${time}], " +
		"Status: ${status}, " +
		"Method: ${method}, " +
		"Url: ${protocol}://${host}${url}, " +
		"Path: ${path}, " +
		"Query Params: ${queryParams}, " +
		"Request Body: ${body}, " +
		"Request Header: ${reqHeaders}, " +
		"Response Body: ${resBody}\n",
	TimeFormat: "01/02/2006 15:04:05",
	TimeZone:   "Local",
	Output:     os.Stdout,
}

func SetupApp() *fiber.App {
	app := fiber.New(FiberConfig)
	app.Use(recover.New())

	app.Use(logger.New(LoggerConfig))
	return app
}
