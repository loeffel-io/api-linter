package lint

import (
	"errors"
	"strings"
)

// Run invokes all rules on the request.
func Run(rules Rules, request Request) (Response, error) {
	return run(rules.All(), request)
}

func run(rules []Rule, req Request) (Response, error) {
	finalResp := Response{}
	errMessages := []string{}
	for _, r := range rules {
		if resp, err := r.Lint(req); err == nil {
			finalResp.merge(resp)
		} else {
			errMessages = append(errMessages, err.Error())
		}
	}

	if len(errMessages) != 0 {
		err := errors.New(strings.Join(errMessages, "; "))
		return finalResp, err
	}

	return finalResp, nil
}
