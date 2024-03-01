package piscine

func PodiumPosition(podium [][]string) [][]string {
		for i:= 0 ; i<len(podium) ; i++{
			for j:=i+1 ; j<len(podium) ; j++{
				if podium[j][0]<podium[i][0]{
					podium[j],podium[i]=podium[i],podium[j]
				}
			}
		}
		return podium
}