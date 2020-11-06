/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/11/5 17:07
 */
package array

func InStrings(s string, array []string) bool {
	for _, v := range array {
		if s == v {
			return true
		}
	}
	return false
}
