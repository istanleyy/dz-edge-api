package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/models"
	"thingularity.co/dz-edge-api/services"
)

type preprocData struct {
	Process       string  `json:"process"`
	WoID          string  `json:"woID"`
	EmptyWeight   float32 `json:"emptyWeight"`
	FullWeight    float32 `json:"fullWeight"`
	SkimTime1     float32 `json:"skimTime1"`
	SkimTime2     float32 `json:"skimTime2"`
	SkimTime3     float32 `json:"skimTime3"`
	PicklingTime1 float32 `json:"picklingTime1"`
	PicklingTime2 float32 `json:"picklingTime2"`
	PicklingTime3 float32 `json:"picklingTime3"`
	FluxTime1     float32 `json:"fluxTime1"`
	GoodsWeight   float32 `json:"goodsWeight"`
}

type galData struct {
	Process           string    `json:"process"`
	WoID              string    `json:"woID"`
	FinishEmptyWeight float32   `json:"finishEmptyWeight"`
	FinishFullWeight  float32   `json:"finishFullWeight"`
	GoodsWeight       float32   `json:"goodsWeight"`
	LittleBlueWeight  []float32 `json:"littleBlueWeight"`
	LittleBlueTime    []float32 `json:"littleBlueTime"`
	LittleBlueTemp    []float32 `json:"littleBlueTemp"`
}

// GetJob ...
func GetJob(c *gin.Context) {
	jobID := c.Param("jobId")         // matching uri definition
	process := c.Param("processType") // matching uri definition

	switch process {
	case "pre":
		i := models.JobExists("pre", jobID)
		if i != -1 {
			pData, err := getPreData(i, jobID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, pData)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot find job in pre-treatment process"})
		return
	case "gal":
		i := models.JobExists("gal", jobID)
		if i != -1 {
			gData, err := getGalData(i, jobID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gData)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot find job in galvanizing process"})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown process. Process must be 'pre' (pre-treatment) or 'gal' (galvanized)."})
	}
}

func getPreData(i int, job string) (*preprocData, error) {
	if i < 0 || i >= 20 {
		return nil, errors.New("invalid index in getPreData")
	}
	valTagStr := models.GetJobMetrics(i)
	val, err := services.GetTagValue("dotzero", valTagStr)
	if err != nil {
		return nil, err
	}
	return &preprocData{"Pre-treatment", job, float32(val[0]), float32(val[1]), float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6]), float32(val[7]), float32(val[8]), float32(val[14])}, nil
	/*
		switch i {
		case 0:
			valTagStr := []string{models.W01WtBktEmpt, models.W01WtBktFull, models.W01TimeDegrease1, models.W01TimeDegrease2, models.W01TimeDegrease3, models.W01TimePickling1, models.W01TimePickling2, models.W01TimePickling3, models.W01TimeFlux, models.W01WtGood}
			val, err := services.GetTagValue("dotzero", valTagStr)
			if err != nil {
				return nil, err
			}
			return &preprocData{"Pre-treatment", job, float32(val[0]), float32(val[1]), float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6]), float32(val[7]), float32(val[8]), float32(val[9])}, nil
		case 1:
			valTagStr := []string{models.W02WtBktEmpt, models.W02WtBktFull, models.W02TimeDegrease1, models.W02TimeDegrease2, models.W02TimeDegrease3, models.W02TimePickling1, models.W02TimePickling2, models.W02TimePickling3, models.W02TimeFlux, models.W02WtGood}
			val, err := services.GetTagValue("dotzero", valTagStr)
			if err != nil {
				return nil, err
			}
			return &preprocData{"Pre-treatment", job, float32(val[0]), float32(val[1]), float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6]), float32(val[7]), float32(val[8]), float32(val[9])}, nil
		case 2:
			valTagStr := []string{models.W03WtBktEmpt, models.W03WtBktFull, models.W03TimeDegrease1, models.W03TimeDegrease2, models.W03TimeDegrease3, models.W03TimePickling1, models.W03TimePickling2, models.W03TimePickling3, models.W03TimeFlux, models.W03WtGood}
			val, err := services.GetTagValue("dotzero", valTagStr)
			if err != nil {
				return nil, err
			}
			return &preprocData{"Pre-treatment", job, float32(val[0]), float32(val[1]), float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6]), float32(val[7]), float32(val[8]), float32(val[9])}, nil
		case 3:
			valTagStr := []string{models.W04WtBktEmpt, models.W04WtBktFull, models.W04TimeDegrease1, models.W04TimeDegrease2, models.W04TimeDegrease3, models.W04TimePickling1, models.W04TimePickling2, models.W04TimePickling3, models.W04TimeFlux, models.W04WtGood}
			val, err := services.GetTagValue("dotzero", valTagStr)
			if err != nil {
				return nil, err
			}
			return &preprocData{"Pre-treatment", job, float32(val[0]), float32(val[1]), float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6]), float32(val[7]), float32(val[8]), float32(val[9])}, nil
		case 4:
			valTagStr := []string{models.W05WtBktEmpt, models.W05WtBktFull, models.W05TimeDegrease1, models.W05TimeDegrease2, models.W05TimeDegrease3, models.W05TimePickling1, models.W05TimePickling2, models.W05TimePickling3, models.W05TimeFlux, models.W05WtGood}
			val, err := services.GetTagValue("dotzero", valTagStr)
			if err != nil {
				return nil, err
			}
			return &preprocData{"Pre-treatment", job, float32(val[0]), float32(val[1]), float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6]), float32(val[7]), float32(val[8]), float32(val[9])}, nil
		}
	*/
}

func getGalData(i int, job string) (*galData, error) {
	switch i {
	case 0:
		tagArray := []string{models.G01WtBktEmpt, models.G01WtBktFull, models.G01WtGood}
		wtLilBlu := models.GetTagArray(models.G01WtLilblu, 5)
		timeLilBlu := models.GetTagArray(models.G01TimeLilblu, 5)
		tempLilBlu := models.GetTagArray(models.G01TempLilblu, 5)
		tags1 := append(tagArray, wtLilBlu...)
		tags2 := append(timeLilBlu, tempLilBlu...)
		tags := append(tags1, tags2...)

		val, err := services.GetTagValue("dotzero", tags)
		if err != nil {
			return nil, err
		}
		wtLilBluVal := []float32{float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6])}
		timeLilBluVal := []float32{float32(val[7]), float32(val[8]), float32(val[9]), float32(val[10]), float32(val[11])}
		tempLilBluVal := []float32{float32(val[12]), float32(val[13]), float32(val[14]), float32(val[15]), float32(val[16])}
		return &galData{"Galvanized", job, float32(val[0]), float32(val[1]), float32(val[2]), wtLilBluVal, timeLilBluVal, tempLilBluVal}, nil
	case 1:
		tagArray := []string{models.G02WtBktEmpt, models.G02WtBktFull, models.G02WtGood}
		wtLilBlu := models.GetTagArray(models.G02WtLilblu, 5)
		timeLilBlu := models.GetTagArray(models.G02TimeLilblu, 5)
		tempLilBlu := models.GetTagArray(models.G02TempLilblu, 5)
		tags1 := append(tagArray, wtLilBlu...)
		tags2 := append(timeLilBlu, tempLilBlu...)
		tags := append(tags1, tags2...)

		val, err := services.GetTagValue("dotzero", tags)
		if err != nil {
			return nil, err
		}
		wtLilBluVal := []float32{float32(val[2]), float32(val[3]), float32(val[4]), float32(val[5]), float32(val[6])}
		timeLilBluVal := []float32{float32(val[7]), float32(val[8]), float32(val[9]), float32(val[10]), float32(val[11])}
		tempLilBluVal := []float32{float32(val[12]), float32(val[13]), float32(val[14]), float32(val[15]), float32(val[16])}
		return &galData{"Galvanized", job, float32(val[0]), float32(val[1]), float32(val[2]), wtLilBluVal, timeLilBluVal, tempLilBluVal}, nil
	}
	return nil, errors.New("invalid index in getGalData")
}
