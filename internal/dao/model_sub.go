/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/11/6 14:46
 */
package dao

type ArticleSubscribe struct {
	Model
	UID       int
	ArticleID int
	TradeID   string
}
