package infermedica

import "os"

func GetTestAppInstance() App {
	return NewApp(os.Getenv("INFERMEDICA_APP_ID"), os.Getenv("INFERMEDICA_APP_KEY"), "infermedica-en", "")
}
