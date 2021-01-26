package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/models"
	"thingularity.co/dz-edge-api/services"
)

// UpdateRoute ...
func UpdateRoute(c *gin.Context) {
	jobID := c.Param("jobId")                    // matching uri definition
	route, err := strconv.Atoi(c.Param("route")) // matching uri definition
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("JobID: " + jobID + "  Route: " + strconv.Itoa(route))

	// Routes define the processes. Each process has its own metric fields to update.
	switch route {
	case 1:
		// index, err := models.UpdateJob("pre", jobID, 1)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
		// 	return
		// }
		// if index < 0 {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update pre-treatment job index out of bound"})
		// 	return
		// }
		// Get metric value from SCADA system
		val, err := services.GetTagValue("dotzero", []string{models.WScale})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		metrics := services.Route01Metrics{
			EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', 2, 64),
		}
		jsonstr, err := json.Marshal(metrics)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		// Update job metrics with iShopFloor API
		if err = services.UpdateRouteMetrics(jobID, "1", string(jsonstr)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": "n/a"})
		return

	case 2:
		// index, err := models.UpdateJob("pre", jobID, 1)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
		// 	return
		// }
		// if index < 0 {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update pre-treatment job index out of bound"})
		// 	return
		// }
		// Get metric value from SCADA system
		val, err := services.GetTagValue("dotzero", []string{models.WScale})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		metrics := services.Route02Metrics{
			FullBucketWeight: strconv.FormatFloat(val[0], 'f', 2, 64),
		}
		jsonstr, err := json.Marshal(metrics)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		// Update job metrics with iShopFloor API
		if err = services.UpdateRouteMetrics(jobID, "2", string(jsonstr)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": "n/a"})
		return

	case 3:
		index, err := models.UpdateJob("pre", jobID, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		if index < 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update pre-treatment job index out of bound"})
			return
		}
		// Get metric value from SCADA system
		val, err := services.GetTagValue("dotzero", models.GetJobMetrics(index))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		metrics := services.Route03Metrics{
			PicklingSmallTime:   strconv.FormatFloat(val[0], 'f', 2, 64) + "," + strconv.FormatFloat(val[1], 'f', 2, 64) + "," + strconv.FormatFloat(val[2], 'f', 2, 64),
			DegreasingSmallTime: strconv.FormatFloat(val[3], 'f', 2, 64) + "," + strconv.FormatFloat(val[4], 'f', 2, 64) + "," + strconv.FormatFloat(val[5], 'f', 2, 64),
			HotWaterTime:        strconv.FormatFloat(val[6], 'f', 2, 64) + "," + strconv.FormatFloat(val[7], 'f', 2, 64) + "," + strconv.FormatFloat(val[8], 'f', 2, 64),
			DegreasingTime:      strconv.FormatFloat(val[9], 'f', 2, 64) + "," + strconv.FormatFloat(val[10], 'f', 2, 64) + "," + strconv.FormatFloat(val[11], 'f', 2, 64),
			DegreasingWaterTime: strconv.FormatFloat(val[12], 'f', 2, 64) + "," + strconv.FormatFloat(val[13], 'f', 2, 64) + "," + strconv.FormatFloat(val[14], 'f', 2, 64),
			FluxTime:            strconv.FormatFloat(val[15], 'f', 2, 64) + "," + strconv.FormatFloat(val[16], 'f', 2, 64) + "," + strconv.FormatFloat(val[17], 'f', 2, 64),
			DryingTime:          strconv.FormatFloat(val[18], 'f', 2, 64) + "," + strconv.FormatFloat(val[19], 'f', 2, 64) + "," + strconv.FormatFloat(val[20], 'f', 2, 64),
			PicklingTime:        strconv.FormatFloat(val[21], 'f', 2, 64) + "," + strconv.FormatFloat(val[22], 'f', 2, 64) + "," + strconv.FormatFloat(val[23], 'f', 2, 64) + "," + strconv.FormatFloat(val[24], 'f', 2, 64) + "," + strconv.FormatFloat(val[25], 'f', 2, 64) + "," + strconv.FormatFloat(val[26], 'f', 2, 64),
			PicklingWaterTime:   strconv.FormatFloat(val[27], 'f', 2, 64) + "," + strconv.FormatFloat(val[28], 'f', 2, 64) + "," + strconv.FormatFloat(val[29], 'f', 2, 64),
			LessFluxTime:        strconv.FormatFloat(val[30], 'f', 2, 64) + "," + strconv.FormatFloat(val[31], 'f', 2, 64) + "," + strconv.FormatFloat(val[31], 'f', 2, 64),
			TurnOverTime:        strconv.FormatFloat(val[33], 'f', 2, 64) + "," + strconv.FormatFloat(val[34], 'f', 2, 64) + "," + strconv.FormatFloat(val[35], 'f', 2, 64),
		}
		jsonstr, err := json.Marshal(metrics)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		// Update job metrics with iShopFloor API
		if err = services.UpdateRouteMetrics(jobID, "3", string(jsonstr)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": index})
		return

	case 4:
		index, err := models.UpdateJob("gal", jobID, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		switch index {
		case 0:
			// Get metric value from SCADA system
			tags := []string{}
			wtLilBlu := models.GetTagArray(models.G01WtLilblu, 5)
			timeLilBlu := models.GetTagArray(models.G01TimeLilblu, 5)
			tempLilBlu := models.GetTagArray(models.G01TempLilblu, 5)
			tags = append(tags, wtLilBlu...)
			tagss := append(timeLilBlu, tempLilBlu...)
			tags = append(tags, tagss...)
			val, err := services.GetTagValue("dotzero", tags)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			metrics := services.Route04Metrics{
				SubBasketWeight: strconv.FormatFloat(val[0], 'f', 2, 64) + "," + strconv.FormatFloat(val[1], 'f', 2, 64) + "," + strconv.FormatFloat(val[2], 'f', 2, 64) + "," + strconv.FormatFloat(val[3], 'f', 2, 64) + "," + strconv.FormatFloat(val[4], 'f', 2, 64),
				SubBasketTime:   strconv.FormatFloat(val[5], 'f', 2, 64) + "," + strconv.FormatFloat(val[6], 'f', 2, 64) + "," + strconv.FormatFloat(val[7], 'f', 2, 64) + "," + strconv.FormatFloat(val[8], 'f', 2, 64) + "," + strconv.FormatFloat(val[9], 'f', 2, 64),
				SubBasketTemp:   strconv.FormatFloat(val[10], 'f', 2, 64) + "," + strconv.FormatFloat(val[11], 'f', 2, 64) + "," + strconv.FormatFloat(val[12], 'f', 2, 64) + "," + strconv.FormatFloat(val[13], 'f', 2, 64) + "," + strconv.FormatFloat(val[14], 'f', 2, 64),
			}
			jsonstr, err := json.Marshal(metrics)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			// Update job metrics with iShopFloor API
			if err = services.UpdateRouteMetrics(jobID, "4", string(jsonstr)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": index})
			return
		case 1:
			// Get metric value from SCADA system
			tags := []string{}
			wtLilBlu := models.GetTagArray(models.G02WtLilblu, 5)
			timeLilBlu := models.GetTagArray(models.G02TimeLilblu, 5)
			tempLilBlu := models.GetTagArray(models.G02TempLilblu, 5)
			tags = append(tags, wtLilBlu...)
			tagss := append(timeLilBlu, tempLilBlu...)
			tags = append(tags, tagss...)
			val, err := services.GetTagValue("dotzero", tags)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			metrics := services.Route04Metrics{
				SubBasketWeight: strconv.FormatFloat(val[0], 'f', 2, 64) + "," + strconv.FormatFloat(val[1], 'f', 2, 64) + "," + strconv.FormatFloat(val[2], 'f', 2, 64) + "," + strconv.FormatFloat(val[3], 'f', 2, 64) + "," + strconv.FormatFloat(val[4], 'f', 2, 64),
				SubBasketTime:   strconv.FormatFloat(val[5], 'f', 2, 64) + "," + strconv.FormatFloat(val[6], 'f', 2, 64) + "," + strconv.FormatFloat(val[7], 'f', 2, 64) + "," + strconv.FormatFloat(val[8], 'f', 2, 64) + "," + strconv.FormatFloat(val[9], 'f', 2, 64),
				SubBasketTemp:   strconv.FormatFloat(val[10], 'f', 2, 64) + "," + strconv.FormatFloat(val[11], 'f', 2, 64) + "," + strconv.FormatFloat(val[12], 'f', 2, 64) + "," + strconv.FormatFloat(val[13], 'f', 2, 64) + "," + strconv.FormatFloat(val[14], 'f', 2, 64),
			}
			jsonstr, err := json.Marshal(metrics)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			// Update job metrics with iShopFloor API
			if err = services.UpdateRouteMetrics(jobID, "4", string(jsonstr)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": index})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u001", "message": "update galvanizing job index out of bound"})
			return
		}

	case 5:
		// index, err := models.UpdateJob("gal", jobID, 1)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
		// 	return
		// }
		// if index < 0 {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update galvanizing job index out of bound"})
		// 	return
		// }
		// Get metric value from SCADA system
		val, err := services.GetTagValue("dotzero", []string{models.GScale})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		metrics := services.Route05Metrics{
			EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', 2, 64),
		}
		jsonstr, err := json.Marshal(metrics)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		// Update job metrics with iShopFloor API
		if err = services.UpdateRouteMetrics(jobID, "5", string(jsonstr)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": "n/a"})
		return

	case 6:
		// index, err := models.UpdateJob("gal", jobID, 1)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
		// 	return
		// }
		// if index < 0 {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update galvanizing job index out of bound"})
		// 	return
		// }
		// Get metric value from SCADA system
		val, err := services.GetTagValue("dotzero", []string{models.GScale})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		metrics := services.Route06Metrics{
			FullBucketWeight: strconv.FormatFloat(val[0], 'f', 2, 64),
		}
		jsonstr, err := json.Marshal(metrics)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		// Update job metrics with iShopFloor API
		if err = services.UpdateRouteMetrics(jobID, "6", string(jsonstr)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": "n/a"})
		return

	case 7:
		// index, err := models.UpdateJob("gal", jobID, 1)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
		// 	return
		// }
		// if index < 0 {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update galvanizing job index out of bound"})
		// 	return
		// }
		// Get metric value from SCADA system
		val, err := services.GetTagValue("dotzero", []string{models.GScale})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		metrics := services.Route07Metrics{
			ProductWeight: strconv.FormatFloat(val[0], 'f', 2, 64),
		}
		jsonstr, err := json.Marshal(metrics)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		// Update job metrics with iShopFloor API
		if err = services.UpdateRouteMetrics(jobID, "7", string(jsonstr)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"woID": jobID, "queueIndex": "n/a"})
		return

	default:
		c.JSON(http.StatusBadRequest, gin.H{"code": "u001", "message": "Unknown route. Route number must be between 1~7."})
	}
}