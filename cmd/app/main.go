package main

import (
	cc "comick-dl/internal/comickclient"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logLevel := flag.String("level", "info", "debug | info | warn | error")
	flag.Parse()

	if flag.NArg() != 1 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	url := flag.Arg(0)

	var lvl slog.Level
	err := (&lvl).UnmarshalText([]byte(*logLevel))
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid --loglevel %q (using info)\n", *logLevel)
		lvl = slog.LevelInfo
	}

	lvlvar := new(slog.LevelVar)
	lvlvar.Set(lvl)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:     lvlvar,
		AddSource: lvl <= slog.LevelDebug,
	}))

	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(lvl)

	header := http.Header{
		"User-Agent": {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"},
	}
	client := cc.New(
		cc.WithHeaders(header),
		cc.WithLogger(logger),
	)

	client.Download(context.Background(), url, "")

}
