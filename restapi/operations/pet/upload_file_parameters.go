package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadFileParams creates a new UploadFileParams object
// with the default values initialized.
func NewUploadFileParams() UploadFileParams {
	var ()
	return UploadFileParams{}
}

// UploadFileParams contains all the bound params for the upload file operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadFile
type UploadFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Additional data to pass to server
	  In: formData
	*/
	AdditionalMetadata *string
	/*file to upload
	  In: formData
	*/
	File runtime.File
	/*ID of pet to update
	  Required: true
	  In: path
	*/
	PetID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UploadFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdAdditionalMetadata, fdhkAdditionalMetadata, _ := fds.GetOK("additionalMetadata")
	if err := o.bindAdditionalMetadata(fdAdditionalMetadata, fdhkAdditionalMetadata, route.Formats); err != nil {
		res = append(res, err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else {
		o.File = runtime.File{Data: file, Header: fileHeader}
	}

	rPetID, rhkPetID, _ := route.Params.GetOK("petId")
	if err := o.bindPetID(rPetID, rhkPetID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadFileParams) bindAdditionalMetadata(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.AdditionalMetadata = &raw

	return nil
}

func (o *UploadFileParams) bindPetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("petId", "path", "int64", raw)
	}
	o.PetID = value

	return nil
}
