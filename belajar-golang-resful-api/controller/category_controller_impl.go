package controller

import (
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/model/domain/web"
	"belajar-golang-resful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryControllerImpl {
	return CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequest(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.CreateCategory(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		DATA:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}
func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}

	helper.ReadFromRequest(r, &categoryUpdateRequest)

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.UpdateCategory(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		DATA:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.DeleteCategory(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *CategoryControllerImpl) FindbyId(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		DATA:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		DATA:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
