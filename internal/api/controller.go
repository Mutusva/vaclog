package api

import (
	"Codenotary/config"
	"Codenotary/internal/immuDB"
	"Codenotary/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type VacRecordController struct {
	Config       config.AppConfig
	ImmuDBClient immuDB.ImmuDBRepository
}

func NewVacController(config config.AppConfig, immuDBClient immuDB.ImmuDBRepository) VacRecordController {
	return VacRecordController{
		Config:       config,
		ImmuDBClient: immuDBClient,
	}
}

// GetAllRecords Get all vaccination records for animal
// @BasePath /api/v1
// @Summary all vaccination records for farm animals
// @Description all vaccination records for farm animals
// @Tags vac_records
// @Accept json
// @Produce json
// @Success 200 {object} VacRecordResponse
// @Failure 403 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /records/ [get]
func (v *VacRecordController) GetAllRecords(c *gin.Context) {
	records, err := v.ImmuDBClient.GetAllRecords(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": model.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "error getting records",
		}})
		// log to a logging system
		log.Printf(err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, VacRecordResponse{
		Records: records,
	})
}

// CreateRecord creates vaccination record
// @BasePath /api/v1
// @Summary creates a vaccination record for an animal
// @Description creates a vaccination record for an animal
// @Tags vac_records
// @Accept json
// @Produce json
// @Param payload body model.VacRecordModel true "The JSON for a vaclog"
// @Success 200 {object} model.RecordCreateResponse
// @Failure 403 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /records/ [put]
func (v *VacRecordController) CreateRecord(c *gin.Context) {
	var vacRequest model.CreateRecordModel

	// Call BindJSON to bind the received JSON to vacRecord
	if err := c.BindJSON(&vacRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": model.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "error parsing json",
		}})
		log.Printf(err.Error())
		return
	}

	res, err := v.ImmuDBClient.CreateRecord(c, vacRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": model.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}})
		log.Printf(err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// SearchRecord searches for a vaccination record
// @BasePath /api/v1
// @Summary search for a vaccination record
// @Description search for a vaccination record for an animal and returns all the revision
// @Tags vac_records
// @Accept json
// @Produce json
// @Success 200 {object} []model.VacRecordModel
// @Failure 403 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /records/:documentID [GET]
func (v *VacRecordController) SearchRecord(c *gin.Context) {
	docID := c.Param("documentID")
	if docID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": model.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "document ID require",
		}})
		return
	}

	res, err := v.ImmuDBClient.SearchRecord(c, docID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": model.ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}})
		return
	}

	c.JSON(http.StatusOK, res)
}
