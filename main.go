package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 轮询时间间隔，可根据需要自定义
	pollInterval := 10 * time.Second
	// 无限循环执行轮询任务
	for {
		// 发起 HTTP GET 请求获取酶标仪数据
		measurements, err := getMeasurements()
		if err != nil {
			fmt.Printf("Error retrieving measurements: %v\n", err)
			continue
		}
		// 计算平均值
		average := calculateAverage(measurements)
		// 打印平均值
		fmt.Printf("Average measurement: %.2f\n", average)
		// 等待下一次轮询
		time.Sleep(pollInterval)
	}
}

// 发起 HTTP GET 请求获取酶标仪数据
func getMeasurements() ([]float64, error) {
	// 发起 HTTP GET 请求
	resp, err := http.Get("http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement")
	if err != nil {
		return nil, fmt.Errorf("HTTP GET request failed: %v", err)
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}
	// 将响应数据解析为浮点数数组
	measurements, err := parseMeasurements(string(body))
	if err != nil {
		return nil, fmt.Errorf("Error parsing measurements: %v", err)
	}
	return measurements, nil
}

// 将响应数据解析为浮点数数组
func parseMeasurements(data string) ([]float64, error) {
	// 按空格拆分数据
	values := strings.Fields(data)
	// 将字符串转换为浮点数
	var measurements []float64
	for _, val := range values {
		measurement, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, fmt.Errorf("Error parsing measurement value: %v", err)
		}
		measurements = append(measurements, measurement)
	}
	return measurements, nil
}

// 计算测量值数组的平均值
func calculateAverage(measurements []float64) float64 {
	var sum float64
	for _, val := range measurements {
		sum += val
	}
	return sum / float64(len(measurements))
}
