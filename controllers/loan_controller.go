package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"loanapi/configs"
	"loanapi/models"
	"loanapi/responses"
	"loanapi/utils"
	"net/http"
	"strconv"
	"time"
)

func CreateLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print the received loan object for debugging
	fmt.Printf("Received Loan: %+v\n", loan)

	// Assuming configs.DB is properly set up to interact with the database
	if err := configs.DB.Create(&loan).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan"})
		return
	}

	res := responses.CreateSuccessPostResponse("Success", "Request Processed Successfully", loan.ID)

	ctx.JSON(http.StatusCreated, res)
}

// GetLoansByCustomerID Api to get loan history
func GetLoansByCustomerID(c *gin.Context) {
	customerID := c.Param("customerID")

	var customer models.Customer
	if err := configs.DB.First(&customer, customerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	var loans []models.Loan
	if err := configs.DB.Where("customer_id = ?", customerID).Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch loans"})
		return
	}

	var loanResponses []responses.Loan
	for _, loan := range loans {
		startDate, err := time.Parse(time.RFC3339, loan.StartDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse StartDate"})
			return
		}
		endDate, err := time.Parse(time.RFC3339, loan.EndDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse EndDate"})
			return
		}
		loanResponses = append(loanResponses, responses.Loan{
			ID:           loan.ID,
			LoanAmount:   loan.LoanAmount,
			LoanCurrency: loan.LoanCurrency,
			InterestRate: loan.InterestRate,
			StartDate:    startDate.Format("2006-01-02"),
			EndDate:      endDate.Format("2006-01-02"),
			Status:       loan.Status,
			CreatedAt:    loan.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    loan.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response := responses.GetLoansByCustomerIDResponse{
		ResponseHeader: responses.ResponseHeader{
			StatusCode:    1000,
			StatusMessage: "Success",
			StatusDesc:    "OK",
		},
		ResponseBody: responses.ResponseBody{
			CustomerID:  customer.ID,
			FirstName:   customer.FirstName,
			LastName:    customer.LastName,
			PhoneNumber: customer.PhoneNumber,
			Gender:      customer.Gender,
			DateOfBirth: customer.DateOfBirth,
			NationalID:  customer.NationalID,
			Address:     customer.Address,
			Loans:       loanResponses,
		},
	}

	c.JSON(http.StatusOK, response)
}

func SummarizeLoans(c *gin.Context) {
	status := c.Query("status")
	var loans []models.Loan

	// Filter loans based on the status parameter
	if status != "" {
		if err := configs.DB.Preload("Payments").Where("status = ?", status).Find(&loans).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loans"})
			return
		}
	} else {
		if err := configs.DB.Preload("Payments").Find(&loans).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve loans"})
			return
		}
	}

	totalLoans := len(loans)
	totalLoanedAmount := 0.0
	totalRepaidAmount := 0.0
	totalInterest := 0.0

	for _, loan := range loans {
		totalLoanedAmount += loan.LoanAmount
		for _, payment := range loan.Payments {
			totalRepaidAmount += payment.PaymentAmount
		}
		totalInterest += loan.LoanAmount * (loan.InterestRate / 100)
	}
	totalRepaymentAmount := totalLoanedAmount + totalInterest

	response := responses.SummaryResponse{
		ResponseHeader: responses.ResponseHeader{
			StatusCode:    1000,
			StatusMessage: "Success",
			StatusDesc:    "OK",
		},
		SummaryResponseBody: responses.SummaryResponseBody{
			TotalLoans:           totalLoans,
			TotalLoanedAmount:    totalLoanedAmount,
			TotalRepaymentAmount: totalRepaymentAmount,
			TotalInterest:        totalInterest,
			TotalRepaidAmount:    totalRepaidAmount,
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetLoans handles fetching loans with pagination
func GetLoans(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	var loans []models.Loan
	if result := configs.DB.Limit(limit).Offset(offset).Find(&loans); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	var totalLoans int64
	configs.DB.Model(&models.Loan{}).Count(&totalLoans)
	type ResponseLoan struct {
		Id       uint `json:"Id"`
		Customer struct {
			CustomerID        uint   `json:"CustomerID"`
			CustomerFirstName string `json:"CustomerFirstName"`
			CustomerLastName  string `json:"CustomerLastName"`
		} `json:"Customer"`
		LoanAmount   float64   `json:"LoanAmount"`
		LoanCurrency string    `json:"LoanCurrency"`
		InterestRate float64   `json:"InterestRate"`
		StartDate    string    `json:"StartDate"`
		EndDate      string    `json:"EndDate"`
		Status       string    `json:"Status"`
		CreatedAt    time.Time `json:"CreatedAt"`
		UpdatedAt    time.Time `json:"UpdatedAt"`
	}
	outputDateFormat := "2006-01-02"

	var responseLoans []ResponseLoan
	for _, loan := range loans {
		var customer models.Customer
		configs.DB.First(&customer, loan.CustomerID)
		startDate, _ := time.Parse(time.RFC3339, loan.StartDate)
		endDate, _ := time.Parse(time.RFC3339, loan.EndDate)
		responseLoan := ResponseLoan{
			Id:           loan.ID,
			LoanAmount:   loan.LoanAmount,
			LoanCurrency: loan.LoanCurrency,
			InterestRate: loan.InterestRate,
			Status:       loan.Status,
			StartDate:    startDate.Format(outputDateFormat), // Format StartDate
			EndDate:      endDate.Format(outputDateFormat),   // Format EndDate
			CreatedAt:    loan.CreatedAt,
			UpdatedAt:    loan.UpdatedAt,
		}
		responseLoan.Customer.CustomerID = customer.ID
		responseLoan.Customer.CustomerFirstName = customer.FirstName
		responseLoan.Customer.CustomerLastName = customer.LastName
		responseLoans = append(responseLoans, responseLoan)
	}

	// Calculate total pages
	totalPages := int(totalLoans) / limit
	if int(totalLoans)%limit != 0 {
		totalPages++
	}

	// Prepare the response
	response := gin.H{
		"CurrentPage": page,
		"TotalPages":  totalPages,
		"TotalItems":  totalLoans,
		"Loans":       responseLoans,
	}

	res := utils.CreateMapResponse("Success", "Request Processed Successfully", response)
	c.JSON(http.StatusOK, res)
}
func UpdateLoans(c *gin.Context) {

	id := c.Param("id")
	var loan models.Loan

	if err := configs.DB.First(&loan, id).Error; err != nil {
		response := utils.CreateErrorResponse("An error was encountered while creating the customer", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, utils.CreateErrorResponse("Invalid input", err.Error()))
		return
	}
	configs.DB.Save(&loan)
	response := responses.CreateSuccessPostResponse("Customer Created Successfully", "Ok", loan.ID)
	c.JSON(http.StatusCreated, response)
}

func GetLoanById(c *gin.Context) {
	id := c.Param("id")
	var loan models.Loan

	// Fetch the loan along with the associated customer details
	if err := configs.DB.Preload("Customer").First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	// Prepare the response
	res := models.LoanRes{
		ID:         loan.ID,
		CustomerID: loan.CustomerID,
		CustomerRes: models.CustomerRes{
			ID:               loan.Customer.ID,
			FirstName:        loan.Customer.FirstName,
			LastName:         loan.Customer.LastName,
			Gender:           loan.Customer.Gender,
			DateOfBirth:      loan.Customer.DateOfBirth,
			NationalID:       loan.Customer.NationalID,
			Address:          loan.Customer.Address,
			City:             loan.Customer.City,
			Province:         loan.Customer.Province,
			PostalCode:       loan.Customer.PostalCode,
			Country:          loan.Customer.Country,
			PhoneNumber:      loan.Customer.PhoneNumber,
			Email:            loan.Customer.Email,
			EmploymentStatus: loan.Customer.EmploymentStatus,
			AnnualIncome:     loan.Customer.AnnualIncome,
			CreditScore:      loan.Customer.CreditScore,
		},
		LoanAmount:   loan.LoanAmount,
		LoanCurrency: loan.LoanCurrency,
		InterestRate: loan.InterestRate,
		StartDate:    loan.StartDate,
		EndDate:      loan.EndDate,
		Status:       loan.Status,
	}

	// Create a response
	xm := utils.CreateGenericSuccessResponse("Success", "Request Processed Successfully", res)
	c.JSON(http.StatusOK, xm)
}
