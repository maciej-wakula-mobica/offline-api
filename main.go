package main

import (
	"encoding/json"
	"fmt"
	"github.com/WPTechInnovation/wpw-sdk-go/wpwithin/psp/onlineworldpay/types"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/tokens", Tokens)
	router.HandleFunc("/v1/orders", Orders)
	fmt.Println("Serving worldpay web service mock on port "+port)
	log.Fatal(http.ListenAndServe(port, router))
}
func Tokens(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/v1/tokens request received")
	trpm := types.TokenResponsePaymentMethod{
		Type:                              "ObfuscatedCard",
		Name:                              "Joe Bloggs",        // TODO
		ExpiryMonth:                       12,                  // TODO
		ExpiryYear:                        2020,                // TODO
		CardType:                          "AMEX",              // TODO
		MaskedCardNumber:                  "**** **** ** 3434", // TODO
		CardSchemeType:                    "unknown",
		CardSchemeName:                    "unknown",
		CardIssuer:                        "unknown",
		CountryCode:                       "XX",
		CardClass:                         "unknown",
		CardProductTypeDescNonContactless: "unknown",
		CardProductTypeDescContactless:    "unknown",
		Prepaid: "unknown",
	}
	t := types.TokenResponse{
		Reusable: false,
		Token:    "TEST_SU_ae2a9e7b-9583-47ad-b47c-0b7199832a1e", // TODO
		TokenResponsePaymentMethod: trpm,
	}
	json.NewEncoder(w).Encode(t)
}
func Orders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/v1/orders request received")
	orpr := types.OrderResponsePaymentResponse{
		Type:                              "ObfuscatedCard",
		Name:                              "Joe Bloggs",
		ExpiryMonth:                       12,
		ExpiryYear:                        2020,
		CardType:                          "AMEX",
		MaskedCardNumber:                  "**** **** ** 3434",
		CardSchemeType:                    "unknown",
		CardSchemeName:                    "unknown",
		CardIssuer:                        "unknown",
		CountryCode:                       "XX",
		CardClass:                         "unknown",
		CardProductTypeDescNonContactless: "unknown",
		CardProductTypeDescContactless:    "unknown",
		Prepaid: "unknown", //IssueNumber int32 ,StartMonth int32 , StartYear int32
	}
	orrs := types.OrderResponseRiskScore{
		Value: "1",
	}
	t := types.OrderResponse{
		OrderCode:         "ad17e09a-ee11-4478-bef9-30da729198ab",         // TODO
		Token:             "TEST_SU_ae2a9e7b-9583-47ad-b47c-0b7199832a1e", // TODO
		OrderDescription:  "RoboWash - 10 units @ GBP 5.00 per unit.",     // TODO
		Amount:            5000,                                           // TODO
		CurrencyCode:      "GBP",                                          // TODO
		PaymentStatus:     "SUCCESS",
		PaymentResponse:   orpr,
		CustomerOrderCode: "075398ec-b719-4dc0-5402-669dd8b7a073", // TODO
		Environment:       "TEST",
		RiskScore:         orrs, // ResultCodes not supported by current WPW API
	}
	json.NewEncoder(w).Encode(t)
}
