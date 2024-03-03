package piscine

func TwoSum(nums []int, target int) []int {
	tab:=[]int{}
	result:=[]int{}
	for i:=0 ; i<len(nums)-1 ; i++{
		for j:=i+1 ; j<len(nums) ; j++{
			if nums[i] + nums[j] == target {
				tab=append(tab,i)
				tab=append(tab,j)
				break
			}

		}
	}
	if len(tab)>=2 {
		result=append(result,tab[0])
		result=append(result, tab[1])
	}else if len(tab)==0 {
		result=nil
	}
	return result
}