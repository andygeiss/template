package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/andygeiss/template/clients"
	"github.com/andygeiss/template/messages"
	"github.com/andygeiss/template/resources"
	"github.com/andygeiss/template/services"
	"github.com/andygeiss/utilities/logging"
	"github.com/andygeiss/utilities/messaging"
)

var (
	name    string = "no-name"
	build   string = "no-build"
	version string = "no-version"
)

func main() {
	log.Printf("%s %s (%s)\n", name, version, build)

	sut := func(w http.ResponseWriter, r *http.Request) {
		// Utilities ...
		logger := logging.NewDefaultLogger()
		bus := messaging.NewDefaultBus(logger)
		// Clients ...
		tradesmanClient := clients.NewTradesman(bus, logger)
		// Resources ...
		memberAccess := resources.NewMemberAccess(logger)
		// Engines ...
		regulationsEngine := services.NewRegulationsEngine(logger)
		// Managers ...
		membershipManager := services.NewMembershipManager(bus, logger, regulationsEngine, memberAccess)
		// Register actors ...
		bus.Subscribe(tradesmanClient)
		bus.Subscribe(membershipManager)
		// Main Logic ...
		logger.Print("Main sends ApplyForMembership")
		bus.Publish(messages.ApplyForMembership{})
	}

	http.HandleFunc("/", sut)
	http.ListenAndServe(":3000", nil)
}
