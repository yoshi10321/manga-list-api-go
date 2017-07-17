CREATE TABLE `mangas` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `img_url` varchar(255) NOT NULL,
  `number` int(11) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  PRIMARY KEY (`id`)
);
