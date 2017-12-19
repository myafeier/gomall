package controller

import (
	"gopkg.in/macaron.v1"
)

// /address/default/:id 设置默认地址
//  Success-Response:
//  HTTP/1.1 200 OK
//  {
//    "meta": {
// "code": 0,
// "message": "调用成功"
//    },
//    "data": {}
//  }
func AddressDefaultPostHandler(ctx *macaron.Context) {

}

// {get} /address/default 获取默认地址
// Response:
// {
//    "meta": {
// "code": 0,
// "message": "调用成功"
//    },
//    "data": {}
//  }
func AddressDefaultGetHandler(ctx *macaron.Context) {

}

// /address 列出所有资源
//    @apiParam {String} [page=1] 指定第几页
// 	  @apiParam {String} [limit=10] 指定每页的记录数
// 	  @apiParam {Boolean} [is_show] 指定is_show过滤

// 	  @apiPermission none
// 	  @apiSampleRequest /address

// 	  @apiUse Header
// 	  @apiUse Success

// 	  @apiSuccessExample Success-Response:
// 	      HTTP/1.1 200 OK
// 	      {
// 	        "meta": {
// 	        	"code": 0,
// 	        	"message": "调用成功"
// 	        },
// 	        "data": [{
// 	        	"_id": "_id",
// 	        	"images": [{
// 	        		"_id": "_id",
// 	        		"name": "name",
// 	        		"path": "path",
// 	        		"create_at": "create_at"
// 	        	}],
// 	        	"is_show": "is_show",
// 	        	"remark": "remark",
// 	        	"sort": "sort",
// 	        	"title": "title",
// 	        	"create_at": "create_at",
// 	        	"update_at": "update_at"
// 	        }]
// 	      }
func AddressListGetHandler(ctx *macaron.Context) {

}

//  {get} /address/:id 获取某个指定资源的信息
// @apiParam {String} id 资源ID
//   * @apiSuccessExample Success-Response:
// 	 *     HTTP/1.1 200 OK
// 	 *     {
// 	 *       "meta": {
// 	 *       	"code": 0,
// 	 *       	"message": "调用成功"
// 	 *       },
// 	 *       "data": {
// 	 *       	"_id": "_id",
// 	 *       	"images": [{
// 	 *       		"_id": "_id",
// 	 *       		"name": "name",
// 	 *       		"path": "path",
// 	 *       		"create_at": "create_at"
// 	 *       	}],
// 	 *       	"is_show": "is_show",
// 	 *       	"remark": "remark",
// 	 *       	"sort": "sort",
// 	 *       	"title": "title",
// 	 *       	"create_at": "create_at",
// 	 *       	"update_at": "update_at"
// 	 *       }
// 	 *     }
// 	 */
func AddressOneGetHandler(ctx *macaron.Context) {

}

// {post} /address 新建一个资源
func AddressPostHandler(ctx *macaron.Context) {

}

// {put} /address/:id 更新某个指定资源的信息
func AddressPutHandler(ctx *macaron.Context) {

}

// {delete} /address/:id 删除某个指定资源
func AddressDeleteHandler(ctx *macaron.Context) {

}
