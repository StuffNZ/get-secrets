package errors

import (
	"os"

	multierror "github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
)

//TODO PanicOnErrors
func PanicOnErrors(err error) {
	if merr, ok := err.(*multierror.Error); ok {
		for _, anErr := range merr.Errors {
			log.Panic(anErr)
		}
	} else {
		log.Panic(err)
	}
}

// Recovery is so that when any other part of the app panics, we can give them a "friendlier" face
func Recovery() {
	if recoveryErr := recover(); recoveryErr != nil {
		// TODO: Re-deliver the stack-trace when debugging
		// TODO: Change to using "github.com/pkg/errors" to capture stack-traces (instead of panic()!)
		log.WithFields(log.Fields{"Err": recoveryErr}).Debug("Panic captured")

		const unexpectedError = 1

		os.Exit(unexpectedError)
	}
}
