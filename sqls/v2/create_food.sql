CREATE TABLE `foods` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `title` varchar(100) NOT NULL COMMENT '商品名',
  `price` float DEFAULT '0' COMMENT '商品价格',
  `stock` int(11) DEFAULT '0' COMMENT '商品库存',
  `type` int(11) DEFAULT '0' COMMENT '商品类型',
  `create_time` datetime NOT NULL COMMENT '商品创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;