package utils

import (
	"errors"

	"github.com/scottdiddy/sal-app/src/repository"
)

//Checks if the supplied merchantID is empty or if the merchant exists. It
//returns the slice of skuids of the products of the merchant and error if any.
func CheckMerchant(merchantID string) ([]string, error) {
	merchantProductSkuids, merchantExists := repository.MerchantData[merchantID]

	//checks if the merchantID is empty
	if merchantID == ":merchantID" {
		return []string{}, errors.New("no merchant ID provided")
	}

	//checks if the merchant exists
	if !merchantExists {
		return []string{}, errors.New("merchant does not exist")
	}
	return merchantProductSkuids, nil
}
