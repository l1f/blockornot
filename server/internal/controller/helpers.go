package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/l1f/blockornot/validator"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type envelope map[string]interface{}

func (c *Controllers) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (c *Controllers) writeJSON(webCtx *WebContext, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, val := range headers {
		webCtx.Response.Header()[key] = val
	}

	webCtx.Response.Header().Set("Content-Type", "application/json")
	webCtx.Response.WriteHeader(status)
	_, _ = webCtx.Response.Write(js)

	return nil
}

func (c *Controllers) readJSON(webCtx *WebContext, dst interface{}) error {
	// Limit the size of the Request body to 1MB
	maxBytes := 1_048_576
	webCtx.Request.Body = http.MaxBytesReader(webCtx.Response, webCtx.Request.Body, int64(maxBytes))

	// field which cannot be mapped to the target destination,
	// the decoder will return an error
	// instead of just ignoring the field
	dec := json.NewDecoder(webCtx.Request.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at charecter %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: Request body to large":
			return fmt.Errorf("body must not be lager than %d bytes", maxBytes)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	// If the Request body not only contained a single JSON value this will
	// a custom error message
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (c *Controllers) readString(qs url.Values, key, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

func (c *Controllers) readCSV(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}

func (c *Controllers) readInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return i
}

func (c *Controllers) background(fn func()) {
	c.ctx.Wg.Add(1)

	go func() {
		defer c.ctx.Wg.Done()
		defer func() {
			if err := recover(); err != nil {
				c.ctx.Logger.Error.Println(err.(string), nil)
			}
		}()

		time.Sleep(5 * time.Second)
		fn()

	}()
}

func (c *Controllers) NotImplementedYet(webCtx *WebContext) {
	c.ServerErrorResponse(webCtx, errors.New("not implemented jet"))
}
