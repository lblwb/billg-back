package invoice

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"os"
)

type InvoiceData struct {
	ShopId         string   `json:"shopId"`
	Sum            float64  `json:"sum"`
	OrderId        string   `json:"orderId"`
	HookUrl        string   `json:"hookUrl"`
	SuccessUrl     string   `json:"successUrl"`
	FailUrl        string   `json:"failUrl"`
	Expire         int      `json:"expire"`
	CustomFields   string   `json:"customFields"`
	Comment        string   `json:"comment"`
	IncludeService []string `json:"includeService"`
	//Signature      string   `json:"signature"`

	// Добавьте поле для хранения сгенерированной подписи
	GeneratedSignature string `json:"-"`
}

func NewInvoiceData(sum float64, orderId string) *InvoiceData {

	//var (hookUrl, successUrl, failUrl string,
	//	sum, expire int,
	//	customFields, comment string,
	//	includeService []string)

	return &InvoiceData{
		//shopId:         os.Getenv("LAVA_SHOP_ID"),
		ShopId:     "2f6da507-19db-4cc7-a9c6-0b59700d3c75",
		Sum:        sum,
		OrderId:    orderId,
		HookUrl:    "https://retry.host/pay.html",  // Webhook url
		SuccessUrl: "https://retry.host/paid.html", // Success url
		FailUrl:    "https://retry.host/fail.html", // Fail url
		Expire:     300,
		//CustomFields:   "",
		Comment:        "Пополнение аккаунта",
		IncludeService: []string{},
	}
}

//func (d *InvoiceData) SetSignature(signature string) {
//	d.Signature = signature
//}

func (d *InvoiceData) GenerateSignature(secretKey string) error {
	jsonData, err := json.Marshal(d)
	if err != nil {
		return err
	}

	signature, err := generateSignature(secretKey, jsonData)
	if err != nil {
		return err
	}

	d.GeneratedSignature = signature
	return nil
}

func (d *InvoiceData) GetBody() []byte {
	jsonData, err := json.Marshal(d)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return nil
	}
	return jsonData
}

func (d *InvoiceData) SendRequest(url string) (*resty.Response, error) {
	client := resty.New()

	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"Signature":    d.GeneratedSignature,
	})

	resp, err := client.R().SetBody(d.GetBody()).Post(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *InvoiceData) HandleResponse(resp *resty.Response) (map[string]interface{}, error) {
	if resp.StatusCode() != http.StatusOK {
		log.Println(string(resp.Body()))
		return nil, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}

	responseData := map[string]interface{}{}
	err := json.Unmarshal(resp.Body(), &responseData)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return responseData, nil
}

func generateSignature(secretKey string, jsonData []byte) (string, error) {
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write(jsonData)
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return signature, nil
}

func (d *InvoiceData) CreateInvoice() (map[string]interface{}, error) {
	// Replace with your actual values
	secretKey := os.Getenv("LAVA_SECRET_KEY")
	url := "https://api.lava.ru/business/invoice/create"
	err := d.GenerateSignature(secretKey)
	if err != nil {
		fmt.Println("Error generating signature:", err)
		return nil, err
	}

	resp, err := d.SendRequest(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	responseData, err := d.HandleResponse(resp)
	if err != nil {
		fmt.Println("Error handling response:", err)
		return nil, err
	}
	fmt.Println("Response data:", responseData)
	return responseData, nil
}
