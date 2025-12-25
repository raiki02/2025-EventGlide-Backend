package ginx

import (
	"github.com/go-playground/validator/v10"
	"github.com/raiki02/EG/api/req"
)

const (
	HolderTypePersonal = "个人"
)

var validate = validator.New()

func GetValidate() *validator.Validate {
	return validate
}

func validateRequest(any any) error {
	return validate.Struct(any)
}

func InitValidation() {
	validate.RegisterStructValidation(func(sl validator.StructLevel) {
		req_ := sl.Current().Interface().(req.CreateActReq)
		if req_.LabelForm.HolderType == HolderTypePersonal {
			if len(req_.LabelForm.Signer) <= 4 {
				sl.ReportError(req_.LabelForm.Signer, "Signer", "signer", "min_signer", "5")
				return
			}

			for _, s := range req_.LabelForm.Signer {
				if len(s.StudentID) != 10 {
					sl.ReportError(req_.LabelForm.Signer, "StudentID", "studentid", "len", "10")
					return
				}
			}
		}
	}, req.CreateActReq{})
}
