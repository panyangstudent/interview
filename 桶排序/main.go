package main
import "fmt"
//实现排序排序
// 根据最大值和最小值，计算出需要的桶数，
// 再将各个元素入桶，通过自身值和最小值相减在和桶的大小取模
// 随后对每个桶进行排序， 这块的排序算法看自己，如果是冒泡，插入这种那整个桶排序就是稳定的排序
func quickSort(nums []int, start, end int) []int {  
	if start < end {    
		privote := nums[start]
		j := start 
		for i := start; i < end; i++ {
			if nums[i] < privote {
				nums[j] ,nums[i] = nums[i],nums[j]
			}
		}
		nums = quickSort(nums, start, j)      
		nums = quickSort(nums, j+1, end)    
	}  
	return nums
}

func bucketSort(nums []int, bucketSize int) []int {
	minValue := nums[0]
	maxValue := nums[0]
	// 获取到最大，最小的元素
	for i := 1; i <len(nums); i++ {
		if nums[i] < minValue {
			minValue = nums[i]
		}
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}

	bucketCount := make([][]int,(maxValue - minValue) / bucketSize +1)
  	//将数据分配到桶中  
  	for i := 0; i < len(nums); i++ {    
		bucketCount[(nums[i]-minValue)/bucketSize] = append(bucketCount[(nums[i]-minValue)/bucketSize], nums[i])  
	}
  	//循环所有的桶进行排序  
	newnums := make([]int,0)
	for i := 0; i < len(bucketCount); i++ {    
		if len(bucketCount[i]) > 1 {      
			//对每个桶内部进行排序,使用快排      
			bucketCount[i] = quickSort(bucketCount[i], 0, len(bucketCount[i])-1)      
		}
		newnums = append(newnums,bucketCount[i]...)
	}
	return newnums
}
func main() {
	nums := []int{45, 63, 3, 1, 29, 77, 20, 4, 30,120,233,3333,22222,222}  
	fmt.Println("排序前：")  
	fmt.Println(nums)  
	nums = bucketSort(nums, 20)  
	fmt.Println("排序后：")  
	fmt.Println(nums)
}