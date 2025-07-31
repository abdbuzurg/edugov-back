package utils

import (
	"backend/internal/infrastructure/http/middleware"
	"backend/internal/shared/custom_errors"
	"encoding/json"
	"log"
	"net/http"
)

// Sends JSON response to the client.
// Code is HTTP STATUS code
func RespondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload any) {
	requestID := middleware.GetRequestIDFromContext(r.Context())

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[%s] Error marshalling response JSON response: %v", requestID, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("[%s] Error writting JSON response: %v", requestID, err)
	}
}

var errorMessages = map[string]map[int]string{
	"en": {
		http.StatusInternalServerError: "An unexpected error occurred. Please try again later.",
		http.StatusMethodNotAllowed:    "Method Not Allowed.",
		http.StatusBadRequest:          "Bad Request. Please check your input.",
		http.StatusNotFound:            "Resource not found.",
		http.StatusUnauthorized:        "You are not authorized or logged in. Please relog in",
		520:                            "Unknown error. Please try again later",
	},
	"ru": {
		http.StatusInternalServerError: "Произошла непредвиденная ошибка. Пожалуйста, попробуйте позже.",
		http.StatusMethodNotAllowed:    "Метод не разрешен.",
		http.StatusBadRequest:          "Неверный запрос. Пожалуйста, проверьте ввод.",
		http.StatusNotFound:            "Ресурс не найден.",
		http.StatusUnauthorized:        "Вы не авторизованы или не вошли в систему. Пожалуйста, войдите снова.",
		520:                            "Неопознанная ошибка. Пожалуйста, попробуйте позже.",
	},
	"tg": {
		http.StatusInternalServerError: "Хатои ғайричашмдошт рух дод. Лутфан баъдтар кӯшиш кунед.",
		http.StatusMethodNotAllowed:    "Усули иҷозатнашуда.",
		http.StatusBadRequest:          "Дархости нодуруст. Лутфан, маълумоти худро тафтиш кунед.",
		http.StatusNotFound:            "Манбаъ ёфт нашуд.",
		http.StatusUnauthorized:        "Шумо ваколатдор нестед ё ворид нашудаед. Лутфан, дубора ворид шавед.",
		520:                            "Хатогии номаълум. Лутфан баъдтар кӯшиш кунед.",
	},
}

func RespondWithError(w http.ResponseWriter, r *http.Request, err error) {
	lang := middleware.GetLanguageFromContext(r.Context())
	requestID := middleware.GetRequestIDFromContext(r.Context())

	if appErr, ok := err.(*custom_errors.AppError); ok {
		log.Print(appErr.Log(requestID))
		if appErr.StatusCode == http.StatusBadRequest {
			RespondWithJSON(w, r, appErr.StatusCode, map[string]string{"message": appErr.Err.Error()})
			return
		}
		RespondWithJSON(w, r, appErr.StatusCode, map[string]string{"message": errorMessages[lang][appErr.StatusCode]})
		return
	}

	log.Printf("[%s] Unidenfied error: %v", requestID, err)
	RespondWithJSON(w, r, 520, map[string]string{"message": errorMessages[lang][520]})
}
