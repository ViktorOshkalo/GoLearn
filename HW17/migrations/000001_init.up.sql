CREATE TABLE products (
    `id` INT NOT NULL AUTO_INCREMENT,
    `catalog_id` INT NULL,
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

/* add test data */
INSERT INTO products (`catalog_id`, `name`, `description`, `created`) VALUES
(1, 'T-Shirt', 'Comfortable and light', UTC_TIMESTAMP()),
(2, 'Pants', 'Nice and warm', UTC_TIMESTAMP()),
(3, 'Sneakers', 'Soft and durable', UTC_TIMESTAMP());

INSERT INTO skus (`product_id`, `amount`, `price`, `unit`, `created`) VALUES
(1, 20, 200, 'unit', UTC_TIMESTAMP()),
(1, 30, 200, 'unit', UTC_TIMESTAMP()),
(1, 18, 200, 'unit', UTC_TIMESTAMP()),
(1, 15, 400, 'unit', UTC_TIMESTAMP()),
(1, 25, 400, 'unit', UTC_TIMESTAMP()),
(2, 5, 500, 'unit', UTC_TIMESTAMP()),
(2, 8, 500, 'unit', UTC_TIMESTAMP()),
(2, 6, 500, 'unit', UTC_TIMESTAMP());

INSERT INTO attributes (`sku_id`, `key`, `value`, `value_type`) VALUES
(1, 'Color', 'Red', 'string'),
(1, 'Size', 'S', 'string'),
(2, 'Color', 'Blue', 'string'),
(2, 'Size', 'M', 'string'),
(3, 'Color', 'Green', 'string'),
(3, 'Size', 'L', 'string'),
(4, 'Color', 'Black', 'string'),
(4, 'Size', 'M', 'string'),
(5, 'Color', 'Black', 'string'),
(5, 'Size', 'L', 'string'),
(6, 'Size', '9', 'integer'),
(7, 'Size', '10', 'integer'),
(8, 'Size', '11', 'integer');







