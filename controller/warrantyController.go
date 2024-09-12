package controller

import (
	"employeeregister/database"
	"employeeregister/models"
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

// // Initialize logger
// var logger = logrus.New()

func WarrntyUpload(c *gin.Context) {
	logger.Info("Received request for WarrntyUpload")

	var warranty models.Warranty
	rand.Seed(time.Now().UnixNano())
	warranty.WarrantyId = rand.Intn(2000000)
	warranty.Status = "pending"

	if err := c.BindJSON(&warranty); err != nil {
		logger.WithError(err).Error("Error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&warranty)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Error creating warranty")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	logger.WithFields(logrus.Fields{
		"warranty": warranty,
	}).Info("Warranty created successfully")
	c.JSON(http.StatusOK, &warranty)
}

func WaarntyGet(c *gin.Context) {
	logger.Info("Received request for WaarntyGet")

	var Warranty []models.Warranty

	result := database.DB.Find(&Warranty)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Error retrieving warranties")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	logger.WithFields(logrus.Fields{
		"warranties": Warranty,
	}).Info("Warranties retrieved successfully")
	c.JSON(http.StatusOK, &Warranty)
}

func WarrantyUpdate(c *gin.Context) {
	logger.Info("Received request for WarrantyUpdate")

	id := c.Param("id")
	body := models.Warranty{}
	var Warranty models.Warranty
	result := database.DB.First(&Warranty, id)

	if err := c.BindJSON(&body); err != nil {
		logger.WithError(err).Error("Error binding JSON")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result.Error != nil {
		logger.WithError(result.Error).Error("Error finding warranty")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	Warranty.AnnualPrice = body.AnnualPrice
	Warranty.ProductName = body.ProductName
	Warranty.Vendor = body.Vendor
	Warranty.Status = body.Status
	database.DB.Save(&Warranty)

	logger.WithFields(logrus.Fields{
		"warranty": Warranty,
	}).Info("Warranty updated successfully")
	c.JSON(http.StatusOK, &Warranty)
}

func DeleteWarranty(c *gin.Context) {
	logger.Info("Received request for DeleteWarranty")

	var Warranty models.Warranty
	id := c.Param("id")
	result := database.DB.First(&Warranty, id)

	if result.Error != nil {
		logger.WithError(result.Error).Error("Error finding warranty")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	database.DB.Delete(&Warranty, id)

	logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Warranty deleted successfully")
	c.JSON(http.StatusOK, "Delete Successful")
}

func PendingWarranty(c *gin.Context) {
	logger.Info("Received request for PendingWarranty")

	var Warranty []models.Warranty
	result := database.DB.Where("status =?", "pending").Find(&Warranty)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Error retrieving pending warranties")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	logger.WithFields(logrus.Fields{
		"warranties": Warranty,
	}).Info("Pending warranties retrieved successfully")
	c.JSON(http.StatusOK, &Warranty)
}

func ReadExcelFile(filePath string) ([]models.WrrantyData, error) {
	logger.Infof("Reading Excel file: %s", filePath)

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		logger.WithError(err).Error("Error opening Excel file")
		return nil, err
	}

	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		err := fmt.Errorf("no sheets found in file")
		logger.WithError(err).Error("Error")
		return nil, err
	}

	sheet := sheetNames[0]
	logger.Infof("Reading data from sheet: %s", sheet)

	rows, err := f.GetRows(sheet)
	if err != nil {
		logger.WithError(err).Error("Error getting rows from sheet")
		return nil, err
	}

	var warrantyData []models.WrrantyData
	for _, row := range rows[1:] {
		if len(row) < 12 {
			logger.WithFields(logrus.Fields{
				"row": row,
			}).Warn("Skipping row with insufficient columns")
			continue
		}

		warrantyData = append(warrantyData, models.WrrantyData{
			Vendor:          row[0],
			WarrantyID:      row[1],
			Name:            row[2],
			MonthlyPrice:    row[3],
			Discount:        row[4],
			AnnualPrice:     row[5],
			PlanDescription: row[6],
			Status:          row[7],
			Picture:         row[8],
			PictureName:     row[12],
		})
	}

	logger.WithFields(logrus.Fields{
		"data": warrantyData,
	}).Info("Excel data read successfully")
	return warrantyData, nil
}

func ExcelUpload(c *gin.Context) {
	logger.Info("Received request for ExcelUpload")

	file, err := c.FormFile("file")
	if err != nil {
		logger.WithError(err).Error("Error getting file from form")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
		return
	}

	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logger.WithError(err).Error("Error saving uploaded file")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
		return
	}

	warrantyData, err := ReadExcelFile(filePath)
	if err != nil {
		logger.WithError(err).Error("Failed to read Excel file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload"})
		return
	}

	for _, data := range warrantyData {
		if err := database.DB.Create(&data).Error; err != nil {
			logger.WithError(err).Error("Failed to insert warranty data")
		}
	}

	logger.Info("File uploaded and data inserted into DB successfully")
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded and data inserted DB successful"})
}

func Warrantydatget(c *gin.Context) {
	logger.Info("Received request for Warrantydatget")

	var Warrantydata []models.WrrantyData
	result := database.DB.Find(&Warrantydata)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Error retrieving warranty data")
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	logger.WithFields(logrus.Fields{
		"data": Warrantydata,
	}).Info("Warranty data retrieved successfully")
	c.JSON(http.StatusOK, &Warrantydata)
}
