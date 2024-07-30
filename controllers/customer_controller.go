package controllers

import (
	"github.com/gin-gonic/gin"
	"loanapi/configs"
	"loanapi/models"
	"loanapi/responses"
	"loanapi/utils"
	"net/http"
	"strconv"
	_ "time"
)

func GetCustomers(c *gin.Context) {
	// Get page and limit query parameters, defaulting to page 1 and limit 10 if not provided
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Calculate the offset
	offset := (page - 1) * limit

	// Retrieve customers from the database with pagination
	var customers []models.Customer
	if result := configs.DB.Limit(limit).Offset(offset).Find(&customers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Count the total number of customers
	var totalCustomers int64
	configs.DB.Model(&models.Customer{}).Count(&totalCustomers)

	// Prepare the response
	var cusRes []models.CustomerRes
	for _, cus := range customers {
		cusRes = append(cusRes, models.CustomerRes{
			ID:               cus.ID,
			FirstName:        cus.FirstName,
			LastName:         cus.LastName,
			Gender:           cus.Gender,
			DateOfBirth:      cus.DateOfBirth,
			NationalID:       cus.NationalID,
			Address:          cus.Address,
			City:             cus.City,
			Province:         cus.Province,
			PostalCode:       cus.PostalCode,
			Country:          cus.Country,
			PhoneNumber:      cus.PhoneNumber,
			Email:            cus.Email,
			EmploymentStatus: cus.EmploymentStatus,
			AnnualIncome:     cus.AnnualIncome,
			CreditScore:      cus.CreditScore,
			CreatedAt:        cus.CreatedAt,
			UpdatedAt:        cus.UpdatedAt,
		})
	}

	// Calculate the total number of pages
	totalPages := int(totalCustomers) / limit
	if int(totalCustomers)%limit != 0 {
		totalPages++
	}

	// Response structure
	response := gin.H{
		"CurrentPage": page,
		"TotalPages":  totalPages,
		"TotalItems":  totalCustomers,
		"Customers":   cusRes,
	}

	resss := utils.CreateMapResponse("Success", "Request Processed Successfully", response)
	c.JSON(http.StatusOK, resss)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := configs.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	res := models.CustomerRes{
		ID:               customer.ID,
		FirstName:        customer.FirstName,
		LastName:         customer.LastName,
		Gender:           customer.Gender,
		DateOfBirth:      customer.DateOfBirth,
		NationalID:       customer.NationalID,
		Address:          customer.Address,
		City:             customer.City,
		Province:         customer.Province,
		PostalCode:       customer.PostalCode,
		Country:          customer.Country,
		PhoneNumber:      customer.PhoneNumber,
		Email:            customer.Email,
		EmploymentStatus: customer.EmploymentStatus,
		AnnualIncome:     customer.AnnualIncome,
		CreditScore:      customer.CreditScore,
		CreatedAt:        customer.CreatedAt,
		UpdatedAt:        customer.UpdatedAt,
	}
	xm := utils.CreateGenericSuccessResponse("Success", "Request Processed Successfully", res)
	c.JSON(http.StatusOK, xm)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, utils.CreateErrorResponse("Invalid input", err.Error()))
		return
	}
	if err := configs.DB.Create(&customer).Error; err != nil {
		response := utils.CreateErrorResponse("An error was encountered while creating the customer", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := responses.CreateSuccessPostResponse("Customer Created Successfully", "Ok", customer.ID)
	c.JSON(http.StatusCreated, response)
}

func UpdateCustomer(c *gin.Context) {

	id := c.Param("id")
	var customer models.Customer

	if err := configs.DB.First(&customer, id).Error; err != nil {
		response := utils.CreateErrorResponse("An error was encountered while creating the customer", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, utils.CreateErrorResponse("Invalid input", err.Error()))
		return
	}
	configs.DB.Save(&customer)
	response := responses.CreateSuccessPostResponse("Customer Created Successfully", "Ok", customer.ID)
	c.JSON(http.StatusCreated, response)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := configs.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	configs.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
