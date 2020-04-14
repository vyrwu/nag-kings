/*
 * Kings API
 *
 * This is a sample Kings Server based on the OpenAPI 3.0 specification.
 *
 * API version: 0.1.0
 * Contact: dev.anowak@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var cachedResponse []byte

func requestData(url string) ([]byte, error) {
	if cachedResponse != nil {
		return cachedResponse, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}
	return body, nil
}

func GetKings(w http.ResponseWriter, r *http.Request) {

	body, err := requestData("https://gist.githubusercontent.com/christianpanton/10d65ccef9f29de3acd49d97ed423736/raw/b09563bc0c4b318132c7a738e679d4f984ef0048/kings")
	if err != nil {
		log.Fatalln(err)
	}

	var kingsRawArray []KingRaw

	json.Unmarshal(body, &kingsRawArray)

	var kingsArray []King

	for _, kingRaw := range kingsRawArray {
		king := King{
			ID:    kingRaw.ID,
			Name:  kingRaw.Nm,
			City:  kingRaw.Cty,
			House: kingRaw.Hse,
			Years: kingRaw.Yrs,
		}
		kingsArray = append(kingsArray, king)
	}

	response, err := json.Marshal(kingsArray)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func GetKingsStatistics(w http.ResponseWriter, r *http.Request) {

	body, err := requestData("https://gist.githubusercontent.com/christianpanton/10d65ccef9f29de3acd49d97ed423736/raw/b09563bc0c4b318132c7a738e679d4f984ef0048/kings")
	if err != nil {
		log.Fatalln(err)
	}

	var kingsRawArray []KingRaw

	json.Unmarshal(body, &kingsRawArray)

	kingByYearsRuled := make(map[King]int)
	houseByYearsRuled := make(map[string]int)
	kingsFirstNameCounter := make(map[string]int)
	for _, kingRaw := range kingsRawArray {
		king := King{
			ID:    kingRaw.ID,
			Name:  kingRaw.Nm,
			City:  kingRaw.Cty,
			House: kingRaw.Hse,
			Years: kingRaw.Yrs,
		}

		yearsRuled, err := calculateYearsRuledFromString(king.Years)
		if err != nil {
			panic(err)
		}

		currentKingRulingYears := kingByYearsRuled[king]
		kingByYearsRuled[king] = currentKingRulingYears + yearsRuled

		currentHouseRulingYears := houseByYearsRuled[king.House]
		houseByYearsRuled[king.House] = currentHouseRulingYears + yearsRuled

		nm := kingRaw.Nm
		firstName := strings.Split(nm, " ")[0]
		if _, ok := kingsFirstNameCounter[firstName]; ok {
			kingsFirstNameCounter[firstName]++
		} else {
			kingsFirstNameCounter[firstName] = 1
		}
	}

	longestRulingKing := make(chan *KingsStatisticsLongestRulingKing)
	go func() {
		getLongestRulingKing := func(kingByYearsRuled map[King]int) *KingsStatisticsLongestRulingKing {
			var longestRulingKing KingsStatisticsLongestRulingKing
			for k, v := range kingByYearsRuled {
				if (longestRulingKing.YearsRuled == 0) || longestRulingKing.YearsRuled == v {
					longestRulingKing = KingsStatisticsLongestRulingKing{
						King:       append(longestRulingKing.King, k),
						YearsRuled: v,
					}
				} else {
					if v > longestRulingKing.YearsRuled {
						longestRulingKing = KingsStatisticsLongestRulingKing{
							King:       []King{k},
							YearsRuled: v,
						}
					}
				}
			}
			return &longestRulingKing
		}
		// function calls create new frames in stack. They are expensive, and should be used
		// only when neccessary
		longestRulingKing <- getLongestRulingKing(kingByYearsRuled)
		close(longestRulingKing)
	}()

	longestRulingHouse := make(chan *KingsStatisticsLongestRulingHouse)
	go func() {
		getLongestRulingHouse := func(houseByYearsRuled map[string]int) *KingsStatisticsLongestRulingHouse {
			var longestRulingHouse KingsStatisticsLongestRulingHouse
			for k, v := range houseByYearsRuled {
				if (longestRulingHouse.YearsRuled == 0) || longestRulingHouse.YearsRuled == v {
					longestRulingHouse = KingsStatisticsLongestRulingHouse{
						HouseName:  append(longestRulingHouse.HouseName, k),
						YearsRuled: v,
					}
				} else {
					if v > longestRulingHouse.YearsRuled {
						longestRulingHouse = KingsStatisticsLongestRulingHouse{
							HouseName:  []string{k},
							YearsRuled: v,
						}
					}
				}
			}
			return &longestRulingHouse
		}
		longestRulingHouse <- getLongestRulingHouse(houseByYearsRuled)
		close(longestRulingHouse)
	}()

	mostCommonFirstName := make(chan []string)
	go func() {
		getMostCommonFirstName := func(kingsFirstNameCounter map[string]int) []string {
			type firstNameCounter struct {
				FirstName []string
				Count     int
			}
			var counter firstNameCounter
			for k, v := range kingsFirstNameCounter {
				if (counter.Count == 0) || counter.Count == v {
					counter = firstNameCounter{
						FirstName: append(counter.FirstName, k),
						Count:     v,
					}
				}
				if v >= counter.Count {
					counter = firstNameCounter{
						FirstName: []string{k},
						Count:     v,
					}
				}
			}
			return counter.FirstName
		}
		mostCommonFirstName <- getMostCommonFirstName(kingsFirstNameCounter)
		close(mostCommonFirstName)
	}()

	statistics := KingsStatistics{
		TotalKings:          len(kingsRawArray),
		LongestRulingKing:   <-longestRulingKing,
		LongestRulingHouse:  <-longestRulingHouse,
		MostCommonFirstName: <-mostCommonFirstName,
	}

	response, err := json.Marshal(statistics)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func calculateYearsRuledFromString(yearsRuledString string) (int, error) {
	splitYears := strings.Split(yearsRuledString, "-")
	if len(splitYears) == 1 {
		return 1, nil
	}

	if len(splitYears) > 2 {
		return -1, errors.New("More than two years in range")
	}

	var splitYearsInt []int
	for _, year := range splitYears {
		i, err := strconv.Atoi(year)
		if err != nil {
			if year != "" {
				return -1, err
			}
			i = time.Now().Year()
		}
		splitYearsInt = append(splitYearsInt, i)
	}

	yearsRuled := splitYearsInt[1] - splitYearsInt[0]
	return yearsRuled, nil
}