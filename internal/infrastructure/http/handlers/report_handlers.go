package handlers

import (
	"backend/internal/application/usecases"
	"backend/internal/shared/utils"
	"log"
	"net/http"
	"os"
)

type ReportHandler struct {
	reportUC usecases.ReportUsecase
}

func NewReportHandler(reportUC usecases.ReportUsecase) *ReportHandler {
	return &ReportHandler{
		reportUC: reportUC,
	}
}

// GET /report/summary-data
// Request body - none
// Response body - excel file
func (h *ReportHandler) SummaryReport(w http.ResponseWriter, r *http.Request) {
	summaryReportFilePath, summaryReportFileName, err := h.reportUC.GenerateSummaryDataReport(r.Context())
	if err != nil {
		utils.RespondWithError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename="+*summaryReportFileName)
	http.ServeFile(w, r, *summaryReportFilePath)

	defer func() {
		err := os.Remove(*summaryReportFilePath)
		if err != nil {
			log.Printf("ERROR: Failed to delete temporary file %s: %v", *summaryReportFilePath, err)
		} else {
			log.Printf("Successfully deleted temporary file: %s", *summaryReportFilePath)
		}
	}()
}
