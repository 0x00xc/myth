/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/26 15:10
 */
package cache

type Storage interface {
	Put(key string, v []byte, exp  ...int64) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}
