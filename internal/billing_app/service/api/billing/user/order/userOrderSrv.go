package order

import (
	"backend/internal/billing_app/models/order"
	"backend/internal/billing_app/models/tariff"
	"backend/internal/billing_app/models/user"
	"backend/internal/billing_app/service/api/auth/jwt_auth"
	"backend/internal/database"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
)

type BillUsrOrderService struct {
	db                  *database.StorageDb
	jwtAuth             *jwt_auth.JwtAuths
	svcOrdEntity        *order.ServiceOrdersEntity
	userEntity          *user.UsersEntity
	userOrderSvcEntity  *order.UserOrderSvcEntity
	tariffServiceEntity *tariff.TariffsServiceEntity
	usersBalanceEntity  *user.UsersBalanceEntity
}

type DNSRecord struct {
	CreatedOn  time.Time   `json:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty"`
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	Content    string      `json:"content,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	ID         string      `json:"id,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	//Priority   *uint16     `json:"priority,omitempty"`
	TTL       int   `json:"ttl,omitempty"`
	Proxied   *bool `json:"proxied,omitempty"`
	Proxiable bool  `json:"proxiable,omitempty"`
	Locked    bool  `json:"locked,omitempty"`
	//Comment   string `json:"comment,omitempty"` // the server will omit the comment field when the comment is empty
	//Tags       []string    `json:"tags,omitempty"`
}

type ResponseDnsInfo struct {
	Result struct {
		Zone interface{} `json:"zone"`
	} `json:"result"`
	Success bool `json:"success"`
}

type ResponseDnsRecInfo struct {
	Result struct {
		DnsRecords []DNSRecord `json:"dns_records"`
	} `json:"result"`
	Success bool `json:"success"`
}

func NewBillUsrOrderService(db *database.StorageDb) *BillUsrOrderService {
	return &BillUsrOrderService{
		//orderEntity:
		db:                  db,
		jwtAuth:             jwt_auth.NewJwtAuths(db),
		userEntity:          user.NewUsersEntity(db),
		svcOrdEntity:        order.NewServiceOrdersEntity(db),
		userOrderSvcEntity:  order.NewUserOrderSvcEntity(db),
		tariffServiceEntity: tariff.NewTariffsServiceEntity(db),
		usersBalanceEntity:  user.NewUserBalanceEntity(db),
		//orderEntity: order.NewUserOrderSvcService(db),
	}
}

func (uos BillUsrOrderService) getDnsRecordsOrderInfo(domainDns string) ([]DNSRecord, error) {
	var data ResponseDnsRecInfo

	if domainDns == "" {
		return nil, errors.New("not domain name")
	}

	var dnsPanelUrl = os.Getenv("DNS_PANEL_URL")
	var dnsPanelKey = os.Getenv("DNS_PANEL_KEY")

	var fullUrl = fmt.Sprintf("%s/%s/%s/%s/%s", dnsPanelUrl, "__sc_api", "zones/byCustomer/2/get", domainDns, "dns_records")
	//http://localhost:9761/__sc_api/zones/byCustomer/2/get/1212stars.site
	log.Println(fullUrl)
	response, err := resty.New().R().
		SetHeader("CsHsG", dnsPanelKey).
		Get(fullUrl)
	if err != nil {
		return nil, err
	}
	//
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		return nil, err
	}

	return data.Result.DnsRecords, nil

}

func (uos BillUsrOrderService) getDnsOrderInfo(domainDns string) (interface{}, error) {
	var data ResponseDnsInfo

	if domainDns == "" {
		return nil, errors.New("not domain name")
	}

	var dnsPanelUrl = os.Getenv("DNS_PANEL_URL")
	var dnsPanelKey = os.Getenv("DNS_PANEL_KEY")

	var fullUrl = fmt.Sprintf("%s/%s/%s/%s", dnsPanelUrl, "__sc_api", "zones/byCustomer/2/get", domainDns)
	//var fullUrl = fmt.Sprintf("%s%s%s%s", dnsPanelUrl, "/__sc_api", "/zones/byCustomer/2/get/", domainDns)
	//http://localhost:9761/__sc_api/zones/byCustomer/2/get/1212stars.site
	log.Println(fullUrl)
	response, err := resty.New().R().
		SetHeader("CsHsG", dnsPanelKey).
		Get(fullUrl)
	if err != nil {
		return nil, err
	}
	//
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		return nil, err
	}

	return data.Result.Zone, nil

}

func (uos BillUsrOrderService) createNewDnsOrder(domainDns string) (interface{}, error) {
	var data ResponseDnsInfo

	if domainDns == "" {
		return nil, errors.New("not domain name")
	}

	var dnsPanelUrl = os.Getenv("DNS_PANEL_URL")
	var dnsPanelKey = os.Getenv("DNS_PANEL_KEY")

	//http://localhost:9761/__sc_api/zones/byCustomer/2/create/retragradmercuriy.site
	var fullUrl = fmt.Sprintf("%s/%s/%s/%s", dnsPanelUrl, "__sc_api", "zones/byCustomer/2/create", domainDns)
	//var fullUrl = fmt.Sprintf("%s%s%s%s", dnsPanelUrl, "/__sc_api", "/zones/byCustomer/2/get/", domainDns)
	//http://localhost:9761/__sc_api/zones/byCustomer/2/get/1212stars.site
	log.Println(fullUrl)
	response, err := resty.New().R().
		SetHeader("CsHsG", dnsPanelKey).
		Post(fullUrl)
	if err != nil {
		return nil, err
	}
	//
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		return nil, err
	}

	return data.Result.Zone, nil

}

func (uos BillUsrOrderService) OrderServiceShow(c *fiber.Ctx) error {
	slugName := c.Params("slug")
	userOrderSvc, err := uos.userOrderSvcEntity.GetUserOrderSvcBySlug(slugName)
	if err != nil {
		return c.JSON(fiber.Map{
			"success":   false,
			"error_msg": "not found",
			"order":     nil,
		})
	}

	log.Println(userOrderSvc.Services[0], userOrderSvc.Services[0].Type)

	if userOrderSvc.Services[0].Type != 4 {
		return c.JSON(fiber.Map{
			"success":   false,
			"error_msg": "not found",
			"order":     nil,
		})
	}

	//fmt.Println(userOrderSvc.OrderParams)

	var domainValue = ""

	if userOrderSvc.Services[0].Resource != "" {
		domainValue = userOrderSvc.Services[0].Resource
	}

	//if len(userOrderSvc.OrderParams) > 0 {
	//	var data map[string]map[string]interface{}
	//	if err := json.Unmarshal([]byte(userOrderSvc.OrderParams), &data); err != nil {
	//		fmt.Println("Ошибка при разборе JSON:", err)
	//	}
	//	if data != nil && len(data["domain name"]) > 1 && data["domain name"]["value"] != "" {
	//		domainValue = data["domain name"]["value"].(string)
	//	}
	//}

	if domainValue != "" {

		//fmt.Println(domainValue)

		dnsOrderInfo, err := uos.getDnsOrderInfo(domainValue)
		//fmt.Println(dnsOrderInfo, err)
		if err != nil {
			return c.JSON(fiber.Map{
				"success": false,
				//"error_msg": "not found dns",
				"error_msg": err,
				"order":     nil,
			})
		}

		dnsRecordsOrderInfo, err := uos.getDnsRecordsOrderInfo(domainValue)
		if err != nil {
			return c.JSON(fiber.Map{
				"success": false,
				//"error_msg": "not found dns",
				"error_msg": err,
				"order":     nil,
			})
		}

		//
		return c.JSON(fiber.Map{
			"success":       true,
			"order":         userOrderSvc,
			"order_dns":     dnsOrderInfo,
			"order_dns_rec": dnsRecordsOrderInfo,
		})

	} else {
		return c.JSON(fiber.Map{
			"success":   false,
			"error_msg": "not found domain",
			"order":     nil,
		})
	}

}
