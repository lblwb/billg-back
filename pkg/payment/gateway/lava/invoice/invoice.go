package invoice

//
//import (
//	"crypto/hmac"
//	"crypto/sha256"
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"github.com/go-resty/resty/v2"
//	"log"
//	"net/http"
//	"os"
//	"reflect"
//	"sort"
//)
//
//type LavaInvoice struct {
//	Data *InvoiceParams
//}
//
//// InvoiceParams представляет параметры для создания Invoice
//type InvoiceParams struct {
//	ShopId       string `json:"shopId"`
//	OrderID      string `json:"orderId"`
//	Comment      string `json:"comment"`
//	FailURL      string `json:"failUrl"`
//	SuccessURL   string `json:"successUrl"`
//	HookURL      string `json:"hookUrl"`
//	Expire       int    `json:"expire"`
//	CustomFields string `json:"customFields"`
//	//Signature    string  `json:"signature"`
//	Sum float64 `json:"sum"`
//}
//
//// Service represents a supported payment method
//type Service struct {
//	Name string `json:"name"`
//}
//
//// InvoiceResponse represents the response from the CreateInvoice API
//type InvoiceResponse struct {
//	Data        *InvoiceData `json:"data"`
//	Error       interface{}  `json:"error"`
//	Status      int64        `json:"status"`
//	StatusCheck bool         `json:"status_check"`
//}
//
//// InvoiceData represents the invoice details within the response
//type InvoiceData struct {
//	ShopId       string `json:"shopId"`
//	Sum          int    `json:"sum"`
//	Comment      string `json:"comment"`
//	FailUrl      string `json:"failUrl"`
//	SuccessUrl   string `json:"successUrl"`
//	HookUrl      string `json:"hookUrl"`
//	Expire       int    `json:"expire"`
//	CustomFields string `json:"customFields"`
//	OrderId      string `json:"orderId"`
//	Signature    string `json:"signature"`
//}
//
//// NewInvoiceParams creates a new InvoiceParams instance
//func NewInvoiceParams(sum float64, orderID string) *LavaInvoice {
//	return &LavaInvoice{
//		Data: &InvoiceParams{
//			FailURL:    "https://retry.host/fail.html",
//			SuccessURL: "https://retry.host/paid.html",
//			HookURL:    "https://retry.host/pay.html",
//			Expire:     300,
//			//CustomFields: "",
//			Sum:     sum,
//			OrderID: orderID,
//			ShopId:  os.Getenv("LAVA_SHOP_ID"),
//		},
//	}
//}
//
//func (li LavaInvoice) ToMap() map[string]interface{} {
//	var s = li.Data
//	//
//	t := reflect.TypeOf(s)
//	v := reflect.ValueOf(s)
//
//	if t.Kind() == reflect.Ptr {
//		t = t.Elem()
//		v = v.Elem()
//	}
//
//	if t.Kind() != reflect.Struct {
//		panic("ToMap expects a struct")
//	}
//
//	m := make(map[string]interface{}, t.NumField())
//
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		value := v.Field(i)
//
//		// Ignore unexported fields
//		if field.PkgPath != "" {
//			continue
//		}
//
//		m[field.Name] = value.String()
//	}
//
//	return m
//}
//
//func (li LavaInvoice) ToJson() ([]byte, error) {
//	marshal, err := json.Marshal(li.Data)
//	if err != nil {
//		return []byte{}, err
//	}
//	return marshal, nil
//}
//
//// SetFailURL sets the FailURL for the InvoiceParams
//func (li *LavaInvoice) SetFailURL(failURL string) {
//	li.Data.FailURL = failURL
//}
//
//// SetSuccessURL sets the SuccessURL for the InvoiceParams
//func (li *LavaInvoice) SetSuccessURL(successURL string) {
//	li.Data.SuccessURL = successURL
//}
//
//// SetHookURL sets the HookURL for the InvoiceParams
//func (li *LavaInvoice) SetHookURL(hookURL string) {
//	li.Data.HookURL = hookURL
//}
//
//// SetExpire sets the Expire for the InvoiceParams
//func (li *LavaInvoice) SetExpire(expire int) {
//	li.Data.Expire = expire
//}
//
//// SetCustomFields sets the CustomFields for the InvoiceParams
//func (li *LavaInvoice) SetCustomFields(customFields string) {
//	li.Data.CustomFields = customFields
//}
//
//// SetComment sets the CustomFields for the InvoiceParams
//func (li *LavaInvoice) SetComment(comment string) {
//	li.Data.Comment = comment
//}
//
//func (li *LavaInvoice) sortDict(data map[string]interface{}) map[string]interface{} {
//	sortedKeys := make([]string, 0, len(data))
//	for k := range data {
//		sortedKeys = append(sortedKeys, k)
//	}
//	sort.Strings(sortedKeys)
//
//	result := make(map[string]interface{})
//	for _, k := range sortedKeys {
//		result[k] = data[k]
//	}
//	return result
//}
//
//func (li *LavaInvoice) generateSign(data map[string]interface{}) string {
//	var secretKey = os.Getenv("LAVA_SECRET_KEY")
//	//
//	sortedData := li.sortDict(data)
//	//
//	jsonData, err := json.Marshal(sortedData)
//	if err != nil {
//		fmt.Println("Error marshalling data:", err)
//		return ""
//	}
//
//	h := hmac.New(sha256.New, []byte(secretKey))
//	_, err = h.Write(jsonData)
//	if err != nil {
//		fmt.Println("Error writing to hmac:", err)
//		return ""
//	}
//	signature := hex.EncodeToString(h.Sum(nil))
//	return signature
//}
//
////func (li *LavaInvoice) generateSignature() (string, error) {
////	var SecretKey = os.Getenv("LAVA_SECRET_KEY")
////	// Compute HMAC signature
////	h := hmac.New(sha256.New, []byte(SecretKey))
////	marshal, err := li.ToJson()
////	if err != nil {
////		return "", err
////	}
////	h.Write(marshal)
////	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
////	//ip.Signature = signature
////
////	log.Println("generateSignature", signature)
////
////	return signature, nil
////}
//
//func (li *LavaInvoice) sendRequest(client *resty.Client, url string, data map[string]interface{}, signature string) (*resty.Response, error) {
//	// Set headers
//	client.SetHeaders(map[string]string{
//		"Signature":    signature,
//		"Accept":       "application/json",
//		"Content-Type": "application/json",
//	})
//
//	// Send POST request with JSON data
//	resp, err := client.R().SetBody(data).Post(url)
//	if err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}
//
//// CreateInvoice creates a new Invoice using the provided parameters
//func (li *LavaInvoice) CreateInvoice() (*InvoiceResponse, error) {
//	url := "https://api.lava.ru/business/invoice/create"
//	// Create a Resty client
//	hc := resty.New()
//	shopID := os.Getenv("LAVA_SHOP_ID")
//	if shopID == "" {
//		return nil, fmt.Errorf("missing SHOP_ID environment variable")
//	}
//
//	li.Data.ShopId = shopID
//
//	//log.Println(shopID)
//
//	signature := li.generateSign(li.ToMap())
//	resp, err := li.sendRequest(hc, url, li.ToMap(), signature)
//	if err != nil {
//		return nil, err
//	}
//
//	if err != nil {
//		log.Println("Error sending request:", err)
//		return nil, err
//	}
//
//	// Handle response (unchanged)
//	if resp.StatusCode() != http.StatusOK {
//		log.Printf("Error: API responded with status code %d\n", resp.StatusCode())
//		log.Println(resp)
//		return nil, err
//	}
//
//	log.Println("Response", string(resp.Body()))
//
//	var responseData InvoiceResponse
//	err = json.Unmarshal(resp.Body(), &responseData)
//	if err != nil {
//		log.Println("Error decoding response:", err)
//		return nil, err
//	}
//
//	log.Println("Response data:", responseData)
//
//	return nil, nil
//}
