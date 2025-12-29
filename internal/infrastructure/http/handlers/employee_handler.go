package handlers

import (
	"backend/internal/application/dtos"
	"backend/internal/application/usecases"
	"backend/internal/infrastructure/http/middleware"
	"backend/internal/shared/custom_errors"
	"backend/internal/shared/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type EmployeeHandler struct {
	employeeUC usecases.EmployeeUsecase
}

func NewEmployeeHandler(employeeUC usecases.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUC: employeeUC,
	}
}

// PUT /employee/profile-picture/{uid}
// Request body - image file
// Response body - none
func (h *EmployeeHandler) UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	const MAX_UPLOAD_SIZE = 10 * 1024 * 1024 * 8

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		utils.RespondWithError(w, r, custom_errors.BadRequest(fmt.Errorf("The uploaded file is too big. Please choose an image that is less than 10MB in size.")))
		return
	}

	file, handler, err := r.FormFile("profilePicture")
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("Error retrieving the file")))
		return
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("Error reading file for type detection")))
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		utils.RespondWithError(w, r, custom_errors.BadRequest(fmt.Errorf("The provided file format is not allowed. Please upload a JPEG or PNG image")))
		return
	}

	uid := r.PathValue("uid")
	profilePictureFileName := uid + filepath.Ext(handler.Filename)

	executablePath, err := os.Executable()
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("Error locating executable path: %w", err)))
		return
	}
	executableDir := filepath.Dir(executablePath)
	profilePictureDir := filepath.Join(executableDir, "/storage/employee/profile_picture/")

	entries, err := os.ReadDir(profilePictureDir)
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("failed to read directory %s: %w", profilePictureDir, err)))
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		currentFilename := entry.Name()
		extension := filepath.Ext(currentFilename)
		currentBaseFilename := strings.TrimSuffix(currentFilename, extension)
		if currentBaseFilename == uid {
			fullPath := filepath.Join(profilePictureDir, currentFilename)
			if err := os.Remove(fullPath); err != nil {
				utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("failed to delete file %s: %w", fullPath, err)))
				return
			}
			break
		}
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("Error resetting file read offset")))
		return
	}

	dst, err := os.Create(filepath.Join(profilePictureDir, profilePictureFileName))
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("%w", err)))
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("%w", err)))
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, nil)
}

// DELETE /employee/{id}
// Request body - None
// Response body - None
func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		err = custom_errors.BadRequest(fmt.Errorf("invalid request path parameter to delete employee by ID: %w", err))
		utils.RespondWithError(w, r, err)
		return
	}

	if err := h.employeeUC.Delete(r.Context(), int64(id)); err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /employee/{uiq}
// Request body - none
// Response body - dtos.EmployeeResponse
func (h *EmployeeHandler) GetByUID(w http.ResponseWriter, r *http.Request) {
	resp, err := h.employeeUC.GetByUniqueID(r.Context(), r.PathValue("uid"))
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, resp)
}

// GET /employee/profile-picture/{uid}
// Request body - none
// Response body - none
func (h *EmployeeHandler) GetProfilePicture(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("uid")

	executablePath, err := os.Executable()
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("Error locating executable path: %w", err)))
		return
	}
	executableDir := filepath.Dir(executablePath)
	profilePictureDir := filepath.Join(executableDir, "/storage/employee/profile_picture/")

	var filePath string
	entries, err := os.ReadDir(profilePictureDir)
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("failed to read directory %s: %w", profilePictureDir, err)))
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		currentFilename := entry.Name()
		extension := filepath.Ext(currentFilename)
		currentBaseFilename := strings.TrimSuffix(currentFilename, extension)
		if currentBaseFilename == uid {
			filePath = filepath.Join(profilePictureDir, currentFilename)
			break
		}
	}

	if filePath == "" {
		utils.RespondWithError(w, r, custom_errors.NotFound(fmt.Errorf("file not found")))
		return
	}

	http.ServeFile(w, r, filePath)
}

// GET /employee/personnel
// Request body - none
// Request param - dtos.PersonnelPaginatedQueryParameters
// Response body - none
func (h *EmployeeHandler) GetPersonnelPaginated(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filter := &dtos.PersonnelPaginatedQueryParameters{
		UID:          query.Get("uid"),
		Name:         query.Get("name"),
		Surname:      query.Get("surname"),
		Middlename:   query.Get("middlename"),
		Workplace:    query.Get("workplace"),
		LanguageCode: middleware.GetLanguageFromContext(r.Context()),
	}

	page, err := strconv.ParseInt(query.Get("page"), 0, 64)
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("invalid page parameter provided: %w", err)))
		return
	}
	filter.Page = page

	limit, err := strconv.ParseInt(query.Get("limit"), 0, 64)
	if err != nil {
		utils.RespondWithError(w, r, custom_errors.InternalServerError(fmt.Errorf("invalid page parameter provided: %w", err)))
		return
	}
	filter.Limit = limit

	personnel, err := h.employeeUC.GetPersonnelPaginated(r.Context(), filter)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, personnel)
}

func (h *EmployeeHandler) GetPersonnelCountPaginated(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filter := &dtos.PersonnelPaginatedQueryParameters{
		UID:          query.Get("uid"),
		Name:         query.Get("name"),
		Surname:      query.Get("surname"),
		Middlename:   query.Get("middlename"),
		Workplace:    query.Get("workplace"),
		LanguageCode: middleware.GetLanguageFromContext(r.Context()),
	}

	fmt.Println("COUNT HANDLER FILTER DATA: ", filter)

	personnel, err := h.employeeUC.GetPersonnelCountPaginated(r.Context(), filter)
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	utils.RespondWithJSON(w, r, http.StatusOK, personnel)
}
