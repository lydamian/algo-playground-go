package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	// config
	var PORT string = "6000"

	// logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	muxRouter := mux.NewRouter()

	// middleware
	RegisterRoutes(muxRouter)
	muxRouter.Use(RequestLoggerMiddleware(log.Logger))

	log.Info().Msg(fmt.Sprintf("Starting server... on PORT:%s", PORT))
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), muxRouter)
	log.Fatal().Msg(fmt.Sprintf("Error starting server on PORT:%s", PORT))
}
