package responses

// Loan history response

type Loan struct {
	ID           uint    `json:"ID"`
	LoanAmount   float64 `json:"LoanAmount"`
	LoanCurrency string  `json:"LoanCurrency"`
	InterestRate float64 `json:"InterestRate"`
	StartDate    string  `json:"StartDate"`
	EndDate      string  `json:"EndDate"`
	Status       string  `json:"Status"`
	CreatedAt    string  `json:"CreatedAt"`
	UpdatedAt    string  `json:"UpdatedAt"`
}

type ResponseBody struct {
	CustomerID  uint   `json:"CustomerID"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	PhoneNumber string `json:"PhoneNumber"`
	Gender      string `json:"Gender"`
	DateOfBirth string `json:"DateOfBirth"`
	NationalID  string `json:"NationalID"`
	Address     string `json:"Address"`
	Loans       []Loan `json:"Loans"`
}

type GetLoansByCustomerIDResponse struct {
	ResponseHeader `json:"ResponseHeader"`
	ResponseBody   `json:"ResponseBody"`
}

// Summary

type SummaryResponseBody struct {
	TotalLoans           int     `json:"TotalLoans"`
	TotalLoanedAmount    float64 `json:"TotalLoanedAmount"`
	TotalRepaymentAmount float64 `json:"TotalRepaymentAmount"`
	TotalRepaidAmount    float64 `json:"TotalRepaidAmount"`
	TotalInterest        float64 `json:"TotalInterest"`
}

type SummaryResponse struct {
	ResponseHeader      ResponseHeader      `json:"ResponseHeader"`
	SummaryResponseBody SummaryResponseBody `json:"ResponseBody"`
}

type ResponseHeader struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
	StatusDesc    string `json:"StatusDesc"`
}

type Response[T any] struct {
	ResponseHeader ResponseHeader `json:"ResponseHeader"`
	ResponseBody   []T            `json:"ResponseBody"`
}

type GenericResponse struct {
	ResponseHeader ResponseHeader `json:"ResponseHeader"`
	RResponseBody  RResponseBody  `json:"ResponseBody"`
}

type RResponseBody struct {
	Message            string `json:"Message"`
	MessageDescription string `json:"MessageDescription"`
	ItemId             uint   `json:"ItemId"`
}

func CreateSuccessPostResponse(message, messageDescription string, itemId uint) GenericResponse {
	return GenericResponse{
		ResponseHeader: ResponseHeader{
			StatusCode:    1000,
			StatusMessage: "Success",
			StatusDesc:    "Request Processed Successfully",
		},
		RResponseBody: RResponseBody{
			Message:            message,
			MessageDescription: messageDescription,
			ItemId:             itemId,
		},
	}
}
