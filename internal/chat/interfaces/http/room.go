package http

import (
	"fmt"

	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// CreateRoom is the handler for the create room path
func (h *handler) CreateRoom() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request dto.CreateRoomReqDTO
		err := f.BodyParser(&request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.RequestBodyParseErrorMessage,
				zap.Error(err),
			)

			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.RequestBodyParseErrorCode, Message: errorhandler.RequestBodyParseErrorMessage, StatusCode: fiber.StatusBadRequest})
		}

		request.Owner = f.Locals("userID").(string)

		validate := validator.New()

		// Validate the request
		err = validate.Struct(request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.ValidationErrorMessage,
				zap.Error(err),
			)

			fields := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.ValidationErrorCode, Message: fmt.Sprintf("invalid fields %s", fields), StatusCode: fiber.StatusBadRequest})
		}

		// Handle the request
		res, resErr := h.application.RoomCommandHandler.CreateRoom(&request)

		if resErr != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.InvalidCredentialsMessage,
				zap.Error(err),
			)

			return f.Status(resErr.StatusCode).JSON(resErr)
		}

		return f.Status(fiber.StatusOK).JSON(res)
	}
}

// QueryRooms is the handler for the query rooms path
func (h *handler) QueryRooms() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request = &dto.QueryRoomsReqDTO{
			Page:      f.QueryInt("page", 1),
			PerPage:   f.QueryInt("per_page", 10),
			SortField: f.Query("sort_field", "created_at"),
			SortOrder: f.Query("sort_order", "desc"),
		}

		// Filters
		if f.Query("user_in") != "" {
			request.UserIn = []string{f.Query("user_in")}

			// Add the current user to the filter
			request.UserIn = append(request.UserIn, f.Locals("userID").(string))
		}

		if f.Query("user_not_in") != "" {
			request.UserNotIn = []string{f.Query("user_not_in")}
		}

		validate := validator.New()

		// Validate the request
		err := validate.Struct(request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.ValidationErrorMessage,
				zap.Error(err),
			)

			fields := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.ValidationErrorCode, Message: fmt.Sprintf("invalid fields %s", fields), StatusCode: fiber.StatusBadRequest})
		}

		// Handle the request
		res, resErr := h.application.RoomQueryHandler.QueryRooms(request)

		if resErr != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.InvalidCredentialsMessage,
				zap.Error(err),
			)

			return f.Status(resErr.StatusCode).JSON(resErr)
		}

		return f.Status(fiber.StatusOK).JSON(res)
	}
}
