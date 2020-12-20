/**
2 * @Author: Nico
3 * @Date: 2020/12/21 0:49
4 */
package proxy

import "github.com/ainilili/stock/model"

type Proxy interface {
	List(query string) ([]model.Stock, error)
	Get(query string)(*model.StockDetails, error)
}

var proxy = SinaProxy{}

func List(query string) ([]model.Stock, error){
	return proxy.List(query)
}

func Get(query string)(*model.StockDetails, error){
	return proxy.Get(query)
}
