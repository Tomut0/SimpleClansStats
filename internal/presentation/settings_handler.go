package presentation

import (
	"bytes"
	"context"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
	_ "time/tzdata"
	"unicode/utf8"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/infrastructure/persistence/sqlite/generated"
)

const maxLogoBytes = 1 << 20
const MaxSettingsRequestBytes = maxLogoBytes + 4096

func (s *ServerHandler) GetAppSettings(ctx context.Context, _ GetAppSettingsRequestObject) (GetAppSettingsResponseObject, error) {
	settings, err := s.settings.GetAppSettings(ctx)
	if err != nil {
		log.Printf("read app settings: %v", err)
		return GetAppSettings500Response{}, nil
	}
	return GetAppSettings200JSONResponse(appSettingsResponse(settings)), nil
}

func (s *ServerHandler) UpdateAppSettings(ctx context.Context, request UpdateAppSettingsRequestObject) (UpdateAppSettingsResponseObject, error) {
	userID, ok := s.sessions.Get(ctx, "user_id").(int64)
	if !ok {
		return UpdateAppSettings401Response{}, nil
	}
	user, err := s.authService.GetByID(userID)
	if err != nil {
		return UpdateAppSettings401Response{}, nil
	}
	if user.Role() != domain.RoleAdmin {
		return UpdateAppSettings403Response{}, nil
	}

	form, err := request.Body.ReadForm(MaxSettingsRequestBytes)
	if err != nil {
		var limit *http.MaxBytesError
		if errors.As(err, &limit) || errors.Is(err, multipart.ErrMessageTooLarge) {
			return UpdateAppSettings413Response{}, nil
		}
		return UpdateAppSettings400Response{}, nil
	}
	defer form.RemoveAll()

	params, ok := parseAppSettings(form)
	if !ok {
		return UpdateAppSettings400Response{}, nil
	}
	logos := form.File["logo"]
	if len(logos) > 1 || len(form.Value["logo"]) != 0 {
		return UpdateAppSettings400Response{}, nil
	}
	if len(logos) == 1 {
		if logos[0].Size > maxLogoBytes {
			return UpdateAppSettings413Response{}, nil
		}
		params.Logo, err = readLogo(logos[0])
		if err != nil {
			return UpdateAppSettings400Response{}, nil
		}
	}

	saved, err := s.settings.UpdateAppSettings(ctx, params)
	if err != nil {
		log.Printf("save app settings: %v", err)
		return UpdateAppSettings500Response{}, nil
	}
	return UpdateAppSettings200JSONResponse(appSettingsResponse(generated.GetAppSettingsRow(saved))), nil
}

func (s *ServerHandler) GetAppLogo(ctx context.Context, _ GetAppLogoRequestObject) (GetAppLogoResponseObject, error) {
	logo, err := s.settings.GetAppLogo(ctx)
	if err != nil {
		log.Printf("read app logo: %v", err)
		return GetAppLogo500Response{}, nil
	}
	switch http.DetectContentType(logo) {
	case "image/png":
		return GetAppLogo200ImagepngResponse{Body: bytes.NewReader(logo), ContentLength: int64(len(logo))}, nil
	case "image/jpeg":
		return GetAppLogo200ImagejpegResponse{Body: bytes.NewReader(logo), ContentLength: int64(len(logo))}, nil
	default:
		return GetAppLogo404Response{}, nil
	}
}

func parseAppSettings(form *multipart.Form) (generated.UpdateAppSettingsParams, bool) {
	var params generated.UpdateAppSettingsParams
	for _, name := range []string{"appName", "removeAuthor", "timeZone", "dateFormat", "timeFormat"} {
		if len(form.Value[name]) != 1 {
			return params, false
		}
	}
	params.AppName = strings.TrimSpace(form.Value["appName"][0])
	params.TimeZone = form.Value["timeZone"][0]
	params.DateFormat = form.Value["dateFormat"][0]
	params.TimeFormat = form.Value["timeFormat"][0]
	removeAuthor := form.Value["removeAuthor"][0]
	params.RemoveAuthor = removeAuthor == "true"

	if params.AppName == "" || !utf8.ValidString(params.AppName) || utf8.RuneCountInString(params.AppName) > 100 ||
		(removeAuthor != "true" && removeAuthor != "false") ||
		!AppSettingsDateFormat(params.DateFormat).Valid() || !AppSettingsTimeFormat(params.TimeFormat).Valid() ||
		params.TimeZone == "" || params.TimeZone == "Local" {
		return params, false
	}
	_, err := time.LoadLocation(params.TimeZone)
	return params, err == nil
}

func readLogo(header *multipart.FileHeader) ([]byte, error) {
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	logo, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	config, format, err := image.DecodeConfig(bytes.NewReader(logo))
	if err != nil || (format != "png" && format != "jpeg") ||
		config.Width < 1 || config.Width > 2048 || config.Height < 1 || config.Height > 2048 {
		return nil, errors.New("invalid logo format or dimensions")
	}
	if _, _, err := image.Decode(bytes.NewReader(logo)); err != nil {
		return nil, err
	}
	return logo, nil
}

func appSettingsResponse(row generated.GetAppSettingsRow) AppSettings {
	return AppSettings{
		AppName:      row.AppName,
		RemoveAuthor: row.RemoveAuthor,
		TimeZone:     row.TimeZone,
		DateFormat:   AppSettingsDateFormat(row.DateFormat),
		TimeFormat:   AppSettingsTimeFormat(row.TimeFormat),
		LogoVersion:  row.LogoVersion,
	}
}
