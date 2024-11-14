package logger

import (
	"fmt"
	"os"
	"strings"

	"log"

	"github.com/otie173/odncore/internal/utils/webhook/discord"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
)

func Register() {
	infoLogger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stdout, "WARN\t", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	fatalLogger = log.New(os.Stdout, "FATAl\t", log.Ldate|log.Ltime)
}

func Info(text ...string) {
	infoLogger.Print(strings.Join(text, " ") + "\n")
}

func Infof(format string, args ...any) {
	infoLogger.Printf(format, args...)
}

func Warn(text ...any) {
	warnLogger.Print(fmt.Sprintln(text...))
}

func Warnf(format string, args ...any) {
	warnLogger.Printf(format, args...)
}

func Error(text ...any) {
	errorLogger.Print(fmt.Sprintln(text...))
}

func Errorf(format string, args ...any) {
	errorLogger.Printf(format, args...)
}

func Fatal(text ...any) {
	fatalLogger.Fatal(fmt.Sprint(text...))
}

func Fatalf(format string, args ...any) {
	fatalLogger.Fatalf(format, args...)
}

func Server(text string) {
	Info(text)

	if discord.WebhookEnabled() {
		if err := discord.SendMessage(text); err != nil {
			Error("Error with send discord webhook message: ", err)
		}
	}
}

func Player(name, text string) {
	Infof("%s %s\n", name, text)

	if discord.WebhookEnabled() {
		if err := discord.PlayerMessage(fmt.Sprintf("**%s**", name), text); err != nil {
			Error("Error with send discord webhook message: ", err)
		}
	}
}
