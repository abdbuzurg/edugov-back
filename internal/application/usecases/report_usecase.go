package usecases

import (
	"backend/internal/infrastructure/http/middleware"
	"backend/internal/infrastructure/persistence/postgres"
	"backend/internal/shared/custom_errors"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"
)

type ReportUsecase interface {
	GenerateSummaryDataReport(ctx context.Context) (*string, *string, error)
}

type reportUsecase struct {
	store     *postgres.Store
	validator *validator.Validate
}

func NewReportUsecase(
	store *postgres.Store,
	validator *validator.Validate,
) ReportUsecase {
	return &reportUsecase{
		store:     store,
		validator: validator,
	}
}

func (uc *reportUsecase) GenerateSummaryDataReport(ctx context.Context) (*string, *string, error) {
	summaryDataReportQueryResult, err := uc.store.GetSummaryData(ctx, middleware.GetLanguageFromContext(ctx))
	if err != nil {
		if custom_errors.IsNotFound(err) {
			return nil, nil, nil
		}

		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("failed to retrive summary data: %w", err))
	}

	type ExcelData struct {
		BachelorCount        int64
		SpecialistCount      int64
		MastersCount         int64
		IntershipCount       int64
		ResidencyCount       int64
		CandidateOfScience   int64
		PhdCount             int64
		DoctorOfScienceCount int64
		Total                int64
	}
	excelData := map[string]ExcelData{}
	degreeKeys := []string{
		"Бакалавр",
		"Магистр",
		"Мутахассис",
		"Номзади илм",
		"Интернатура",
		"Ординатура",
		"PhD (Доктори фалсафа)",
		"Доктори илм",
	}

	totalPerDegree := make(map[string]*int64)

	for _, key := range degreeKeys {
		var count int64 = 0
		totalPerDegree[key] = &count
	}

	for _, entry := range summaryDataReportQueryResult {
		if _, ok := excelData[entry.Workplace]; !ok {
			excelData[entry.Workplace] = ExcelData{}
		}

		tempData := excelData[entry.Workplace]

		switch entry.DegreeLevel {
		case "Бакалавр":
			tempData.BachelorCount++

		case "Магистр":
			tempData.MastersCount++

		case "Мутахассис":
			tempData.SpecialistCount++

		case "Номзади илм":
			tempData.CandidateOfScience++

		case "Интернатура":
			tempData.IntershipCount++

		case "Ординатура":
			tempData.ResidencyCount++

		case "PhD (Доктори фалсафа)":
			tempData.PhdCount++

		case "Доктори илм":
			tempData.DoctorOfScienceCount++
		}

		tempData.Total++
		*totalPerDegree[entry.DegreeLevel]++
		excelData[entry.Workplace] = tempData
	}

	executablePath, err := os.Executable()
	if err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error getting the executable file path: %w", err))
	}
	executableDir := filepath.Dir(executablePath)
	excelTemplateFilepath := filepath.Join(executableDir, "/internal/files/report_templates/summary_data.xlsx")

	f, err := excelize.OpenFile(excelTemplateFilepath)
	if err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error openning excel file: %w", err))
	}

	sheetName := "Лист1"
	startingRow := 2
	f.InsertRows(sheetName, 2, len(excelData))

	institutionNamesCellStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
		Fill: excelize.Fill{
			Type:  "pattern",
			Color: []string{"#DAE9F7"},
		},
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	if err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error setting up cell style for institutionNames: %w", err))
	}

	numbersCellStyle, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error setting up cell style for numbers: %w", err))
	}

	index := 0
	for institutionName, values := range excelData {
		if err := f.SetCellStyle(sheetName, "A"+fmt.Sprint(startingRow+index), "A"+fmt.Sprint(startingRow+index), institutionNamesCellStyle); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error setting up the style in the excel for institutionNames: %w", err))
		}

		if err := f.SetCellStyle(sheetName, "B"+fmt.Sprint(startingRow+index), "J"+fmt.Sprint(startingRow+index), numbersCellStyle); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error setting up the style in the excel for degree cells: %w", err))
		}

		if err := f.SetCellStr(sheetName, "A"+fmt.Sprint(startingRow+index), institutionName); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the A%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "B"+fmt.Sprint(startingRow+index), values.BachelorCount); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the B%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "C"+fmt.Sprint(startingRow+index), values.SpecialistCount); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the C%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "D"+fmt.Sprint(startingRow+index), values.MastersCount); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the D%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "E"+fmt.Sprint(startingRow+index), values.IntershipCount); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the E%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "F"+fmt.Sprint(startingRow+index), values.ResidencyCount); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the F%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "G"+fmt.Sprint(startingRow+index), values.CandidateOfScience); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the G%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "H"+fmt.Sprint(startingRow+index), values.PhdCount); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the H%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "I"+fmt.Sprint(startingRow+index), values.DoctorOfScienceCount); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the I%d: %w", startingRow+index, err))
		}

		if err := f.SetCellInt(sheetName, "J"+fmt.Sprint(startingRow+index), values.Total); err != nil {
			return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the J%d: %w", startingRow+index, err))
		}

		index++
	}

	overallTotal := 0
	for _, values := range totalPerDegree {
		overallTotal += int(*values)
	}

	if err := f.SetCellInt(sheetName, "B"+fmt.Sprint(startingRow+index), *totalPerDegree["Бакалавр"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the B%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "C"+fmt.Sprint(startingRow+index), *totalPerDegree["Магистр"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the C%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "D"+fmt.Sprint(startingRow+index), *totalPerDegree["Мутахассис"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the D%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "E"+fmt.Sprint(startingRow+index), *totalPerDegree["Номзади илм"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the E%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "F"+fmt.Sprint(startingRow+index), *totalPerDegree["Интернатура"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the F%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "G"+fmt.Sprint(startingRow+index), *totalPerDegree["Ординатура"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the G%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "H"+fmt.Sprint(startingRow+index), *totalPerDegree["PhD (Доктори фалсафа)"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the H%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "I"+fmt.Sprint(startingRow+index), *totalPerDegree["Доктори илм"]); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the I%d: %w", startingRow+index, err))
	}

	if err := f.SetCellInt(sheetName, "J"+fmt.Sprint(startingRow+index), int64(overallTotal)); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error writing on the J%d: %w", startingRow+index, err))
	}

	currentTime := time.Now()
	reportFileName := fmt.Sprintf(
		"Summary Report - %s.xlsx",
		currentTime.Format("02-01-2006"),
	)

	reportFilePath := filepath.Join(executableDir, "/internal/files/tmp/", reportFileName)
	if err := f.SaveAs(reportFilePath); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error savinging file: %w", err))
	}

	if err := f.Close(); err != nil {
		return nil, nil, custom_errors.InternalServerError(fmt.Errorf("Error closing file: %w", err))
	}

	return &reportFilePath, &reportFileName, nil
}
