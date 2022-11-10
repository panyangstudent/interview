package main


func subsets(nums []int) (ans [][]int) {
    set := []int{}
    var dfs func(int)
    dfs = func(cur int) {
        if cur == len(nums) {
            ans = append(ans, append([]int(nil), set...))
            return
        }
        set = append(set, nums[cur])
        dfs(cur + 1)
        set = set[:len(set)-1]
        dfs(cur + 1)
    }
    dfs(0)
    return
}

func subStr(nums []int) (ans [][]int)  {
    set := make([]int,0)
    var def func(int)
    def = func(cur int) {
        if cur == len(nums) {
            ans = append(ans, append([]int(nil),set...))
            return
        }
        //要当前位置
        set = append(set ,nums[cur])
        def(cur+1)

        // 不要当前位置
        set = set[: len(set)-1]
        def(cur+1)
    }
    def(0)
    return ans
}