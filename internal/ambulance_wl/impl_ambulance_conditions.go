package ambulance_wl

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_ambulance_conditions.go
func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		result := ambulance.PredefinedConditions
		if result == nil {
			result = []Condition{}
		}
		return nil, result, http.StatusOK
	})
}

// Copy following section to separate file, uncomment, and implement accordingly
// CreateCondition - Saves new condition into ambulance
func (this *implAmbulanceConditionsAPI) CreateCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		var condition Condition
		if err := ctx.ShouldBindJSON(&condition); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}
		conflictIndx := slices.IndexFunc(ambulance.PredefinedConditions, func(c Condition) bool {
			return c.Value == condition.Value
		})
		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Condition already exists",
			}, http.StatusConflict
		}
		ambulance.PredefinedConditions = append(ambulance.PredefinedConditions, condition)
		return ambulance, condition, http.StatusCreated

	})

}

// DeleteCondition - Deletes specific condition
func (this *implAmbulanceConditionsAPI) DeleteCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		// TODO: UPDATE these functions to use correct endpoint and update by code in the URL
		conditionId := ctx.Param("conditionId")
		indx := slices.IndexFunc(ambulance.PredefinedConditions, func(c Condition) bool {
			return c.Code == conditionId
		})
		if indx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Condition not found",
			}, http.StatusNotFound
		}
		ambulance.PredefinedConditions = append(ambulance.PredefinedConditions[:indx], ambulance.PredefinedConditions[indx+1:]...)
		return ambulance, nil, http.StatusOK
	})
}

// UpdateCondition - Updates specific condition
func (this *implAmbulanceConditionsAPI) UpdateCondition(ctx *gin.Context) {
	// TODO: UPDATE these functions to use correct endpoint and update by code in the URL

	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		var condition Condition
		if err := ctx.ShouldBindJSON(&condition); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}
		indx := slices.IndexFunc(ambulance.PredefinedConditions, func(c Condition) bool {
			return c.Code == condition.Code
		})
		if indx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Condition not found",
			}, http.StatusNotFound
		}
		ambulance.PredefinedConditions[indx] = condition
		return ambulance, condition, http.StatusOK
	})
}

//
