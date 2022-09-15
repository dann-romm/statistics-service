package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"statistics-service/internal/entity"
	"statistics-service/internal/service"
	"strconv"
)

type StatisticsRoutes struct {
	s service.Statistics
}

func newStatisticsRoutes(g *echo.Group, s service.Statistics) {
	r := &StatisticsRoutes{s: s}

	g.Group("/statistics")
	{
		g.POST("/save", r.saveData)
		g.GET("/get", r.getData)
		g.GET("/get/:record_id", r.getDataById)
		g.PUT("/update", r.updateData)
		g.DELETE("/delete", r.deleteData)
	}
}

func (r *StatisticsRoutes) saveData(c echo.Context) error {
	userId, ok := c.Get(userIdCtx).(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return nil
	}

	var input entity.StatisticalData
	err := c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return nil
	}

	recordId, err := r.s.SaveData(c.Request().Context(), userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"userId": recordId,
	})
}

func (r *StatisticsRoutes) getData(c echo.Context) error {
	userId, ok := c.Get(userIdCtx).(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return nil
	}

	records, err := r.s.GetData(c.Request().Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"entity": records,
	})
}

func (r *StatisticsRoutes) getDataById(c echo.Context) error {
	userId, ok := c.Get(userIdCtx).(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return nil
	}

	recordIdParam := c.Param("record_id")
	recordId, err := strconv.Atoi(recordIdParam)
	if recordIdParam == "" || err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return nil
	}

	record, err := r.s.GetDataById(c.Request().Context(), userId, recordId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"entity": record,
	})
}

func (r *StatisticsRoutes) updateData(c echo.Context) error {
	userId, ok := c.Get(userIdCtx).(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return nil
	}

	var input entity.StatisticalData
	err := c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return nil
	}

	recordIdParam := c.QueryParam("record_id")
	if recordIdParam == "" {
		newErrorResponse(c, http.StatusBadRequest, "record_id is required")
		return nil
	}

	recordId, err := strconv.Atoi(recordIdParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return nil
	}

	err = r.s.UpdateData(c.Request().Context(), userId, recordId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (r *StatisticsRoutes) deleteData(c echo.Context) error {
	userId, ok := c.Get(userIdCtx).(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return nil
	}

	recordIdParam := c.QueryParam("record_id")
	if recordIdParam == "" {
		newErrorResponse(c, http.StatusBadRequest, "record_id is required")
		return nil
	}

	recordId, err := strconv.Atoi(recordIdParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return nil
	}

	err = r.s.DeleteData(c.Request().Context(), userId, recordId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
