package main

import (
	"fmt"

	"github.com/sunmi-OS/gocore/algorithm/magicalbayes/brain"
	"github.com/sunmi-OS/gocore/algorithm/magicalbayes/classifier"
)

func main() {
	test0()
}

func test0() {
	bayesBrain := brain.NewBayesBrain()
	bayesClassifier := classifier.BayesClassifier{Brain: bayesBrain}
	//训练样本
	bayesBrain.Learn("Chinese", "Chinese", "Beijing", "Chinese")
	bayesBrain.Learn("Chinese", "Chinese", "Chinese", "Shanghai")
	bayesBrain.Learn("Chinese", "Chinese", "Macao")
	bayesBrain.Learn("Not Chinese", "Tokyo", "Japan", "Chinese")
	//应用tf-idf算法对特征加权
	bayesBrain.ApplyTfIdf()
	bayesBrain.Show()

	//测试样本
	features := []string{"Chinese", "Chinese", "Chinese", "Tokyo", "Japan"}

	//计算属于Chinese的概率
	fmt.Println(bayesClassifier.BayesProbabilityOf("Chinese",
		features...))
	//计算属于Not Chinese的概率
	fmt.Println(bayesClassifier.BayesProbabilityOf("Not Chinese",
		features...))

	//对测试样本进行分类
	fmt.Println(bayesClassifier.Classify(features...))
}
