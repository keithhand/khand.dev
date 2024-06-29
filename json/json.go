package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type logger interface {
	Warn(string, ...any)
	Error(string, ...any)
}

type jsonHandler[T any] struct {
	logger logger
}

func New(lgr logger) jsonHandler[any] {
	return jsonHandler[any]{
		logger: lgr,
	}
}

func (j jsonHandler[T]) UnmarshallUrl(url string, obj T) T {
	resp, err := http.Get(url)
	if err != nil {
		j.logger.Error(fmt.Errorf("initializing request to %s: %w", url, err).Error())
		return obj
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.logger.Error(fmt.Errorf("reading %T reponse: %w", obj, err).Error())
		return obj
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		j.logger.Error(fmt.Errorf("unmarshalling %T from json: %w", obj, err).Error())
		return obj
	}

	return obj
}
