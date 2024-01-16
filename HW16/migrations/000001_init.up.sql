CREATE TABLE products (
    `id` INT NOT NULL AUTO_INCREMENT,
    `category_id` INT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(1024) NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NULL,
    `archived` DATETIME NULL,

    PRIMARY KEY (`id`)
);

/* product codes */
CREATE TABLE skus (
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_id` INT NOT NULL,
    `amount` FLOAT NOT NULL,
    `price` FLOAT NOT NULL,
    `unit` VARCHAR(10) NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NULL,
    `archived` DATETIME NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES products(`id`)
);

CREATE TABLE attributes (
    `sku_id` INT NOT NULL,
    `key` VARCHAR(20) NOT NULL,
    `value` VARCHAR(100) NOT NULL,
    `value_type` VARCHAR(10) NOT NULL,

    PRIMARY KEY (`sku_id`, `key`),
    FOREIGN KEY (`sku_id`) REFERENCES skus(`id`)
);



