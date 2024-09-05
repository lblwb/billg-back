package lava

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Invoice struct {
	ShopId  string  `json:"shopId"`
	OrderId string  `json:"orderId"`
	Sum     float64 `json:"sum"`
	Comment string  `json:"comment"`
	Expire  int     `json:"expire"`
	Status  string  `json:"status"`
	Url     string  `json:"url"`
}

type InvoiceResponse struct {
	Status      int                 `json:"status"`
	StatusCheck bool                `json:"status_check"`
	Data        InvoiceResponseData `json:"data"`
}

type InvoiceResponseData struct {
	ID             string      `json:"id"`
	Amount         float64     `json:"amount"`
	Expired        string      `json:"expired"`
	Status         int         `json:"status"`
	ShopID         string      `json:"shop_id"`
	URL            string      `json:"url"`
	Comment        string      `json:"comment"`
	MerchantName   string      `json:"merchantName"`
	ExcludeService interface{} `json:"exclude_service"`
	IncludeService interface{} `json:"include_service"`
}

type Client struct {
	ProjectId string
	SecretKey string
	Client    *http.Client
}

func NewClient(projectId, secretKey string) *Client {
	return &Client{
		ProjectId: projectId,
		SecretKey: secretKey,
		Client:    &http.Client{},
	}
}

func (c *Client) CreateInvoice(amount float64, comment, orderID string, expire int) (InvoiceResponse, error) {
	if orderID == "" {
		orderID = fmt.Sprintf("LavaBusiness-%d-%d", int(time.Now().UnixNano()/1000000), int(time.Now().UnixNano()%1000000))
	}

	if expire > 43200 {
		expire = 43200
	} else if expire < 1 {
		expire = 1
	}

	params := map[string]interface{}{
		"shopId":  c.ProjectId,
		"orderId": orderID,
		"sum":     amount,
		"comment": comment,
		"expire":  expire,
	}

	headers, err := c.createSignHeaders(params)
	if err != nil {
		return InvoiceResponse{}, err
	}

	marshal, err := json.Marshal(params)
	if err != nil {
		return InvoiceResponse{}, err
	}
	req, err := http.NewRequest("POST", "https://api.lava.ru/business/invoice/create", bytes.NewBufferString(string(marshal)))
	if err != nil {
		return InvoiceResponse{}, err
	}

	req.Header.Set("Signature", headers["Signature"])
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return InvoiceResponse{}, err
	}

	defer resp.Body.Close()

	var invoice InvoiceResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return InvoiceResponse{}, err
	}
	if err := json.Unmarshal(body, &invoice); err != nil {
		return InvoiceResponse{}, err
	}

	//log.Println(invoice.Data)

	return invoice, nil
}

func (c *Client) createSignHeaders(params map[string]interface{}) (map[string]string, error) {
	secretKey := []byte(c.SecretKey)
	paramsStr, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	//hash := hmac.New(sha256.New, secretKey, paramsStr)
	hash := hmac.New(sha256.New, secretKey)
	hash.Write(paramsStr)
	signature := hex.EncodeToString(hash.Sum(nil))
	//hexSignature := hex.EncodeToString(signature)

	return map[string]string{
		"Signature":    signature,
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}, nil
}
