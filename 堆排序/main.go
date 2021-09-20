// 思路堆排序实际是数组排序，使用数组下标构造一个二叉树，
// 加入有数组a，那可以将a[0]视为根节点，它的左子节点为a[2*i + 1]，右子节点为a[2*i +2]
// 构建一个大顶堆，构建成功后a[0]便是最大的
// 之后交换a[0]和a[len(a) - 1]
// 然后在调整a[:len(a)-2]为大顶堆
// 重复上面的步骤，得到有序数组
// 时间复杂度为nlog(n)
 

// func adjustHeap(a []int64，pos int64，length int64) {
// 	for {
// 		// 计算孩子位置, 如果超过则break
// 		child := pos * 2 + 1
// 		if child >= length - 1 {
// 			break
// 		} 
// 		// 找出孩子中较大的那个
// 		if a[child+1] > a[child] {
// 			child ++ 
// 		}
// 		// 检查孩子是否大于父亲，如果大于则交换
// 		if a[pos] < a[child] {
// 		   a[pos],a[child] = a[child],a[pos]
// 		   // 更新位置，指向子节点，把子节点的个个孩子都构建下， 这个pos只能不断在增大
// 		   pos = child
// 		} else {
// 			// 如果子节点小于父节点，则停止调整
// 			break
// 		}
// 	}
// }
// func buildHeap(a []int64){
// 	// 从底层向上构建，len(a)/2-1是第一个非叶子节点， 这一步是为了构建一个堆
// 	for i := len(a)/2-1; i>=0 ; i-- {
// 		adjustHeap(a, i, len(a))
// 	}

// 	// 首尾交换，重新调整构建大顶堆，
// 	for i := len(a) - 1 ; i >= 0; i -- {
// 		a[0], a[i] = a[i], a[0]
// 		adjustHeap(a, 0, i)
// 	}
// }
package main
import (
    "fmt"
)
func main(){
	num := []int64{98, 48, 777, 63, 57, 433, 23, 1112, 1}
	heapSort(num)
	
	fmt.Println("heap sort over:", num)

}



// func heapSort (arr []int64) {
// 	// 从第一个不是叶子节点开始构建堆
// 	for i := int64(len(arr))/2 - 1; i>=0; i-- {
// 		// 调整
// 		adjustHeap(arr, i, int64(len(arr)))
// 	}

// 	// 从第一个节点开始构建
// 	for i := int64(len(arr)) - 1; i>=0 ; i-- {
// 		arr[0], arr[i] = arr[i], arr[0]
// 		adjustHeap(arr, 0 ,i)
// 	}
// }
// func adjustHeap(arr []int64,pos int64, length int64) {
// 	for {
// 		child := pos * 2 + 1
// 		if child >= length - 1 {
// 			break
// 		}

// 		if arr[child + 1] > arr[child] {
// 			child ++ 
// 		}

// 		if arr[pos] < arr[child] {
// 			arr[pos], arr[child] = arr[child], arr[pos]
// 			pos = child
// 		} else {
// 			break
// 		}
// 	}
// }

func heapSort(arr []int64) {
	//构建
	for i := int64(len(arr))/2-1;i>=0;i-- {
		adjectHeap(arr, i, int64(len(arr)))
	}
	// 调整
	for i := int64(len(arr)) -1 ; i >= 0 ; i-- {
		arr[0],arr[i] = arr[i],arr[0]
		adjectHeap(arr, 0, i)
	}
}

func adjectHeap(arr []int64, pos int64, length int64) {
	for {
		child := pos * 2 + 1

		if child >= length-1 {
			break
		} 
			
		if arr[child] < arr[child+1] {
			child++
		}

		if arr[pos] < arr[child] {
			arr[pos], arr[child] = arr[child], arr[pos]
			pos = child
		} else {
			break
		}

	}
}