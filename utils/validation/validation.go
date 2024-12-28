package validation

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	// Register custom validate methods
	govalidator.TagMap["required"] = govalidator.Validator(func(str string) bool {
		return len(str) > 0
	})

	//govalidator.TagMap["bool"] = govalidator.Validator(func(str string) bool {
	//	i, err := strconv.Atoi(str)
	//	if err != nil {
	//		return false
	//	}
	//	if i == constant.FEMALE || i == constant.MALE {
	//		return true
	//	}
	//	return false
	//})
	//govalidator.TagMap["meal"] = govalidator.Validator(func(str string) bool {
	//	if str == constant.BREAKFAST || str == constant.LUNCH || str == constant.DINER || str == constant.SNACK {
	//		return true
	//	}
	//	return false
	//})

}
func CommonlyValidate(obj interface{}) (bool, error) {
	isvalid, err := govalidator.ValidateStruct(obj)
	if err != nil || !isvalid {
		return false, err
	} else {
		return true, nil
	}
}
