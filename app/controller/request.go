package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Status   string      `json:"status"`
	Code     interface{} `json:"code"`
	Message  string      `json:"message"`
	Revision string      `json:"revision"`
	Method   string      `json:"method"`
	Data     []struct {
		LocalOrderHeaderID       string      `json:"localOrderHeaderId"`
		OrderHeaderID            string      `json:"orderHeaderId"`
		StoreSettingID           interface{} `json:"storeSettingId"`
		StoreSettingName         string      `json:"storeSettingName"`
		TableNo                  string      `json:"tableNo"`
		TableName                string      `json:"tableName"`
		TableCategory            string      `json:"tableCategory"`
		ApplyPlans               string      `json:"applyPlans"`
		Numbers                  string      `json:"numbers"`
		NumbersMale              interface{} `json:"numbersMale"`
		NumbersFemale            interface{} `json:"numbersFemale"`
		NumbersUnknown           interface{} `json:"numbersUnknown"`
		EnterDateTime            string      `json:"enterDateTime"`
		LastOrderDateTime        string      `json:"lastOrderDateTime"`
		CheckoutDateTime         string      `json:"checkoutDateTime"`
		Subtotal                 string      `json:"subtotal"`
		SubtotalTaxIncluded      string      `json:"subtotalTaxIncluded"`
		SubtotalTaxExcluded      string      `json:"subtotalTaxExcluded"`
		SubtotalTaxExempt        string      `json:"subtotalTaxExempt"`
		SubtotalForServiceCharge string      `json:"subtotalForServiceCharge"`
		Amount                   string      `json:"amount"`
		Tax                      string      `json:"tax"`
		TaxExclude               string      `json:"taxExclude"`
		SellTaxDivision          string      `json:"sellTaxDivision"`
		TaxRate                  string      `json:"taxRate"`
		RoundingDivision         string      `json:"roundingDivision"`
		DiscountPrice            string      `json:"discountPrice"`
		DiscountRate             string      `json:"discountRate"`
		DiscountDivision         string      `json:"discountDivision"`
		ServiceChargeRate        string      `json:"serviceChargeRate"`
		TableChargePerPerson     string      `json:"tableChargePerPerson"`
		TimeZonePriceType        string      `json:"timeZonePriceType"`
		TimeZonePriceDivision    string      `json:"timeZonePriceDivision"`
		TimeZonePriceValue       string      `json:"timeZonePriceValue"`
		Total                    string      `json:"total"`
		GroupingID               interface{} `json:"groupingId"`
		GroupingIds              string      `json:"groupingIds"`
		GroupingName             string      `json:"groupingName"`
		StaffID                  interface{} `json:"staffId"`
		StaffName                interface{} `json:"staffName"`
		UUID                     string      `json:"uuid"`
		ReservationID            interface{} `json:"reservationId"`
		Memo                     interface{} `json:"memo"`
		Status                   string      `json:"status"`
		PlanOperatingStatus      string      `json:"planOperatingStatus"`
		Revision                 string      `json:"revision"`
		Barcode                  string      `json:"barcode"`
		SmaregiTransactionHeadID interface{} `json:"smaregiTransactionHeadId"`
		UnCaterCount             string      `json:"unCaterCount"`
	} `json:"data"`
}

func filter(result *Response) []Result {
	var Results []Result
	// TODO: 条件部分は関数に切り出した方がよさげ
	for _, d := range result.Data {
		if d.Status != "1" && d.Status != "4" {
			continue
		}
		if strings.Contains(d.TableName, "S") {
			continue
		}
		Results = append(Results, Result{Seat: d.TableName, Category: d.TableCategory})
	}
	return Results
}

// もう少し関数細かくしたい
func request() ([]Result, error) {
	u := os.Getenv("URL")
	contractID := os.Getenv("CONTRACT_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")
	storeID := os.Getenv("STORE_ID")

	now := time.Now().Format("20060102030405")

	re := regexp.MustCompile(`^........`)
	revision := re.FindAllStringSubmatch(now, -1)[0][0] + "110000"

	r, err := strconv.Atoi(revision)
	q := fmt.Sprintf(`service=OrderService&method=getUpdatedOrderHeaderList&params={"revision":"%v"}`, r)
	q = url.PathEscape(q)

	req, err := http.NewRequest("GET", u+"?"+q, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-access-token", accessToken)
	req.Header.Set("X-contract-id", contractID)
	req.Header.Set("X-store-id", storeID)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	var result Response
	if err := json.Unmarshal(byteArray, &result); err != nil {
		return nil, err
	}
	seat := filter(&result)
	log.Print(seat)
	return seat, nil
}
