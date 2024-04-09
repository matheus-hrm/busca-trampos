CREATE TABLE IF NOT EXISTS jobposts (
    `postId` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `ùserId` INT UNSIGNED NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `address` TEXT NOT NULL   
    `salary`    
    `status` ENUM('active','inactive','closed') NOT NULL DEFAULT 'active',
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`postId`)
    FOREIGN KEY (`userId`) REFERENCES users(`id`)
)