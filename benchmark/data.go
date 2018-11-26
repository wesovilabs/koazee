package benchmark

import "github.com/wesovilabs/koazee/benchmark/utils"

var strings10 = utils.ArrayOfString(1, 10, 10)
var strings100 = utils.ArrayOfString(1, 10, 100)
var strings1000 = utils.ArrayOfString(1, 10, 1000)
var strings5000 = utils.ArrayOfString(1, 10, 5000)
var numbers10 = utils.ArrayOfInt(0, 6, 10)
var numbers100 = utils.ArrayOfInt(0, 25, 100)
var numbers1000 = utils.ArrayOfInt(0, 200, 1000)
var numbers2500 = utils.ArrayOfInt(0, 2000, 2500)
